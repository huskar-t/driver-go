package main

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
)

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

//import _ "github.com/taosdata/driver-go/driver"
func main() {
	a := STaosHeader{
		Version:  0x1,
		Comp:     0x0,
		Tcp:      0x0,
		Spi:      0x1,
		Encrypt:  0x0,
		TranID:   0x2efe,
		UID:      0x18ae358,
		SourceID: 0x1000000,
		DestID:   0x0,
		MeterID:  [24]byte{'r', 'o', 'o', 't'},
		Port:     0x0,
		Empty:    0x0,
		MsgType:  0x1f,
		MsgLen:   0xb9,
		Content: []byte{},
	}
	data := make([]byte,a.MsgLen)
	fmt.Println(int(a.MsgLen))
	data[0] = unionVersionComp(&a)
	data[1] = unionTcpSpiEncrypt(&a)
	binary.LittleEndian.PutUint16(data[2:4], a.TranID)
	binary.LittleEndian.PutUint32(data[4:8], a.UID)
	binary.LittleEndian.PutUint32(data[8:12], a.SourceID)
	binary.LittleEndian.PutUint32(data[12:16], a.DestID)
	for i := 0; i < 24; i++ {
		data[16+i] = a.MeterID[i]
	}
	binary.LittleEndian.PutUint16(data[40:42], a.Port)
	data[42] = a.Empty
	data[43] = a.MsgType
	binary.BigEndian.PutUint32(data[44:48], a.MsgLen)
	fmt.Printf("% x\n", data)

	dataLen := len(data)
	digestBuf := data[dataLen-20:]
	//now := time.Now()
	//unix := now.Unix()
	binary.BigEndian.PutUint32(digestBuf, uint32(1578241357))
	fmt.Printf("% x\n", digestBuf)
	h := md5.New()
	h.Write([]byte("root"))
	encodeUser := h.Sum(nil)
	h2 := md5.New()
	h2.Write(encodeUser)
	h2.Write(data[:dataLen-16])
	h2.Write(encodeUser)
	fmt.Printf("% x\n", h2.Sum(nil))

}

func unionVersionComp(d *STaosHeader) byte {
	low := d.Version & 0xf
	high := d.Comp & 0xf
	return high<<4 + low
}

func unionTcpSpiEncrypt(d *STaosHeader) byte {
	low := d.Tcp & 3
	mid := d.Spi & 7
	high := d.Encrypt & 7
	return low + mid<<2 + high<<5
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
