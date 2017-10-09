package main

import (
	"time"
	"fmt"
	"os"
	"syscall"
	"os/signal"
	_ "sync"
)

func main() {
	defer func() {
		fmt.Println("\033[?25h")
	}()
	fmt.Print("\033[?25l")
	doneChan := make(chan bool, 1)
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-stopChan
		doneChan <- true
	}()

	tickChan := time.NewTicker(500 * time.Millisecond).C
	for {
		select {
		case <- tickChan:
			ReadBitBay()
		case <- doneChan:
			fmt.Println("\nFinish")
			return
		}
	}
}
