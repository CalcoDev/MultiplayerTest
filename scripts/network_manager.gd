extends Node

var udp: PacketPeerUDP
var prev_packet_count: int = 0

func _ready() -> void:
    udp = PacketPeerUDP.new()
    var error_code = udp.set_dest_address("127.0.0.1", 25565)
    if error_code != OK:
        push_error("ERROR: Failed to bind udp peer.")
    print("CONNECTED TO UDP PEER.")
    prev_packet_count = udp.get_available_packet_count()

    var byte_array = PackedByteArray()
    byte_array.append_array("Hello world!".to_ascii_buffer())
    udp.put_packet(byte_array)
    print("SENT STUFF")

# func _process(delta: float) -> void:
#     var packet_count = udp.get_available_packet_count()
#     if packet_count != prev_packet_count:
#         print("PACKETS CHANGED?")
#         print(udp.get_packet())
#         var packet_error = udp.get_packet_error()
#         if packet_error != OK:
#             push_warning("ERROR: Couldn't recieve packet!")
