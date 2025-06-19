package main

// CAS操作，当时还没有抽象出atomic包
func cas(val *int32, old, new int32) bool
func semacquire(*int32) // sem  acquire
func semrelease(*int32) // sem  release

// 互拆锁的结构
type Mutex struct {
	key  int32 // 锁是否被持有的标识
	sema int32 // 信号量专用，用以阻塞/唤醒goroutine
}

// 保证成功在val上增加delta的值
func xadd(val *int32, delta int32) (new int32) {
	for {
		v := *val
		if cas(val, v, v+delta) {
			return v + delta
		}
	}
	panic("unreached")
}

// 加锁
func (m *Mutex) Lock() {
	// 如果锁已被持有的话，进入等待状态
	if xadd(&m.key, 1) == 1 { //标识加1，如果等于1，成功获取到锁
		return
	}
	semacquire(&m.sema) // 否则阻塞等待
}

func (m *Mutex) Unlock() {
	if xadd(&m.key, -1) == 0 { // 标识减1，如果等于0，说明没有其他的goroutine等待这把锁
		return
	}
	semrelease(&m.sema) // 唤醒其他等待（阻塞）的goroutine
}
