package workerpool1

import (
	"errors"
	"fmt"
	"sync"
)

type Pool struct {
	capacity int // 池的容量

	active chan struct{} // 对应上图中的active channel
	tasks  chan Task     // 对应上图中的task channel

	wg   sync.WaitGroup // 用于在pool销毁时等待所有worker退出
	quit chan struct{}  // 用于通知各个worker退出的信号channel
}

// Task 是一个对用户提交的请求的抽象，它的本质就是一个函数类型
type Task func()

const (
	defaultCapacity = 100
	maxCapacity     = 10000
)

func New(capacity int) *Pool {
	if capacity <= 0 {
		capacity = defaultCapacity
	}
	if capacity > maxCapacity {
		capacity = maxCapacity
	}
	p := &Pool{
		capacity: capacity,
		active:   make(chan struct{}, capacity),
		tasks:    make(chan Task),
		quit:     make(chan struct{}),
	}

	fmt.Printf("worker pool started\n")
	go p.run()

	return p
}

/*
循环体中使用select监视Pool类型实例的两个channel：quit和active。这种在for中使用select监视多个channel的实现，在Go代码中十分常见，是一种惯用法。

当接收到来自quit channel的退出“信号”时，这个Goroutine就会结束运行。而当active channel可写时，run方法就会创建一个新的worker Goroutine。
此外，为了方便在程序中区分各个worker输出的日志，我这里将一个从1开始的变量idx作为worker的编号，并把它以参数的形式传给创建worker的方法。
*/
func (p *Pool) run() {
	idx := 0
	for {
		select {
		case <-p.quit:
			return
		case p.active <- struct{}{}:
			idx++
			p.newWorker(idx)

		}
	}
}

// 封装创建新的worker goroutine的职责
// 在创建一个新的worker goroutine之前，newWorker方法会先调用p.wg.Add方法将WaitGroup的等待计数加一。
// 由于每个worker运行于一个独立的Goroutine中，newWorker方法通过go关键字创建了一个新的Goroutine作为worker。
//
// 新worker的核心，依然是一个基于for-select模式的循环语句，在循环体中，新worker通过select监视quit和tasks两个channel。
// 和前面的run方法一样，当接收到来自quit channel的退出“信号”时，这个worker就会结束运行。
// tasks channel中放置的是用户通过Schedule方法提交的请求，新worker会从这个channel中获取最新的Task并运行这个Task。
//
// 在新worker中，为了防止用户提交的task抛出panic，进而导致整个workerpool受到影响，
// 我们在worker代码的开始处，使用了defer+recover对panic进行捕捉，捕捉后worker也是要退出的，于是我们还通过<-p.active更新了worker计数器。
// 并且一旦worker goroutine退出，p.wg.Done也需要被调用，这样可以减少WaitGroup的Goroutine等待数量。
func (p *Pool) newWorker(i int) {
	p.wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("worker[%03d]: recover from panic[%s] and exit\n", i, err)
				<-p.active
			}
			p.wg.Done()
		}()

		fmt.Printf("worker[%03d]: is started\n", i)

		for {
			select {
			case <-p.quit:
				fmt.Printf("worker[%03d]: is exit\n", i)
				<-p.active
				return
			case task := <-p.tasks:
				fmt.Printf("worker[%03d]: received a task\n", i)
				task()
			}
		}
	}()
}

var ErrWorkerPoolFreed = errors.New("worker pool freed") // workerpool已终止运行

/*
Schedule方法的核心逻辑，是将传入的Task实例发送到workerpool的tasks channel中。
但考虑到现在workerpool已经被销毁的状态，我们这里通过一个select，检视quit channel是否有“信号”可读，
如果有，就返回一个哨兵错误ErrWorkerPoolFreed。如果没有，一旦p.tasks可写，提交的Task就会被写入tasks channel，以供pool中的worker处理。

这里要注意的是，这里的Pool结构体中的tasks是一个无缓冲的channel，如果pool中worker数量已达上限，而且worker都在处理task的状态，
那么Schedule方法就会阻塞，直到有worker变为idle状态来读取tasks channel，schedule的调用阻塞才会解除。
*/
func (p *Pool) Schedule(t Task) error {
	select {
	case <-p.quit:
		return ErrWorkerPoolFreed
	case p.tasks <- t:
		return nil
	}
}

func (p *Pool) Free() {
	close(p.quit)
	p.wg.Wait()
	fmt.Printf("worker pool freed \n")
}
