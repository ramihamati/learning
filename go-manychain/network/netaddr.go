package network

type NetAddr struct {
	Value string
}

func (n NetAddr) Equals(addr NetAddr) bool {
	return n.Value == addr.Value
}
