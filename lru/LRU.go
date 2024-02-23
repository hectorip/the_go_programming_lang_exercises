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

func (lru *LRU) Get(key int) int {
	if val, ok := lru.cache[key]; ok {
		return val
	}
	return -1
}

func (lru *LRU) Put(key, value int) {
	c_value := lru.Get(key)
	if c_value != -1 {
		lru.cache[key] = value
		lru.queue[capacity-1] = key
	} else {
		if lru.curretSize >= lru.capacity {
			key_to_remove := lru.queue[0]
			lru.queue = lru.queue[1:]
			delete(lru.cache, key_to_remove)
			lru.curretSize--
		}
		lru.cache[key] = value
		n := 0
		for _, v := range lru.queue {
			if v != key {
				lru.queue[n] = v
				n++
			}
		}
		lru.queue[n] = key
	}
}
