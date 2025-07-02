package workerpool2

import (
	"errors"
	"fmt"
	"sync"
)

type Pool struct {
	capacity int  // 池的容量
	preAlloc bool // 是否在创建pool的时候就预创建workers，默认值为：false

	// 当pool满的情况下，新的Schedule调用是否阻塞当前goroutine。默认值：true
	// 如果block = false，则Schedule返回ErrNoWorkerAvailInPool
	block  bool
	active chan struct{} // 对应上图中的active channel
	tasks  chan Task     // 对应上图中的task channel

	wg   sync.WaitGroup // 用于在pool销毁时等待所有worker退出
	quit chan struct{}  // 用于通知各个worker退出的信号channel
}

type Task func()

const (
	defaultCapacity = 100
	maxCapacity     = 10000
)

func New(capacity int, opts ...Option) *Pool {
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
		block:    true,
	}

	for _, opt := range opts {
		opt(p)
	}

	fmt.Printf("worker pool started(preAlloc=%t)\n", p.preAlloc)
	if p.preAlloc {
		for i := 0; i < p.capacity; i++ {
			p.newWorker(i)
			p.active <- struct{}{}
		}
	}

	go p.run()

	return p
}

func (p *Pool) run() {
	idx := len(p.active)

	if !p.preAlloc {
	loop:
		for t := range p.tasks {
			p.returnTask(t)
			select {
			case <-p.quit:
				return
			case p.active <- struct{}{}:
				idx++
				p.newWorker(idx)
			default:
				break loop
			}
		}
	}

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

func (p *Pool) returnTask(t Task) {
	go func() {
		p.tasks <- t
	}()
}

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

var ErrWorkerPoolFreed = errors.New("worker pool freed")         // workerpool已终止运行
var ErrNoIdleWorkerInPool = errors.New("no idle worker in pool") // workerpool已满，，没有空闲goroutine用于处理新任务

func (p *Pool) Schedule(t Task) error {
	select {
	case <-p.quit:
		return ErrWorkerPoolFreed
	case p.tasks <- t:
		return nil
	default:
		if p.block {
			p.tasks <- t
			return nil
		}
		return ErrNoIdleWorkerInPool
	}
}

func (p *Pool) Free() {
	close(p.quit)
	p.wg.Wait()
	fmt.Printf("worker pool freed(preAlloc=%t) \n", p.preAlloc)
}
