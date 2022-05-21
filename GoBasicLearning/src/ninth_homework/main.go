package main

import (
	"encoding/binary"
	"fmt"
)

/*
goim协议解码器
解包方式：length field based frame decoder
数据包结构：
package_length 		包长度		4bytes
header_length 		头长度		2bytes
protocol_version 	协议版本		2bytes
operation 			操作码		4bytes
sequence 			请求序号ID	4bytes
body 				包内容		package_length - header_length
*/

func main() {
	body := "hello, I am tianlq"
	data := encoder(body)
	decoder(data)
}

type dataPackage struct {
	packageLen  uint32
	headerLen   uint16
	protocolVer uint16
	operation   uint32
	sequence    uint32
	body        string
}

// 编码器
func encoder(body string) []byte {
	// 整个head的长度 4+2+2+4+4
	headerLen := 16
	// 数据包总长度 包头+包体
	packageLen := headerLen + len(body)
	result := make([]byte, packageLen)

	// 字节序：大端序
	binary.BigEndian.PutUint32(result[0:4], uint32(packageLen))
	binary.BigEndian.PutUint16(result[4:6], uint16(headerLen))
	protocolVer := 1
	binary.BigEndian.PutUint16(result[6:8], uint16(protocolVer))
	operation := 2
	binary.BigEndian.PutUint32(result[8:12], uint32(operation))
	sequence := 1
	binary.BigEndian.PutUint32(result[12:16], uint32(sequence))

	byteBody := []byte(body)
	copy(result[16:], byteBody)
	return result
}

// 解码器
func decoder(data []byte) {
	// 数据包长度应 ≥ 16
	if len(data) < 16 {
		println("data length not enough")
		return
	}
	// 解析 packageLen
	packageLen := binary.BigEndian.Uint32(data[0:4])
	// 解析 headerLen
	headerLen := binary.BigEndian.Uint16(data[4:6])
	// 解析 protocolVer
	protocolVer := binary.BigEndian.Uint16(data[6:8])
	// 解析 operation
	operation := binary.BigEndian.Uint32(data[8:12])
	// 解析 sequence
	sequence := binary.BigEndian.Uint32(data[12:16])

	// 解析包体
	body := string(data[16:])

	result := dataPackage{
		packageLen:  packageLen,
		headerLen:   headerLen,
		protocolVer: protocolVer,
		operation:   operation,
		sequence:    sequence,
		body:        body,
	}
	fmt.Printf("data: %+v\n", result)
}
