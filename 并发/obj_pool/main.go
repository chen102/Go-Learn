package main

import "fmt"
import "sync"

type New func() interface{}
type Pool struct {
	o       sync.Mutex
	objects []interface{}
	new     New
}

func NewPool(size int, new New) *Pool {
	objects := make([]interface{}, size)
	for i := 0; i < size; i++ {
		objects[i] = new() //初始化多少个
	}
	return &Pool{
		objects: objects,
		new:     new,
	}
}
func (p *Pool) Get() interface{} {
	p.o.Lock()
	defer p.o.Unlock()
	if len(p.objects) > 0 {
		obj := p.objects[0] //先进先出
		p.objects = p.objects[1:]
		return obj
	} else {
		return p.new()
	}
}
func (p *Pool) Put(obj interface{}) {
	p.o.Lock()
	defer p.o.Unlock()
	p.objects = append(p.objects, obj)
}
func main() {
	pool := NewPool(5, func() interface{} {
		fmt.Println("new")
		return 1
	})
	x := pool.Get()
	fmt.Println(x)
	pool.Put(x)

	x = pool.Get()
	x = pool.Get()
}
