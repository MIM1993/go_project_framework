package internal

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func ListeningExitSignal(){
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-signalChan
	logrus.Infof("receive signal: %v", sig)
	os.Exit(0)
}