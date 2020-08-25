package lrucache

import (
	"container/list"
	"errors"
	"fmt"
	"sync"

	"github.com/y851592226/cate/convert"
)

type Cacher interface {
	Set(interface{}, interface{})
	Get(interface{}) (interface{}, bool)
	GetDefault(interface{}, interface{}) interface{}
	GetInt(interface{}) (int, bool)
	GetIntDefault(interface{}, int) int
	GetString(interface{}) (string, bool)
	GetStringDefault(interface{}, string) string
	Len() int
}

type Options struct {
	Capacity  int
	OnDeleted func(key, value interface{})
}

type cache struct {
	capacity  int
	onDeleted func(key, value interface{})
	mutex     sync.Mutex
	data      map[interface{}]*list.Element
	list      *list.List
}

type entry struct {
	key   interface{}
	value interface{}
}

func NewCacher(opt *Options) (Cacher, error) {
	if opt.Capacity == 0 {
		return nil, errors.New("capacity should not be zero")
	}
	return &cache{
		capacity:  opt.Capacity,
		onDeleted: opt.OnDeleted,
		mutex:     sync.Mutex{},
		data:      map[interface{}]*list.Element{},
		list:      list.New(),
	}, nil
}

func (c *cache) Set(key, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	el, ok := c.data[key]
	if ok {
		el.Value.(*entry).value = value
		c.list.MoveToFront(el)
		return
	}
	el = c.list.PushFront(&entry{key, value})
	c.data[key] = el
	if c.list.Len() > c.capacity {
		c.removeOldest()
	}
}

func (c *cache) Get(key interface{}) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	el, ok := c.data[key]
	if !ok {
		return nil, false
	}
	c.list.MoveToFront(el)
	return el.Value.(*entry).value, true
}

func (c *cache) removeOldest() {
	el := c.list.Back()
	if el != nil {
		e := c.list.Remove(el)
		delete(c.data, e.(*entry).key)
		if c.onDeleted != nil {
			c.onDeleted(e.(*entry).key, e.(*entry).value)
		}
	}
}

func (c *cache) GetDefault(key, def interface{}) interface{} {
	v, ok := c.Get(key)
	if !ok {
		return def
	}
	return v
}

func (c *cache) GetInt(key interface{}) (int, bool) {
	v, ok := c.Get(key)
	if !ok {
		return 0, false
	}
	return v.(int), true
}

func (c *cache) GetIntDefault(key interface{}, def int) int {
	v, ok := c.Get(key)
	if !ok {
		return def
	}
	return convert.AsIntDefault(v, def)
}

func (c *cache) GetString(key interface{}) (string, bool) {
	v, ok := c.Get(key)
	if !ok {
		return "", false
	}
	return v.(string), true
}

func (c *cache) GetStringDefault(key interface{}, def string) string {
	v, ok := c.Get(key)
	if !ok {
		return def
	}
	return convert.AsStringDefault(v, def)
}

func (c *cache) Len() int {
	fmt.Println(c.data)
	return len(c.data)
}
