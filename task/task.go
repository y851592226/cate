package task

import (
	"context"
	"fmt"
	"runtime/debug"
)

type Func func(context.Context) error

type Task interface {
	Run(context.Context)
	Wait() error
}

type task struct {
	ctx  context.Context
	run  Func
	err  error
	done chan struct{}
}

func (t *task) Run(ctx context.Context) {
	t.ctx = ctx
	go func() {
		defer func() {
			if err := recover(); err != nil {
				stack := debug.Stack()
				t.err = fmt.Errorf("%v\n %s", err, string(stack))
				close(t.done)
			}
		}()
		t.err = t.run(ctx)
		close(t.done)
	}()
}

func (t *task) Wait() error {
	select {
	case <-t.ctx.Done():
		return t.ctx.Err()
	case <-t.done:
		return t.err
	}
}

func NewTask(f Func) Task {
	return &task{
		run:  f,
		done: make(chan struct{}),
	}
}
