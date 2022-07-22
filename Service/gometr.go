package main

import (
	"strconv"
	"time"
)

type GoMetrClient struct {
	ServiceID int
	timeOut   time.Duration
}

func (g GoMetrClient) GetID() string {
	return strconv.Itoa(g.ServiceID)
}

func (g GoMetrClient) GetMetrics() string {
	return ""
}

func (g GoMetrClient) Ping() error {
	return nil
}

func (g GoMetrClient) getHealth(ch chan HealthCheck) HealthCheck {
	var s string
	time.Sleep(2 * time.Second)
	if g.ServiceID%2 == 0 {
		s = PassStatus
	} else {
		s = FailStatus
	}
	ch <- HealthCheck{ServiceId: g.GetID(), Status: s}
	return HealthCheck{ServiceId: g.GetID(), Status: s}
}

func (g GoMetrClient) Health() bool {
	timer := time.NewTimer(g.timeOut * time.Second)
	ch := make(chan HealthCheck, 1)
	go g.getHealth(ch)
	select {
	case <-timer.C:
		return false
	case h := <-ch:
		if h.Status == PassStatus {
			return true
		} else {
			return false
		}
	}
}

func NewGoMetrClient(id int, timeOut time.Duration) GoMetrClient {
	return GoMetrClient{
		ServiceID: id,
		timeOut:   timeOut,
	}
}
