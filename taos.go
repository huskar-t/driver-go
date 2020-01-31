package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//1字节     uint8_t
//2字节     uint16_t
//4字节     uint32_t
//8字节     uint64_t
type STaosHeader struct {
	Version  byte //4
	Comp     byte //4
	Tcp      byte //2
	Spi      byte //3
	Encrypt  byte //3
	TranID   uint16
	UID      uint32
	SourceID uint32
	DestID   uint32
	MeterID  [24]byte
	Port     uint16
	Empty    byte
	MsgType  uint8
	MsgLen   uint32 //185总数据长度 ?+48+len(auth)(4+16)=185 ?=185-48-20=117 185为空content 其他在此基础上加具体见taosmsg.h
	Content  []byte
}
type SConnectRsp struct {
	AcctID    [24]byte
	Version   [12]byte
	WriteAuth byte
	SuperAuth byte
}

//import _ "github.com/taosdata/driver-go/driver"
func main() {
	Content := []byte{
		0x00, 0x30, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x31, 0x2E, 0x36, 0x2E, 0x34, 0x2E, 0x34, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0x00,
		0x00, 0x00, 0x00}
	data := Content[1:]

	buf := bytes.NewBuffer(data)

	obj := &SConnectRsp{}

	if err := binary.Read(buf, binary.BigEndian, obj); err != nil {
		panic(err)
	}
	fmt.Println(string(obj.AcctID[:]))
	fmt.Println(string(obj.Version[:]))
	fmt.Println(obj.WriteAuth)
}

func parse(data byte) {
	fmt.Println(data&0x3, (data>>2)&0x7, (data>>5)&0x7)
}
func strDequote(z string) (string, int) {
	if z == "" {
		return z, 0
	}
	quote := z[0]
	if quote != '\'' && quote != '"' {
		return z, len(z)
	}
	b := []byte(z)
	j := 0
	for i := 1; i < len(b); i++ {
		if b[i] == quote {
			if b[i+1] == quote {
				j += 1
				b[j] = quote
			} else {
				return string(b), j
			}
		} else {
			j += 1
			b[j] = b[i]
		}
	}
	return string(b), j + 1
}
