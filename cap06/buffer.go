package main

import "sync"

type Buffer struct {
	mu   sync.Mutex
	cond *sync.Cond
	data []int
}

func NewBuffer() *Buffer {
	b := &Buffer{data: []int{}}
	b.cond = sync.NewCond(&b.mu)
	return b
}
