package service

import (
	"os"
	"os/signal"

	"github.com/tosone/worklyzer/travel"
)

// Initialize ..
func Initialize() {
	go travel.Travel()

	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel, os.Interrupt)
	for {
		select {
		case <-signalChanel:
			return
		}
	}
}
