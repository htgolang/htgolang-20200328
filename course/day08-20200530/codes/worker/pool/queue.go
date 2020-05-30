package pool

import (
	"fmt"
	"sync"
)

const (
	defaultCap = 1024
)

// 定义队列
type Queue struct {
	elements []interface{} //队列中元素
	locker   sync.Mutex    //并发安全定义锁
	limit    int           //定义队列中元素数量上限
}

// 创建队列
func NewQueue(limit int) *Queue {
	return &Queue{
		elements: make([]interface{}, 0, defaultCap),
		limit:    limit,
	}
}

// 队列中加入元素，当达到数量上限返回错误
func (q *Queue) Append(e interface{}) error {
	q.locker.Lock()
	defer q.locker.Unlock()
	if q.limit != -1 && len(q.elements) >= q.limit {
		return fmt.Errorf("queue is full", q.limit)
	}
	q.elements = append(q.elements, e)
	return nil
}

// 队列中获取元素，当队列为空返回错误
func (q *Queue) Front() (interface{}, error) {
	if len(q.elements) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}
	e := q.elements[0]
	q.elements = q.elements[1:]
	return e, nil
}

// 队列中元素数量
func (q *Queue) Len() int {
	return len(q.elements)
}
