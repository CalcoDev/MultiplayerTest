package main

import (
	"fmt"

	"grow.graphics/gd"
	"grow.graphics/gd/gdextension"
)

const SERVER_IP = "127.0.0.1"
const SERVER_PORT = 25565
const SERVER_LOGS = "server.log"

const CLIENT_LOGS = "client.log"

var GL gd.Lifetime

type GoNode struct {
	gd.Class[GoNode, gd.Node] `gd:"GoNode"`

	ExportedInt   gd.Int
	privateString gd.String
}

func (node *GoNode) Ready() {
	node.privateString = GL.String("Ooga booga godot code.")
}

func (node *GoNode) SayHiGo() {
	fmt.Println("Hello world Go! | Private string: ", node.privateString.String())
}

func (node *GoNode) SayHiGodot() {
	GL.Print(GL.Variant(node.privateString))
}

func main() {
	godot, ok := gdextension.Link()
	if !ok {
		return
	}
	GL = godot
	gd.Register[GoNode](godot)
	gd.Register[DummyClient](godot)
	gd.Register[ServerComponent](godot)
	gd.Register[ClientComponent](godot)
}
