package main

import (
	"strconv"
	"fmt"
	"time"
	"github.com/my-stocks-pro/earnings-scheduler/vendor.orig/github.com/gin-gonic/gin"
	"net/http"
)

func NewTick() *Ticker {
	return &Ticker{
		Status:     false,
		Scheduler:  map[string]*time.Ticker{},
		APIGateWay: "http://127.0.0.1:9001/gateway",
		HTTPClient: http.Client{},
	}
}

func (t *Ticker) Run(schedulerType string, c *gin.Context) {

	go func() {
		for tick := range t.Scheduler[schedulerType].C {
			statusCode, err := t.RPC(schedulerType)
			c.JSON(statusCode, gin.H{
				"SchedulerType": schedulerType,
				"Error":         err,
				"Tick":          tick.String(),
			})

			fmt.Println(err)
		}
	}()
}

func (s *Scheduler) ParseUint(tick string) uint64 {
	u64, err := strconv.ParseUint(tick, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	return u64
}
