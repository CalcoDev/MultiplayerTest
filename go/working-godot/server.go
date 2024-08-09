package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/calcodev/working"
	"grow.graphics/gd"
)

type ServerComponent struct {
	gd.Class[ServerComponent, gd.Node] `gd:ServerComponent`

	LogFile gd.Bool

	wrkServer *working.Server
}

func (s *ServerComponent) GetAddress() gd.String {
	return GL.String(s.wrkServer.GetAddress())
}

func (s *ServerComponent) InitServer() {
	if s.LogFile {
		file, err := os.OpenFile(CLIENT_LOGS, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("ERROR: Failed opening %s: [%q]", SERVER_LOGS, err)
		}
		defer file.Close()
		log.SetOutput(file)
	}

	ctx, cancel := context.WithCancel(context.Background())

	s.wrkServer = working.NewServer(ctx, cancel, SERVER_IP, SERVER_PORT)
	s.wrkServer.OnPacketReceived.Subscribe(func(c *working.DummyClient, n int, data []byte) {
		fmt.Println("RECV: ", data, " FROM ", c.ClientId)
	})
}

func (s *ServerComponent) Start() {
	go s.handleStart()
}

func (s *ServerComponent) handleStart() {
	s.wrkServer.Start()
}

func (s *ServerComponent) Stop() {
	go s.wrkServer.Stop()
}
