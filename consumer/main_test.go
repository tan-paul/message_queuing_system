package main

import (
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestSignals(t *testing.T) {
	flag := false
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// handeling interrupts or running consumer
	for {
		select {
		case sig := <-sigs:
			t.Log("signal :: ", sig)
			flag = true
			break
		default:
			t.Log("Sleeping for a sec as no signal found")
			time.Sleep(1 * time.Second)
		}
		if flag {
			break
		}
	}
}
