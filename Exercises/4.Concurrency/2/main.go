package main

import (
	"errors"
	"fmt"
	"time"
)

const timeout = time.Second * 2

func main() {
	fmt.Println(realMain(1))
}

func realMain(goroutineDuration time.Duration) error {
	ch := make(chan int, 2)
	go g(ch, goroutineDuration)
	go g(ch, goroutineDuration)
	timer := time.NewTimer(timeout)
	for i := 0; i < 2; i++ {
		select {
		case <-ch:
			{
			}
		case <-timer.C:
			{
				return errors.New("время вышло")
			}
		}
	}
	close(ch)
	return nil
}

func g(ch chan int, t time.Duration) {
	timer := time.NewTimer(timeout)
	time.Sleep(t)
	select {
	case <-timer.C:
		fmt.Println("Время вышло")
		return
	default:
		ch <- 1
	}
}
