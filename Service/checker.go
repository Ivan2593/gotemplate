package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Measurable interface {
	GetMetrics() string
}

type Checkable interface {
	Measurable
	Ping() error
	GetID() string
	Health() bool
}

type Checker struct {
	slice []Checkable
	mu    sync.Mutex
}

func (c *Checker) Add(checkable Checkable) {
	c.mu.Lock()
	c.slice = append(c.slice, checkable)
	c.mu.Unlock()
}

func (c *Checker) String() string {
	var str string
	for _, value := range c.slice {
		str += value.GetID() + " "
	}
	return str
}

func (c *Checker) Check(ctx context.Context, value Checkable) {
	if !value.Health() {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(value.GetID() + " не работает")
		}
	}
}

func (c *Checker) Run(ctx context.Context) {
	fmt.Println("Метод Run начался")
	go c.run(ctx)
	fmt.Println("Метод Run завершился")
}

func (c *Checker) run(ctx context.Context) {
	fmt.Println("Метод run начался")
	ticker := time.NewTicker(5 * time.Second)
	for _ = range ticker.C {
		select {
		case <-ctx.Done():
			fmt.Println("Метод run завершился")
			return
		default:
			c.mu.Lock()
			for _, value := range c.slice {
				go c.Check(ctx, value)
			}
			c.mu.Unlock()
		}
	}
}

func (c *Checker) Stop(cancel context.CancelFunc) {
	fmt.Println("Метод Stop начался")
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("Метод Stop завершился")
}

func NewChecker() Checker {
	return Checker{}
}
