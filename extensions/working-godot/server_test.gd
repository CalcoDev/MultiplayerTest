extends Node

@export var server: ServerComponent = null

func _ready() -> void:
    server.InitServer()
    server.Start()

    get_tree().create_timer(60).timeout.connect(
    func():
        server.Stop()
    )

func _exit_tree() -> void:
    server.Stop()