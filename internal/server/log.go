package server

import (
	"fmt"
	"sync"
)

type Log struct {
	mu      sync.Mutex
	records []Record
}

func NewLog() *Log {
	return &Log{}
}

func (c *Log) Apprend(record Record) (uinit64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
}
