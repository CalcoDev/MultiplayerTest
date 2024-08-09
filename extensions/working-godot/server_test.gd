extends Node

@export var server: ServerComponent = null
@export var client: ClientComponent = null


func _enter_tree() -> void:
    # print(server.GetClients())
    pass
    # server.OnStarted.connect(func():
    #     print("SERVER STARTED")
    # )
    # server.OnStopped.connect(func():
    #     print("SERVER STOPPED")
    # )
    # server.OnClientConnected.connect(func(clientId: int):
    #     print("CLIENT ID ", clientId, " CONNECTED")
    # )
    # server.OnPacketReceived.connect(func(clientId: int, _n: int, data: PackedByteArray):
    #     print("CLIENT ID ", clientId, " SEND DATA: ", data.get_string_from_utf8())
    # )

    # print(DummyClient.new().Clien)
    # print(ClientComponent.ClientState.ClientNone)
    # print(ClientComponent.ClientState.ClientStarted)
    # print(ClientComponent.ClientState.ClientStopping)
    # print(ClientComponent.ClientState.ClientStopped)

func _ready() -> void:
    pass
#     server.InitServer()
#     ($"../Start" as Button).pressed.connect(func():
#         server.Start())
#     ($"../Stop" as Button).pressed.connect(func():
#         server.Stop())

# func _exit_tree() -> void:
#     server.Stop()
