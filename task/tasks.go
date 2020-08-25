package task

import (
	"context"
	"sync"
)

type Tasks interface {
	AddTask(Task)
	AddTaskFunc(Func)
	Run(context.Context)
	Wait() error
	Cancel()
}

func NewTasks() Tasks {
	return &tasks{}
}

type tasks struct {
	response   Response
	mutex      sync.Mutex
	tasks      []Task
	cancelFunc func()
	wg         sync.WaitGroup
	canceld    bool
}

func (t *tasks) AddTask(task Task) {
	t.tasks = append(t.tasks, task)
}

func (t *tasks) AddTaskFunc(f Func) {
	task := NewTask(f)
	t.AddTask(task)
}

func (t *tasks) Cancel() {
	t.cancelFunc()
}

func (t *tasks) Run(ctx context.Context) {
	ctx, cancelFunc := context.WithCancel(ctx)
	t.cancelFunc = cancelFunc
	t.response = Response(make([]error, len(t.tasks)))
	t.mutex = sync.Mutex{}
	t.wg = sync.WaitGroup{}
	t.wg.Add(len(t.tasks))
	for i, task := range t.tasks {
		go func(i int, task Task) {
			defer t.wg.Done()
			task.Run(ctx)
			err := task.Wait()
			t.mutex.Lock()
			if err != nil && !t.canceld {
				t.response[i] = err
				t.canceld = true
				t.Cancel()
			}
			t.mutex.Unlock()
		}(i, task)
	}
}

func (t *tasks) Wait() error {
	defer t.cancelFunc()
	t.wg.Wait()
	response := t.response
	return response.Error()
}
