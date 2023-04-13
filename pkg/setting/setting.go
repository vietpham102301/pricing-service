package setting

import (
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func WaitOSSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	s := <-c
	zap.S().Infof("Receive os.Signal: %s", s.String())
}
