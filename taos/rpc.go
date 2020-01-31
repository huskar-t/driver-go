package taos

import (
	"crypto/md5"
	"encoding/binary"
	"errors"
	"math/rand"
	"sync"
	"time"
)

type protocolBuilder func(RpcProtocol RpcProtocol, Content []byte) ([]byte, error)

var reqProtocolMap = map[int]protocolBuilder{
	TSDB_SQL_CONNECT: RpcProtocol.tscBuildConnectMsg,
	TSDB_SQL_HB:      RpcProtocol.tscBuildHeartBeatMsg,
}

type RpcProtocol struct {
	sync.RWMutex
	EncodePassword [TSDB_AUTH_LEN]byte //16
	DB             [TSDB_METER_ID_LEN]byte
	header         *STaosHeader
}
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
	MeterID  [TSDB_UNI_LEN]byte
	Port     uint16
	Empty    byte
	MsgType  uint8
	MsgLen   uint32 //185初始总数据长度  48基础线程信息+117数据库信息+4时间戳+16位秘钥
	Content  []byte
}

func (rpcProtocol *RpcProtocol) GetReqMsg(MsgType int, content []byte) ([]byte, error) {
	rpcProtocol.Lock()
	defer rpcProtocol.Unlock()
	builder := reqProtocolMap[MsgType]
	return builder(*rpcProtocol, content)
}

func (rpcProtocol *RpcProtocol) buildMsg() ([]byte, error) {
	//请求字节码涉及到高低位操作无法直接用 binary包直接对齐
	if rpcProtocol.header == nil {
		return nil, errors.New("STaosHeader is nil")
	}
	contentLen := len(rpcProtocol.header.Content)
	rpcProtocol.header.TranID += 1
	rpcProtocol.header.MsgLen = uint32(48 + TSDB_METER_ID_LEN + contentLen + 4 + 16)
	//48基础信息+117数据库信息+4时间戳+16位秘钥 = 185
	var data = make([]byte, rpcProtocol.header.MsgLen)
	var header = rpcProtocol.header
	data[0] = rpcProtocol.unionVersionComp()
	data[1] = rpcProtocol.unionTcpSpiEncrypt()
	binary.LittleEndian.PutUint16(data[2:4], header.TranID)
	binary.LittleEndian.PutUint32(data[4:8], header.UID)
	binary.LittleEndian.PutUint32(data[8:12], header.SourceID)
	binary.LittleEndian.PutUint32(data[12:16], header.DestID)
	for i := 0; i < TSDB_UNI_LEN; i++ {
		data[16+i] = header.MeterID[i]
	}
	binary.LittleEndian.PutUint16(data[40:42], header.Port)
	data[42] = header.Empty
	data[43] = header.MsgType
	binary.BigEndian.PutUint32(data[44:48], header.MsgLen)
	for i := 0; i < 117; i++ {
		data[48+i] = rpcProtocol.DB[i]
	}
	for i := 0; i < contentLen; i++ {
		data[165+i] = rpcProtocol.header.Content[i]
	}
	err := rpcProtocol.addDigest(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (rpcProtocol *RpcProtocol) unionVersionComp() byte {
	low := rpcProtocol.header.Version & 0xf
	high := rpcProtocol.header.Comp & 0xf
	return high<<4 + low
}

func (rpcProtocol *RpcProtocol) unionTcpSpiEncrypt() byte {
	low := rpcProtocol.header.Tcp & 0x3
	mid := rpcProtocol.header.Spi & 0x7
	high := rpcProtocol.header.Encrypt & 0x7
	return low + mid<<2 + high<<5
}

func (rpcProtocol *RpcProtocol) addDigest(data []byte) error {
	//calculate auth
	dataLen := len(data)
	if dataLen < 20 {
		return errors.New("date length less than 20")
	}
	digestBuf := data[dataLen-20:]
	now := time.Now()
	unix := now.Unix()
	binary.BigEndian.PutUint32(digestBuf, uint32(unix))
	h := md5.New()
	h.Write(rpcProtocol.EncodePassword[:])
	h.Write(data[:dataLen-16])
	h.Write(rpcProtocol.EncodePassword[:])
	auth := h.Sum(nil)
	if len(auth) != 16 {
		return errors.New("generate digest error")
	}
	for i := 0; i < 16; i++ {
		digestBuf[dataLen-16+i] = auth[i]
	}
	return nil
}

func NewRpcProtocol(user, db string, encodePassword []byte) *RpcProtocol {
	var rpcProtocol RpcProtocol
	meterID := [24]byte{}
	for i := 0; i < 24; i++ {
		meterID[i] = user[i]
	}
	initHeader := STaosHeader{
		Version:  0x1,
		Comp:     0x0,
		Tcp:      0x0,
		Spi:      0x1,
		Encrypt:  0x0,
		TranID:   0x0,
		UID:      rand.Uint32(),
		SourceID: 0x1000000,
		DestID:   0x0,
		MeterID:  meterID,
		Port:     0x0,
		Empty:    0x0,
		MsgType:  TSDB_MSG_TYPE_CONNECT, //预先设置成TSDB_MSG_TYPE_CONNECT
		//48基础线程信息+117数据库信息+4时间戳+16位秘钥
		MsgLen:  48 + TSDB_METER_ID_LEN + 4 + 16,
		Content: []byte{},
	}
	for i := 0; i < TSDB_AUTH_LEN; i++ {
		rpcProtocol.EncodePassword[i] = encodePassword[i]
	}
	rpcProtocol.DB = [117]byte{}
	for i := range db {
		rpcProtocol.DB[i] = db[i]
	}
	rpcProtocol.header = &initHeader
	return &rpcProtocol
}
