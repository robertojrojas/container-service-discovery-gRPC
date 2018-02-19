package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	Version   string
	BuildTime string
)

func main() {

	errChan := make(chan error, 1)
	log.Printf("Greeter-Go: %s %s\n", Version, BuildTime)
	go func() {
		errChan <- startServer()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		case s := <-signalChan:
			log.Println(fmt.Sprintf("Captured %v. Exciting...", s))
			os.Exit(0)
		}
	}
}
