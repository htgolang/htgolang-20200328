package pool

import (
	"fmt"
	"sync"
)

type Task func() interface{}

type Pool struct {
	queue []Task
	caps	int //queue 最大长度
	worker int // 并发数
	taskChan chan Task
	result chan interface{} //任务结果队列
	wg sync.WaitGroup //等待
	wg2 sync.WaitGroup
	lock	sync.Mutex //定义锁
}

func NewPool(caps,worker int) *Pool  {
	return &Pool{
		caps:caps,
		worker:worker,
		taskChan:make(chan Task,worker),
		result:make(chan interface{},caps*2),
	}
}

func (p *Pool)AddTask(task Task) error {
	defer p.lock.Unlock()
	p.lock.Lock()
	if len(p.queue)<p.caps{
		p.queue = append(p.queue,task)
	}else {
		return fmt.Errorf("tasks pool is full")
	}
	return nil
}

func (p *Pool)PopTask()( interface{} ,error){
	defer p.lock.Unlock()
	p.lock.Lock()
	if len(p.queue)==0{
		return nil,fmt.Errorf("queue is empty")
	}else {
		task := p.queue[0]
		p.queue = p.queue[1:]
		return task,nil
	}
}

func (p *Pool) Start() chan interface{}{
	for i:=0;i<len(p.queue);i++{
		elemt,err := p.PopTask()
		if err != nil{
			break
		}
		if task,ok := elemt.(Task);ok{
			p.wg2.Add(1)
			go func(i int) {
				fmt.Println(i,"queue start")

				p.taskChan <- task

				fmt.Println(i,"queue end")
				p.wg2.Done()
			}(i)
			p.wg.Add(1)
			go func(i int) {
				fmt.Println(i,"task start")
				task :=<- p.taskChan
				p.result <- task()
				fmt.Println(i,"task end")
				p.wg.Done()
			}(i)
		}
	}

	p.wg2.Wait()
	close(p.taskChan)
	p.wg.Wait()
	close(p.result)
	return p.result
}