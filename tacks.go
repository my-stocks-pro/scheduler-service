package main

import (
	"fmt"
	"github.com/my-stocks-pro/earnings-scheduler/vendor.orig/github.com/jasonlvhit/gocron"
	"net/http"
	"bytes"
)

func (s Scheduler) Run() {

	go func() {
		fmt.Println("NewScheduler EarningsTask Create")

		approvedScheduler := gocron.NewScheduler()
		approvedScheduler.Every(EarningsTick).Seconds().Do(func() {
			s.RPC("Earnings")
		})
		<-approvedScheduler.Start()
	}()

	go func() {
		fmt.Println("NewScheduler ApprovedTask Create")

		approvedScheduler := gocron.NewScheduler()
		approvedScheduler.Every(ApprovedTick).Seconds().Do(func() {
			s.RPC("Approved")
		})
		<-approvedScheduler.Start()
	}()

	go func() {
		fmt.Println("NewScheduler RejectedTask Create")

		approvedScheduler := gocron.NewScheduler()
		approvedScheduler.Every(RejectedTick).Seconds().Do(func() {
			s.RPC("Rejected")
		})
		<-approvedScheduler.Start()
	}()
}

func (s Scheduler) RPC(schedulerType string) {

	fmt.Println(schedulerType)

	req, errReq := http.NewRequest(http.MethodPost, API, bytes.NewReader([]byte(schedulerType)))
	if errReq != nil {
		fmt.Println(errReq)
	}

	_, errResp := s.HTTPClient.Do(req)
	if errResp != nil {
		fmt.Println(errReq)
	}
}
