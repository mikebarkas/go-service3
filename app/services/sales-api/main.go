package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"go.uber.org/automaxprocs/maxprocs"
)

var build = "develop"

func main() {

	if _, err := maxprocs.Set(); err != nil {
		fmt.Println("maxprocs: %w", err)
		os.Exit(1)
	}
	g := runtime.GOMAXPROCS(0)

	log.Printf("Starting service build[%s] CPI[%d]", build, g)
	defer log.Println("Service stopped")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	log.Println("stopping service")
}
