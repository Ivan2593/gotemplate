package main

import (
	"context"
	"fmt"
	"time"
)

const timeout = time.Second * 2

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	realMain(ctx, 5)
	time.Sleep(timeout)
}

func realMain(ctx context.Context, num int) {
	myCtx, cancel := context.WithCancel(ctx)
	for i := 1; i < num+1; i++ {
		go goCont(myCtx, i)
	}
	time.Sleep(timeout)
	cancel()
}

func goCont(ctx context.Context, i int) {
	<-ctx.Done()
	fmt.Printf("завершилась горутина - %v\n", i)
	return

}
