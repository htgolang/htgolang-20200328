package taskpool

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"runtime"
	"sync"
)

type Task func() interface{}

type TaskPool struct {
	tasks   []Task
	limit   int
	workers int
	events  chan struct{}
	result  chan interface{}
	wg      sync.WaitGroup
}

func NewTaskPool(workers int) *TaskPool {
	return &TaskPool{
		workers: workers,
		limit: -1,
		events: make(chan struct{}, math.MaxInt8),
	}
}

func (t *TaskPool) AddTask(task Task) error {
	if t.limit != -1 && len(t.tasks) == t.limit {
		return errors.New("Can't add any more tasks!")
	}
	t.tasks = append(t.tasks, task)
	t.events <- struct{}{}
	return nil
}

func (t *TaskPool) getTask() (Task, error) {
	if len(t.tasks) == 0 {
		return nil, errors.New("There isn't any task.")
	}
	task := t.tasks[0]
	t.tasks = t.tasks[1:]
	return task, nil
}

func (t *TaskPool) Run() {
	close(t.events)
	t.result = make(chan interface{}, len(t.tasks))
	for i := 0; i < t.workers; i++ {
		t.wg.Add(1)
		go func(i int) {
			defer t.wg.Done()
			for range t.events {
				task, err := t.getTask()
				if err != nil {
					continue
				}
				taskname := runtime.FuncForPC(reflect.ValueOf(task).Pointer()).Name()
				fmt.Printf("Work %d: %s\n", i, taskname)
				t.result <- task()
			}
		}(i)
	}
	t.wg.Wait()
	close(t.result)
}

func (t *TaskPool) GetResult() {
	for result := range t.result {
		fmt.Printf("%v\n", result)
	}
}
