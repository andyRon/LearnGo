package main

import (
	"fmt"
	"sync"
	"time"
)

/*
goroutine使用模式: 发布-订阅（Pub/Sub）模式

在 Go 语言里，goroutine 的发布-订阅（Pub/Sub）模式是一种消息传递模式，允许发布者（Publisher）发送消息到一个或多个主题（Topic），而订阅者（Subscriber）可以订阅感兴趣的主题并接收消息。
该模式实现了发布者和订阅者之间的解耦，提高了系统的可扩展性和灵活性。
*/

func main() {
	broker := NewBroker()
	defer broker.Close()

	// 启动生产者
	go producer(broker)

	// 启动多个消费者
	for i := 1; i <= 3; i++ {
		go consumer(broker, i)
	}

	// 等待一段时间，确保消息处理完成
	time.Sleep(2 * time.Second)
}

// Broker 表示消息代理，负责管理主题和订阅者
type Broker struct {
	subscribers map[string][]chan interface{}
	mu          sync.Mutex
}

// NewBroker 创建一个新的消息代理
func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[string][]chan interface{}),
	}
}

// Subscribe 订阅指定主题
func (b *Broker) Subscribe(topic string) chan interface{} {
	b.mu.Lock()
	defer b.mu.Unlock()
	ch := make(chan interface{}, 10)
	b.subscribers[topic] = append(b.subscribers[topic], ch)
	return ch
}

// Publish 发布消息到指定主题
func (b *Broker) Publish(topic string, message interface{}) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if subscribers, ok := b.subscribers[topic]; ok {
		for _, ch := range subscribers {
			select {
			case ch <- message:
			default:
				// 通道满时，丢弃消息
			}
		}
	}
}

// Close 关闭所有订阅通道
func (b *Broker) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()
	for topic, subscribers := range b.subscribers {
		for _, ch := range subscribers {
			close(ch)
		}
		delete(b.subscribers, topic)
	}
}

// 生产者函数，发布消息
func producer(broker *Broker) {
	for i := 0; i < 5; i++ {
		broker.Publish("topic1", i)
		fmt.Printf("Published: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
}

// 消费者函数，订阅消息
func consumer(broker *Broker, id int) {
	ch := broker.Subscribe("topic1")
	for msg := range ch {
		fmt.Printf("Consumer %d received: %v\n", id, msg)
		time.Sleep(200 * time.Millisecond)
	}
}
