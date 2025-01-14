package fyf

import "container/list"

type LRUCache struct {
	l        *list.List
	m        map[int]*Node
	capacity int
}

type Node struct {
	v int
	e *list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		l:        list.New(),
		m:        map[int]*Node{},
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		this.l.MoveToFront(v.e)
		return v.v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.m[key]; ok {
		this.l.MoveToFront(v.e)
		this.m[key].v = value
	} else {
		elem := this.l.PushFront(key)
		this.m[key] = &Node{
			v: value,
			e: elem,
		}
	}
	if len(this.m) > this.capacity {
		delete(this.m, this.l.Remove(this.l.Back()).(int))
	}
}
