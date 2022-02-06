// 协议处理，处理封包和解包
package protocol

import (
	"bytes"
	"encoding/binary"
)

const (
	ConstHeader         = "www.acfun.cn"
	ConstHeaderLength   = 12
	ConstSaveDataLength = 4
)

// 封包
func Packet(message []byte) []byte {
	// header + body length + message
	return append(append([]byte(ConstHeader), IntToBytes(len(message))...), message...)
}

// 解包
func Unpack(buffer []byte, readerChannel chan []byte) []byte {
	length := len(buffer)
	var i int
	for i = 0; i < length; i++ {
		if length < i+ConstHeaderLength+ConstSaveDataLength {
			// 长度为最小头部信息长
			break
		}
		if string(buffer[i:i+ConstHeaderLength]) == ConstHeader {
			// 是否为约定的头部信息
			messageLength := BytesToInt(buffer[i+ConstHeaderLength : i+ConstHeaderLength+ConstSaveDataLength])
			if length < i+ConstHeaderLength+ConstSaveDataLength+messageLength {
				break
			}
			data := buffer[i+ConstHeaderLength+ConstSaveDataLength : i+ConstHeaderLength+ConstSaveDataLength+messageLength]
			readerChannel <- data

			i += ConstHeaderLength + ConstSaveDataLength + messageLength - 1 // end index
		}
	}
	if i == length {
		// 没有找到约定的头部信息， return empty []byte
		return make([]byte, 0)
	}
	return buffer[i:] // return message

}

// 整型转换成字节 int32 4个字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// 字节转换成整型
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}
