package main

import "grow.graphics/gd"

// TODO(calco): This is absolutely horrendous lmfao
func ToPackedByteArray(buf []byte) gd.PackedByteArray {
	byte_array := GL.PackedByteArray()
	for _, b := range buf {
		byte_array.Append(gd.Int(b))
	}
	return byte_array
}
