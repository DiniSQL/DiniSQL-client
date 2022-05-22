package Type

//go:generate msgp

type PacketHead struct {
	P_Type  int
	Op_Type int
	Spare   string
}

type Packet struct {
	Head     PacketHead
	Signal   bool
	Payload  []byte
	IPResult []byte
}
