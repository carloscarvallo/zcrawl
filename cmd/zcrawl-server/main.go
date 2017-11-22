package main

import (
	"github.com/matiasinsaurralde/zcrawl-platform/server"
)

const (
	defaultAddr = "localhost:9999"
)

func main() {
	settings := &server.Settings{
		ListenAddr: defaultAddr,
	}
	s := server.New(settings)
	err := s.Start()
	if err != nil {
		panic(err)
	}
}