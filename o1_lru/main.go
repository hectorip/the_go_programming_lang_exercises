package main

import "fmt"

type Node struct {
	Key   int
	Value int
	Next  *Node
	Prev  *Node
}

type Deque struct {
	Head *Node
	Tail *Node
}

func (d *Deque) PushFront(n Node) {
	if d.Head == nil {
		d.Head = &n
		d.Tail = &n
	} else {
		n.Next = d.Head
		d.Head.Prev = &n
		d.Head = &n
	}
}

func (d *Deque) PopBack() int {
	if d.Tail != nil {
		tail := d.Tail
		d.Tail = d.Tail.Prev
		d.Tail.Next = nil
		return tail.Key
	}
	return -1
}

func (d *Deque) MoveToFront(n Node) {
	if n.Prev != nil {
		n.Prev.Next = n.Next
	}
	if n.Next != nil {
		n.Next.Prev = n.Prev
	}
	d.PushFront(n)
}

type LRUCache struct {
	capacity    int
	cache       map[int]Node
	queue       Deque
	currentSize int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:    capacity,
		cache:       make(map[int]Node),
		queue:       Deque{},
		currentSize: 0,
	}
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cache[key]; ok {
		lru.queue.MoveToFront(node)
		return node.Value
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.cache[key]; ok {
		node.Value = value
		lru.queue.MoveToFront(node)
		return
	}
	if lru.currentSize == lru.capacity {
		delete(lru.cache, lru.queue.PopBack())
		lru.currentSize--
	}
	new_node := Node{Key: key, Value: value}
	lru.queue.PushFront(new_node)
	lru.cache[key] = new_node
	lru.currentSize++
}

func main() {
	my_lru := Constructor(2)
	my_lru.Put(1, 1)
	my_lru.Put(2, 2)

	fmt.Println(my_lru.Get(1))
	my_lru.Put(3, 3)
	fmt.Println(my_lru.Get(2))
	my_lru.Put(4, 4)
	fmt.Println(my_lru.Get(1))
	fmt.Println(my_lru.Get(3))
	fmt.Println(my_lru.Get(4))
}
