package ratelimit

import (
	"sync"
	"sync/atomic"
	"time"
)

type TokenBucket struct {
	Num          int64         //每次生成令牌的数量
	Capacity     int64         //bucket的总容量
	fillInterval time.Duration //令牌生成的时间间隔间隔
	count        int64         // 当前可用令牌数
	mutex        *sync.Mutex
	cond         *sync.Cond
}

func NewTokenBucket(num, capacity int64, fillInterval time.Duration) Limiter {
	mutex := &sync.Mutex{}
	tb := &TokenBucket{
		Num:          num,
		Capacity:     capacity,
		fillInterval: fillInterval,
		count:        num,
		mutex:        mutex,
		cond:         sync.NewCond(mutex),
	}

	go func() {
		for range time.Tick(tb.fillInterval) {
			tb.mutex.Lock()
			count := atomic.LoadInt64(&tb.count)
			if count < tb.Capacity {
				atomic.AddInt64(&tb.count, tb.Num)
				// 通知阻塞的协程
				if count <= 0 {
					tb.cond.Broadcast()
				}
			}
			tb.mutex.Unlock()
		}
	}()
	return tb
}

func (tb *TokenBucket) SetNum(num int64) {
	tb.Num = num
}

func (tb *TokenBucket) SetCapacity(capacity int64) {
	tb.Capacity = capacity
}

func (tb *TokenBucket) Get() bool {
	c := atomic.AddInt64(&tb.count, -1)
	if c >= 0 {
		return true
	}
	atomic.AddInt64(&tb.count, 1)
	return false
}

func (tb *TokenBucket) BlockingGet() {
	ok := tb.Get()
	if ok {
		return
	}
	tb.BlockingGetSlow()
}

func (tb *TokenBucket) BlockingGetSlow() {
	for {
		tb.mutex.Lock()
		ok := tb.Get()
		if ok {
			tb.mutex.Unlock()
			return
		}
		tb.cond.Wait()
		tb.mutex.Unlock()
	}
}
