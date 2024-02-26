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

func (d *Deque) PushFront(n *Node) {
	if d.Head == nil {
		d.Head = n
		d.Tail = n
	} else {
		n.Next = d.Head
		d.Head.Prev = n
		fmt.Printf("Current Head: %v, Next: %v, PRev: %v \n", d.Head.Key, d.Head.Next, d.Head.Prev)
		d.Head = n
		fmt.Printf("New Head: %v, Next: %v, PRev: %v \n", d.Head.Key, d.Head.Next, d.Head.Prev)
	}
}

func (d *Deque) PopBack() int {
	if d.Tail != nil && d.Tail == d.Head {
		tail := d.Tail
		d.Tail = nil
		d.Head = nil
		return tail.Key
	}
	if d.Tail != nil {
		fmt.Printf("Popping %d\n", d.Tail.Key)
		fmt.Printf("HEAD %d\n", d.Head.Key)
		tail := d.Tail
		d.Tail = tail.Prev
		fmt.Printf("New tail: %d\n", d.Tail.Key)
		d.Tail.Next = nil
		fmt.Printf("Current Head %d\n", d.Head.Key)
		return tail.Key
	}
	return -1
}

func (d *Deque) MoveToFront(n *Node) {
	if d.Head == n {
		return
	}
	if d.Tail == n {
		d.Tail = n.Prev
	}
	fmt.Printf("Next node is: %v \n", n.Next)
	fmt.Printf("PREV node is: %v \n", n.Prev)
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
	cache       map[int]*Node
	queue       Deque
	currentSize int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:    capacity,
		cache:       make(map[int]*Node),
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
		fmt.Println("Cache is full")
		c := lru.queue.Head
		for c != nil {
			fmt.Printf("%d -> ", c.Key)
			c = c.Next
		}
		fmt.Println()
		delete(lru.cache, lru.queue.PopBack())
		lru.currentSize--
		c = lru.queue.Head
		for c != nil {
			fmt.Printf("%d -> ", c.Key)
			c = c.Next
		}
		fmt.Println()
	}
	new_node := Node{Key: key, Value: value}
	lru.queue.PushFront(&new_node)
	lru.cache[key] = &new_node
	fmt.Printf("Next of new node: %v\n", new_node.Next)
	lru.currentSize++
	c := lru.queue.Head
	for c != nil {
		fmt.Printf("%d ->", c.Key)
		c = c.Next
	}
	fmt.Println()
}

func main() {
	my_lru := Constructor(1)
	my_lru.Put(2, 1)

	fmt.Println(my_lru.Get(2))
	// my_lru.Put(2, 2)

	my_lru.Put(3, 2)
	fmt.Println(my_lru.Get(2))
	fmt.Println(my_lru.Get(3))
	// my_lru.Put(4, 4)
	// fmt.Println(my_lru.Get(1))
	// fmt.Println(my_lru.Get(3))
	// fmt.Println(my_lru.Get(4))
}
