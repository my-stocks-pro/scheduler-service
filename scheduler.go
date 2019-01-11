package main

import (
	"os"
	"github.com/my-stocks-pro/earnings-scheduler/vendor.orig/github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Ticker struct {
	Status     bool
	Scheduler  map[string]*time.Ticker
	HTTPClient http.Client
	APIGateWay string
	Data       Schedul
}

type Scheduler struct {
	ServiceName string
	StartTime   string
	QuitOS      chan os.Signal
	QuitRPC     chan bool
	QuitTick    chan bool
	Router      *gin.Engine
	Server      *http.Server
	Tick        *Ticker
}

func NewScheduler() *Scheduler {

	router := gin.Default()
	return &Scheduler{
		ServiceName: "scheduler-service",
		StartTime:   time.Now().Format("2006-01-02 15:04"),
		QuitOS:      make(chan os.Signal),
		QuitRPC:     make(chan bool),
		QuitTick:    make(chan bool),
		Router:      router,
		Server: &http.Server{
			Addr:    ":9000",
			Handler: router,
		},
		Tick: NewTick(),
	}
}
