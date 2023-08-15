package gutils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
Listen Signal: Ctrl+C
*/
type CtrlCSignal struct{}

func (that *CtrlCSignal) exitHandle(exitChan chan os.Signal) {
	for {
		select {
		case <-exitChan:
			fmt.Println("Exiting...")
			os.Exit(1)
		default:
			time.Sleep(30 * time.Millisecond)
		}
	}
}

func (that *CtrlCSignal) ListenSignal() {
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt, syscall.SIGTERM)
	go that.exitHandle(exitChan)
}
