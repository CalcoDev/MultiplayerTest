package main

import (
	"context"
	"log"
	"os"

	"github.com/calcodev/working"
	"grow.graphics/gd"
)

type DummyClient struct {
	gd.Class[DummyClient, gd.RefCounted] `gd:"DummyClient"`

	ClientId gd.Int `gd:"client_id"`
}

func (c *DummyClient) Initialize(id gd.Int) {
	c.ClientId = id
}

type ServerComponent struct {
	gd.Class[ServerComponent, gd.Node] `gd:"ServerComponent"`

	LogFile gd.Bool `gd:"log_file"`

	wrkServer *working.Server

	// SIGNALS
	OnStarted            gd.SignalAs[func()]       `gd:"on_started"`
	OnStopped            gd.SignalAs[func()]       `gd:"on_stopped"`
	OnClientConnected    gd.SignalAs[func(gd.Int)] `gd:"on_client_connected"`
	OnClientDisconnected gd.SignalAs[func(gd.Int)] `gd:"on_client_disconnected"`

	OnPacketReceived gd.SignalAs[func(gd.Int, gd.Int, gd.PackedByteArray)] `gd:"on_packet_received"`
}

func (s *ServerComponent) OnRegister(godot gd.Context) {
	godot.Register(gd.Enum[ServerComponent, int]{
		Name: "ServerState",
		Values: map[string]int{
			"ServerNone":     0,
			"ServerStarted":  1,
			"ServerStopping": 2,
			"ServerStopped":  3,
		},
	})
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

	s.wrkServer.OnStarted.Subscribe(func() {
		// TODO(calco): This should use it's own gd.Lifetime.
		GL.Callable(func() {
			s.OnStarted.Emit()
		}).CallDeferred()
	})
	s.wrkServer.OnStopped.Subscribe(func() {
		GL.Callable(func() {
			s.OnStopped.Emit()
		}).CallDeferred()
	})
	s.wrkServer.OnClientConnected.Subscribe(func(c *working.DummyClient) {
		GL.Callable(func() {
			s.OnClientConnected.Emit(gd.Int(c.ClientId))
		}).CallDeferred()
	})
	s.wrkServer.OnClientDiconnected.Subscribe(func(c *working.DummyClient) {
		GL.Callable(func() {
			s.OnClientDisconnected.Emit(gd.Int(c.ClientId))
		}).CallDeferred()
	})
	s.wrkServer.OnPacketReceived.Subscribe(func(c *working.DummyClient, n int, data []byte) {
		GL.Callable(func() {
			s.OnPacketReceived.Emit(gd.Int(c.ClientId), gd.Int(n), ToPackedByteArray(data))
		}).CallDeferred()
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

func (s *ServerComponent) HasClientId(clientId gd.Int) gd.Bool {
	_, exists := s.wrkServer.HasClientId(working.ClientId(clientId))
	return exists
}

// TODO(calco): Should send stuff be blocking or non blocking ?!?!?!?
func (s *ServerComponent) SendToClient(clientId gd.Int, data gd.PackedByteArray) {
	go s.wrkServer.SendToClient(working.ClientId(clientId), data.Bytes())
}

func (s *ServerComponent) Broadcast(data gd.PackedByteArray) {
	go s.wrkServer.Broadcast(data.Bytes())
}

func (s *ServerComponent) GetIP() gd.String {
	return GL.String(s.wrkServer.IP)
}

func (s *ServerComponent) GetPort() gd.Int {
	return gd.Int(s.wrkServer.Port)
}
func (s *ServerComponent) GetState() gd.Int {
	return gd.Int(s.wrkServer.State)
}

func (s *ServerComponent) GetClients() gd.ArrayOf[DummyClient] {
	dc := gd.NewArrayOf[DummyClient](GL)
	for _, c := range s.wrkServer.Clients {
		dc.Append(DummyClient{ClientId: gd.Int(c.ClientId)})
	}
	return dc
}

func (s *ServerComponent) GetOwner() gd.Int {
	return gd.Int(s.wrkServer.Owner)
}

func (s *ServerComponent) GetCurrClientId() gd.Int {
	return gd.Int(s.wrkServer.CurrClientId)
}
