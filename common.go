package main

import "encoding/binary"

func ByteArrayToInt(bytes []byte) uint64 {
	switch len(bytes) {
	case 1:
		return uint64(bytes[0])
	case 2:
		return uint64(binary.LittleEndian.Uint16(bytes))
	case 4:
		return uint64(binary.LittleEndian.Uint32(bytes))
	case 8:
		return uint64(binary.LittleEndian.Uint64(bytes))
	default:
		return 0
	}
}
