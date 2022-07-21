package main

import (
	"fmt"
	"sync"
)

var (
	a = []int{1, 3, 5, 7, 9}
	b = []int{2, 4, 6, 8, 10}
	c = []int{}
)

func main() {
	realMain(a, b, &c)
	fmt.Println(c)
}

func realMain(a, b []int, c *[]int) {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	chA := make(chan int, 1)
	chB := make(chan int, 1)
	chB <- 0
	go oddNumbers(a, c, chA, chB, wg)
	go numbers(b, c, chA, chB, wg)
	wg.Wait()
}

func oddNumbers(a []int, c *[]int, chA chan int, chB chan int, wg *sync.WaitGroup) {
	for _, value := range a {
		<-chB
		*c = append(*c, value)
		chA <- value
	}
	wg.Done()
}

func numbers(b []int, c *[]int, chA chan int, chB chan int, wg *sync.WaitGroup) {
	for _, value := range b {
		<-chA
		*c = append(*c, value)
		chB <- value
	}
	wg.Done()
}
