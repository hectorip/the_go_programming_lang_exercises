package main

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
		tail = d.Tail
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


