package network

type RemotePeer struct {
	Address   NetAddr
	Transport ITransport
}

func NewRemotePeer(transport ITransport, address NetAddr) *RemotePeer {
	return &RemotePeer{
		Transport: transport,
		Address:   address,
	}
}
