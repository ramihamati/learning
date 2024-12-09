package network

type NetworkAddress struct {
	Value string
}

func (n NetworkAddress) Equals(addr NetworkAddress) bool {
	return n.Value == addr.Value
}
