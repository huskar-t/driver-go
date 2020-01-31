package taos

func (rpcProtocol *RpcProtocol) tscBuildConnectMsg(Content []byte) ([]byte, error) {
	//创建连接请求
	return rpcProtocol.buildMsg()
}

func (rpcProtocol *RpcProtocol) tscBuildHeartBeatMsg(Content []byte) ([]byte, error) {
	return nil, nil
}
