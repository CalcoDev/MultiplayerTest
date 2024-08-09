extends Node

@export var gonode: GoNode = null

func _enter_tree() -> void:
    process_priority = 999

func _ready() -> void:
    get_tree().create_timer(0.2).timeout.connect(func():
        gonode.SayHiGo()
        gonode.SayHiGodot()
        print(gonode.ExportedInt))
