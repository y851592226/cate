package ttlcache

import (
	"context"
	"sync"
	"time"
)

type Cacher interface {
	Get(context.Context, interface{}) (interface{}, error)
	GetDefault(context.Context, interface{}, interface{}) interface{}
}

type Options struct {
	TTL     time.Duration
	Fetcher func(context.Context, interface{}) (interface{}, error)
}

type item struct {
	data       interface{}
	expireTime time.Time
}

func newItem(data interface{}, TTL time.Duration) *item {
	now := time.Now()
	if TTL > 0 {
		return &item{
			data:       data,
			expireTime: now.Add(TTL),
		}
	}
	return &item{
		data: data,
	}
}

type cache struct {
	mutex       sync.Mutex
	waitChan    map[interface{}]chan struct{}
	dataRWMutex sync.RWMutex
	data        map[interface{}]*item
	opt         Options
}

func NewCache(opt Options) Cacher {
	return &cache{
		mutex:       sync.Mutex{},
		waitChan:    map[interface{}]chan struct{}{},
		dataRWMutex: sync.RWMutex{},
		data:        map[interface{}]*item{},
		opt:         opt,
	}
}

func (c *cache) Get(ctx context.Context, key interface{}) (interface{}, error) {
	c.dataRWMutex.RLock()
	i, found := c.data[key]
	c.dataRWMutex.RUnlock()
	now := time.Now()
	if found && (c.opt.TTL == 0 || i.expireTime.After(now)) {
		return i.data, nil
	}
	// not found OR expired
	return c.getSingleflight(ctx, key)
}

func (c *cache) GetDefault(ctx context.Context, key interface{}, def interface{}) interface{} {
	v, err := c.Get(ctx, key)
	if err != nil {
		return def
	}
	return v
}

func (c *cache) getSingleflight(ctx context.Context, key interface{}) (interface{}, error) {
	c.mutex.Lock()
	wait, ok := c.waitChan[key]
	if !ok {
		wait = make(chan struct{})
		c.waitChan[key] = wait
	}
	c.mutex.Unlock()
	// another goroutine is in process, just wait
	if ok {
		<-wait
		return c.Get(ctx, key)
	}
	data, err := c.get(ctx, key)
	c.notifyWait(key, wait)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (c *cache) get(ctx context.Context, key interface{}) (interface{}, error) {
	data, err := c.opt.Fetcher(ctx, key)
	if err != nil {
		return nil, err
	}
	c.dataRWMutex.Lock()
	c.data[key] = newItem(data, c.opt.TTL)
	c.dataRWMutex.Unlock()
	return data, nil
}

func (c *cache) notifyWait(key interface{}, wait chan struct{}) {
	c.mutex.Lock()
	delete(c.waitChan, key)
	close(wait)
	c.mutex.Unlock()
}
