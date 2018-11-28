package main

import (
	"os"
	"github.com/my-stocks-pro/earnings-scheduler/vendor.orig/github.com/gin-gonic/gin"
	"net/http"
	"github.com/my-stocks-pro/earnings-scheduler/config"
)

type Scheduler struct {
	QuitOS     chan os.Signal
	QuitRPC    chan bool
	Router     *gin.Engine
	Server     *http.Server
	Config     *config.TypeConfig
	Services   *map[string][]string
	HTTPClient http.Client
}

func NewScheduler() Scheduler {

	router := gin.Default()
	return Scheduler{
		QuitOS:  make(chan os.Signal),
		QuitRPC: make(chan bool),
		Router:  router,
		Server: &http.Server{
			Addr:    ":9000",
			Handler: router,
		},
		HTTPClient: http.Client{},
	}
}
