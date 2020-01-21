package taos

type SConnectRsp struct {
	AcctID    [TSDB_ACCT_LEN]byte
	Version   [TSDB_VERSION_LEN]byte
	WriteAuth byte
	SuperAuth byte
}

func (RpcProtocol *RpcProtocol) parseResp(data []byte) {

}
func (rpcProtocol *RpcProtocol) tscProcessConnectRsp(data []byte) ([]byte, error) {

}
