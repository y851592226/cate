# tasks

> Concurrent task control

**Simple Example**

```go
func ExampleTasks() {
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
	tasks := task.NewTasks()
	tasks.AddTaskFunc(task0)
	tasks.AddTaskFunc(task1)
	tasks.AddTaskFunc(task2)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*250)
	defer cancel()
	tasks.Run(ctx)
	err := tasks.Wait()
	fmt.Println(err) // task:2 error:context deadline exceeded
	fmt.Println(a) // 3
}
```