package Type

//go:generate msgp

type PacketHead struct {
	P_Type  int
	Op_Type int
}

type Packet struct {
	Head    PacketHead
	Payload []byte
}

type MasterClientPacket struct {
	Head      PacketHead
	Signal    bool
	SQLResult []byte
	IPResult  []byte
}

