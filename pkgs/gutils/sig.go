package gutils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
Listen Signal: Ctrl+C
*/
type CtrlCSignal struct {
	Sweeper func() error
}

func (that *CtrlCSignal) RegisterSweeper(f func() error) {
	that.Sweeper = f
}

func (that *CtrlCSignal) exitHandle(exitChan chan os.Signal) {
	for range exitChan {
		fmt.Println("\nExiting...")
		if that.Sweeper != nil {
			if err := that.Sweeper(); err != nil {
				fmt.Println(err)
			}
		}
		os.Exit(1)
	}
}

func (that *CtrlCSignal) ListenSignal() {
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt, syscall.SIGTERM)
	go that.exitHandle(exitChan)
}
