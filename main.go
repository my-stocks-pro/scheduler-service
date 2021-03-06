package main

import (
	"os/signal"
	"os"
	"log"
	"time"
	"net/http"
	"context"
	"fmt"
)

const (
	//ServiceName = "scheduler-service"
	//EarningsTick = 5
	//ApprovedTick = 10
	//RejectedTick = 15
	//API = "127.0.0.1:9001/scheduler"
)

func main() {
	scheduler := NewScheduler()

	fmt.Println(fmt.Sprintf("%s INITIALIZE", scheduler.ServiceName))

	scheduler.Routing()

	go func() {
		if err := scheduler.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	signal.Notify(scheduler.QuitOS, os.Interrupt)
	select {
	case <-scheduler.QuitOS:
		log.Println("Shutdown Server by OS signal...")
	case <-scheduler.QuitRPC:
		log.Println("Shutdown Server by RPC signal...")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := scheduler.Server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}
