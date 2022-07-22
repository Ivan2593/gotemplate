package main

import (
	"context"
	"fmt"
	"time"
)

type HealthCheck struct {
	ServiceId string
	Status    string
}

const PassStatus = "pass"
const FailStatus = "fail"

func main() {
	ch := NewChecker()
	g1 := NewGoMetrClient(1, 2)
	g2 := NewGoMetrClient(2, 2)
	g3 := NewGoMetrClient(3, 2)
	g4 := NewGoMetrClient(4, 4)
	ch.Add(g1)
	ch.Add(g2)
	ch.Add(g3)
	ch.Add(g4)
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println(ch)
	ch.Run(ctx)
	time.Sleep(15 * time.Second)
	ch.Stop(cancel)
	time.Sleep(6 * time.Second)
}
