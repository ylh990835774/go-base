package protocol

import (
	"bytes"
	"encoding/binary"
)

const (
	ConstHeader         = "Mon"
	ConstHeaderLength   = 3
	ConstSaveDataLength = 4
)

// 封装包
func Packet(message []byte) []byte {
	return append(append([]byte(ConstHeader), IntToBytes(len(message))...), message...)
}

// 解包
func Unpack(buffer []byte, readerChannel chan []byte) []byte {
	length := len(buffer)

	var i int
	for i = 0; i < length; i = i + 1 {
		iHeaderLen := i + ConstHeaderLength
		iHeaderDataLen := iHeaderLen + ConstSaveDataLength

		if length < iHeaderDataLen {
			break
		}
		if string(buffer[i:iHeaderLen]) == ConstHeader {
			messageLength := BytesToInt(buffer[iHeaderLen:iHeaderDataLen])
			if length < iHeaderDataLen+messageLength {
				break
			}
			data := buffer[iHeaderDataLen : iHeaderDataLen+messageLength]
			readerChannel <- data

			i += ConstHeaderLength + ConstSaveDataLength + messageLength - 1
		}
	}

	if i == length {
		return make([]byte, 0)
	}

	return buffer[i:]
}

// 字节转为整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	_ = binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

// 整形转为字节
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}
