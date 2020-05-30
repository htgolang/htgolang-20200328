package pool

import (
	"log"
	"math"
	"sync"
)

type Task func() interface{}

//定义worker 池
type Pool struct {
	worker  int              //定义并发数量
	tasks   *Queue           //任务队列
	events  chan struct{}    // 任务通知管道
	Results chan interface{} //结果管道
	wg      sync.WaitGroup   //等待组
}

//创建worker池
func NewPool(worker int) *Pool {
	return &Pool{
		worker:  worker,
		tasks:   NewQueue(-1),
		events:  make(chan struct{}, math.MaxInt32),
		Results: make(chan interface{}, worker*3),
	}
}

//像任务池添加任务
func (p *Pool) AddTask(task Task) {
	p.tasks.Append(task)
	p.events <- struct{}{}
}

// 启动任务
func (p *Pool) Start() {
	for i := 0; i < p.worker; i++ {
		p.wg.Add(1)
		go func(i int) {
			// 忽略从管道中读取的数据
			for range p.events {
				// 从任务队列中读取任务
				e, err := p.tasks.Front()
				if err != nil {
					continue
				}
				log.Printf("worker %d run task", i)

				// 将队列中空接口数据转换为Task并进行执行
				if task, ok := e.(Task); ok {
					// 将结果放入到results管道
					p.Results <- task()
				}
			}
			p.wg.Done()
		}(i)
	}
}

// 等待任务执行完成
func (p *Pool) Wait() {
	close(p.events)
	p.wg.Wait()
	close(p.Results)
}
