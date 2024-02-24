package main

type Node struct {
  Key int
  Value int
  Next *Node
  Prev *Node
}

type Deuque struct {
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
    d.Head = n
  }


