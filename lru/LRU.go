package main

import "fmt"

func main() {
	fmt.Println("Funcinamiento de LRU en Golang!")
}

type LRU struct {
	capacity   int
	cache      map[int]int
	queue      []int
	curretSize int
}

func NewLRU(capacity int) *LRU {
	return &LRU{
		capacity:   capacity,
		cache:      make(map[int]int),
		queue:      make([]int, capacity),
		curretSize: 0,
	}
}

func (lru *LRU) move_to_end(key int) {
	n := 0
	for _, v := range lru.queue {
		if v != key {
			lru.queue[n] = v
			n++
		}
	}
	lru.queue = lru.queue[:n]
	lru.queue = append(lru.queue, key)
}

func (lru *LRU) Get(key int) int {
	if val, ok := lru.cache[key]; ok {
		lru.move_to_end(key)
		return val
	}
	return -1
}

func (lru *LRU) Put(key, value int) {
	if _, ok := lru.cache[key]; ok {
		lru.cache[key] = value
		lru.move_to_end(key)
	} else {
		if lru.curretSize >= lru.capacity {
			key_to_remove := lru.queue[0]
			lru.queue = lru.queue[1:]
			delete(lru.cache, key_to_remove)
			lru.curretSize--
		}
		lru.cache[key] = value
		lru.queue = append(lru.queue, key)
	}
}
