package main

import (
	"github.com/my-stocks-pro/earnings-scheduler/vendor.orig/github.com/gin-gonic/gin"
	"time"
)

func (s Scheduler) Routing() {

	s.Router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"service":  ServiceName,
			"currDate": time.Now().Format("2006-01-02 15:04"),
			"version":  "1.0",
		})
	})

	s.Router.GET("/", func(c *gin.Context) {
		go s.Run()
	})

	s.Router.GET("/stop", func(c *gin.Context) {
		s.QuitRPC <- true
	})

}
