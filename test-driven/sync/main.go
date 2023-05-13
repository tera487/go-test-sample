package main

import "sync"

// Counter は
type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter returns a new Counter.
func NewCounter() *Counter {
	return &Counter{}
}

// Inc はプラス
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

// Value は値を取得
func (c *Counter) Value() int {
	return c.value
}
