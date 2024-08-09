package main

import (
	"log"
	"os"

	"github.com/calcodev/working"
	"grow.graphics/gd"
)

type ClientComponent struct {
	gd.Class[ClientComponent, gd.Node] `gd:ClientComponent`

	LogFile gd.Bool

	wrkClient *working.Client

	// SIGNALS
	OnStarted      gd.SignalAs[func()]          `gd:on_started`
	OnStopped      gd.SignalAs[func()]          `gd:on_stopped`
	OnConnected    gd.SignalAs[func(gd.String)] `gd:on_client_connected`
	OnDisconnected gd.SignalAs[func(gd.String)] `gd:on_client_disconnected`

	OnPacketReceived gd.SignalAs[func(gd.Int, gd.PackedByteArray)] `gd:on_packet_received`
}

func (c *ClientComponent) OnRegister(godot gd.Context) {
	godot.Register(gd.Enum[ClientComponent, int]{
		Name: "ClientState",
		Values: map[string]int{
			"ClientNone":     0,
			"ClientStarted":  1,
			"ClientStopping": 2,
			"ClientStopped":  3,
		},
	})
}

func (c *ClientComponent) GetAddress() gd.String {
	return GL.String(c.wrkClient.GetAddress())
}

func (c *ClientComponent) InitClient() {
	if c.LogFile {
		file, err := os.OpenFile(CLIENT_LOGS, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("ERROR: Failed opening %s: [%q]", SERVER_LOGS, err)
		}
		defer file.Close()
		log.SetOutput(file)
	}

	c.wrkClient.OnStarted.Subscribe(func() {
		// TODO(calco): This should use it's own gd.Lifetime.
		GL.Callable(func() {
			c.OnStarted.Emit()
		}).CallDeferred()
	})
	c.wrkClient.OnStopped.Subscribe(func() {
		GL.Callable(func() {
			c.OnStopped.Emit()
		}).CallDeferred()
	})
	c.wrkClient.OnDiconnected.Subscribe(func(s *working.DummyServer) {
		GL.Callable(func() {
			c.OnConnected.Emit(GL.String(s.Address.String()))
		}).CallDeferred()
	})
	c.wrkClient.OnDiconnected.Subscribe(func(s *working.DummyServer) {
		GL.Callable(func() {
			c.OnDisconnected.Emit(GL.String(s.Address.String()))
		}).CallDeferred()
	})
	c.wrkClient.OnPacketReceived.Subscribe(func(n int, data []byte) {
		GL.Callable(func() {
			c.OnPacketReceived.Emit(gd.Int(n), ToPackedByteArray(data))
		}).CallDeferred()
	})
}

func (c *ClientComponent) Start(address gd.String) {
	go c.handleStart(address.String())
}

func (c *ClientComponent) handleStart(address string) {
	c.wrkClient.Start(address)
}

func (c *ClientComponent) Stop() {
	go c.wrkClient.Stop()
}

// TODO(calco): Should send stuff be blocking or non blocking ?!?!?!?
func (c *ClientComponent) Send(data gd.PackedByteArray) {
	go c.wrkClient.Send(data.Bytes())
}

func (c *ClientComponent) GetState() gd.Int {
	return gd.Int(c.wrkClient.State)
}
