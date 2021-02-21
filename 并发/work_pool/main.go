package main

import (
	"fmt"
	"math"
	"sync"
)

const (
	defaultCap = 1024
)

type Queue struct {
	elements []interface{}
	locker   sync.Mutex
	limit    int //队列元素数量上限
}

func NewQueue(limit int) *Queue {
	return &Queue{
		elements: make([]interface{}, 0, defaultCap),
		limit:    limit,
	}
}
func (q *Queue) Append(e interface{}) error {
	q.locker.Lock()
	defer q.locker.Unlock()
	if q.limit != -1 && len(q.elements) >= q.limit {
		return fmt.Errorf("queue is full", q.limit)
	}
	q.elements = append(q.elements, e)
	return nil
}
func (q *Queue) Front() (interface{}, error) {
	if len(q.elements) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}
	e := q.elements[0]
	q.elements = q.elements[1:]
	return e, nil
}
func (q *Queue) Len() int {
	return len(q.elements)
}

type Pool struct {
	worker  int
	tasks   *Queue
	Results chan interface{}
	events  chan struct{} //任务通知
	wg      sync.WaitGroup
}
type Task func() interface{}

func NewPool(worker int) *Pool {
	return &Pool{
		worker:  worker, //并发数量
		tasks:   NewQueue(-1),
		events:  make(chan struct{}, math.MaxInt32),
		Results: make(chan interface{}, worker*3),
	}
}
func (p *Pool) AddTask(task Task) {
	p.tasks.Append(task)
	p.events <- struct{}{}
}
func (p *Pool) Start() {
	for i := 0; i < p.worker; i++ {
		p.wg.Add(1)
		go func() {
			for range p.events {
				e, err := p.tasks.Front()
				if err != nil {
					continue
				}
				if task, ok := e.(Task); ok {
					p.Results <- task()
				}
			}
			p.wg.Done()
		}()
	}
}
func (p *Pool) Wait() {
	close(p.events)
	p.wg.Wait()
	close(p.Results)
}
func main() {

	worker := NewPool(5)
	worker.AddTask(func() interface{} {
		return 1
	})
	worker.AddTask(func() interface{} {
		return 2
	})
	worker.AddTask(func() interface{} {
		return 3
	})
	worker.Start()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for result := range worker.Results {
			fmt.Println(result)
		}
		wg.Done()
	}()
	worker.Wait()
	wg.Wait()
}
