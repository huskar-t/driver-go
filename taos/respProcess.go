package taos

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

//mgmtProcessShellMsg[TSDB_MSG_TYPE_CONNECT] = mgmtProcessConnectMsg;
type processFunc func(header STaosHeader) error

var respProcessMap = map[byte]processFunc{
	TSDB_MSG_TYPE_CONNECT_RSP: STaosHeader.mgmtProcessConnectMsg,
}

type STaosRsp struct {
	code uint8
	more []byte
}

type SConnectRsp struct {
	AcctID    [TSDB_ACCT_LEN]byte
	Version   [TSDB_VERSION_LEN]byte
	WriteAuth byte
	SuperAuth byte
}

func (rpcProtocol *RpcProtocol) parseResp(data []byte) error {
	//	解析STaosHeader
	dataLen := len(data)
	if dataLen <= 0 {
		return errors.New("resp data is null")
	}
	var header STaosHeader
	header.parseVersionComp(data[0])
	header.parseTcpSpiEncrypt(data[1])
	header.TranID = binary.LittleEndian.Uint16(data[2:4])
	header.UID = binary.LittleEndian.Uint32(data[4:8])
	header.SourceID = binary.LittleEndian.Uint32(data[8:12])
	//设置DestID
	rpcProtocol.header.DestID = header.SourceID
	header.DestID = binary.LittleEndian.Uint32(data[12:16])
	for i := 0; i < TSDB_UNI_LEN; i++ {
		header.MeterID[i] = data[16+i]
	}
	header.Port = binary.LittleEndian.Uint16(data[40:42])
	header.Empty = data[42]
	header.MsgType = data[43]
	header.MsgLen = binary.BigEndian.Uint32(data[44:48])
	if header.Spi == 1 {
		header.Content = data[48 : dataLen-20]
	} else {
		header.Content = data[48:]
	}
	if len(header.Content) == 0 {
		return fmt.Errorf("mgmtProcessConnectMsg error: header.Content % x", header.Content)
	}
	code := int(header.Content[0])
	if code != 0 {
		return GetErrorStr(code)
	}
	processor, exist := respProcessMap[header.MsgType]
	if !exist {
		return fmt.Errorf("can not find processor, msgType = %d", header.MsgType)
	}
	return processor(header)
}
func (header *STaosHeader) parseVersionComp(data byte) {
	header.Version = data & 0xf
	header.Comp = (data >> 4) & 0xf
}

func (header *STaosHeader) parseTcpSpiEncrypt(data byte) {
	header.Tcp = data & 0x3
	header.Spi = (data >> 2) & 0x7
	header.Encrypt = (data >> 5) & 0x7
}

func (header *STaosHeader) mgmtProcessConnectMsg() error {
	data := header.Content[1:]
	buf := bytes.NewBuffer(data)
	obj := &SConnectRsp{}
	if err := binary.Read(buf, binary.BigEndian, obj); err != nil {
		return err
	}
	//暂不考虑版本号对应以及权限
	return nil
}
