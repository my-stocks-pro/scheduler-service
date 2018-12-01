package main

import (
	"github.com/my-stocks-pro/earnings-scheduler/vendor.orig/github.com/gin-gonic/gin"
	"time"
	"net/http"
)

type Schedul struct {
	Status       string `json:"status"`
	EarningsTick int    `json:"earnings_tick"`
	ApprovedTick int    `json:"approved_tick"`
	RejectedTick int    `json:"rejected_tick"`
}

func (s *Scheduler) Routing() {

	s.Router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"startTime":    s.StartTime,
			"currDate":     time.Now().Format("2006-01-02 15:04"),
			"version":      "1.0",
			"EarningsTick": s.Tick.Scheduler["earnings"],
			"ApprovedTick": s.Tick.Scheduler["approved"],
			"RejectedTick": s.Tick.Scheduler["rejected"],
			"service":      s.ServiceName,
		})
	})

	s.Router.POST("/scheduler", func(c *gin.Context) {
		schedulerUpdate := Schedul{}

		if err := c.BindJSON(&schedulerUpdate); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		switch schedulerUpdate.Status {
		case "terminate":
			s.QuitRPC <- true
		case "stop":
			if s.Tick.Status == false {
				c.JSON(http.StatusOK, gin.H{
					"Error": "Schedulers Ticks is not defined",
					"Tick":  nil,
				})
				return
			}
			s.Tick.Stop(c)
		case "run":
			s.Tick.Update("earnings", schedulerUpdate.EarningsTick)
			s.Tick.Update("approved", schedulerUpdate.ApprovedTick)
			s.Tick.Update("rejected", schedulerUpdate.RejectedTick)
			s.Tick.Status = true
			go s.Tick.Run("earnings", c)
			go s.Tick.Run("approved", c)
			go s.Tick.Run("rejected", c)
		}
	})
}

func (t *Ticker) Update(schedulerType string, schedulerTick int) {
	if schedulerTick > 0 {
		t.Scheduler[schedulerType] = time.NewTicker(time.Duration(schedulerTick) * time.Second)
	}
}

func (t *Ticker) Stop(c *gin.Context) {
	t.Scheduler["earnings"].Stop()
	t.Scheduler["approved"].Stop()
	t.Scheduler["rejected"].Stop()

	c.JSON(200, gin.H{
		"SchedulerType": "earnings, approved, rejected",
		"Error":         nil,
		"Tick":          "STOP",
	})
}
