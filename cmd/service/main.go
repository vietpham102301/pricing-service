package main

import (
	"booking-service/server"
	"go.uber.org/zap"
)

func main() {
	s, err := server.NewServer()
	if err != nil {
		zap.S().Errorf("create server fail with err: %v", err)
		panic(err)
	}

	s.Init()

	err = s.Listen()
	if err != nil {
		zap.S().Errorf("start server fail with err: %v", err)
		panic(err)
	}
}
