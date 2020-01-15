package taos

func (rpcProtocol *RpcProtocol) tscBuildConnectMsg(Content []byte) ([]byte, error) {
	connectMsg := make([]byte, 0xb9)
	copy(connectMsg, rpcProtocol.BaseMsg)
	err := rpcProtocol.addDigest(connectMsg)
	if err != nil {
		return nil, err
	}
	return connectMsg, nil
}

func (rpcProtocol *RpcProtocol) tscBuildHeartBeatMsg(Content []byte) ([]byte, error) {
	return nil, nil
}
