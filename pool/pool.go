package pool

import (
	"context"
	"sync"
	"time"
)

type Pool interface {
	Get(ctx context.Context) (interface{}, error)
	Put(interface{}, error)
}

type Options struct {
	NewFunc     func() (interface{}, error)
	MaxRetries  int
	MaxSize     int
	MaxIdleSize int
	MaxIdlelime time.Duration
}

type pool struct {
	mutex  sync.Mutex
	values chan *Value
	count  int
	opt    *Options
}

type Value struct {
	expireTime time.Time
	value      interface{}
}

func (v *Value) Expired() bool {
	return !v.expireTime.IsZero() && time.Now().After(v.expireTime)
}

func NewPool(opt *Options) Pool {
	return &pool{
		mutex:  sync.Mutex{},
		values: make(chan *Value, opt.MaxIdleSize),
		count:  0,
		opt:    opt,
	}
}

func (p *pool) Get(ctx context.Context) (interface{}, error) {
	// 优先使用缓存channel中的value
ForEnd:
	for {
		select {
		case v := <-p.values:
			if !v.Expired() {
				return v.value, nil
			}
			p.mutex.Lock()
			p.count--
			p.mutex.Unlock()

		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			break ForEnd
		}
	}
	return p.getSlow(ctx)
}

func (p *pool) getSlow(ctx context.Context) (interface{}, error) {
	p.mutex.Lock()
	if p.opt.MaxSize > p.count {
		p.count++
		p.mutex.Unlock()
		return p.getNew()
	}
	p.mutex.Unlock()
	for {
		select {
		case v := <-p.values:
			if v.Expired() {
				return v.value, nil
			}
			p.mutex.Lock()
			p.count--
			p.mutex.Unlock()

		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

func (p *pool) getNew() (value interface{}, err error) {
	for i := 0; i <= p.opt.MaxRetries; i++ {
		value, err = p.opt.NewFunc()
		if err != nil {
			continue
		}
		return value, nil
	}
	return nil, err
}

func (p *pool) Put(value interface{}, err error) {
	if err != nil {
		p.mutex.Lock()
		p.count--
		p.mutex.Unlock()
		return
	}
	v := &Value{
		value: value,
	}
	if p.opt.MaxIdlelime > 0 {
		v.expireTime = time.Now().Add(p.opt.MaxIdlelime)
	}
	select {
	case p.values <- v:
		return
	default:
		p.mutex.Lock()
		p.count--
		p.mutex.Unlock()
		return
	}
}
