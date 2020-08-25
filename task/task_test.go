package task

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"context"
	"errors"
	"sync/atomic"
	"time"
)

var _ = Describe("Task", func() {
	var a int64
	task0 := func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Millisecond * 100):
			atomic.AddInt64(&a, 1)
		}
		return nil
	}
	task1 := func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Millisecond * 200):
			atomic.AddInt64(&a, 2)
		}
		return nil
	}
	task2 := func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Millisecond * 300):
			atomic.AddInt64(&a, 3)
		}
		return nil
	}
	task3 := func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Millisecond * 400):
			return errors.New("test error")
		}
	}
	It("test error1", func() {
		tasks := NewTasks()
		tasks.AddTaskFunc(task0)
		tasks.AddTaskFunc(task1)
		tasks.AddTaskFunc(task2)
		tasks.AddTaskFunc(task3)
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, time.Millisecond*350)
		defer cancel()
		tasks.Run(ctx)
		err := tasks.Wait()
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("task:3 error:context deadline exceeded"))
	})
})
