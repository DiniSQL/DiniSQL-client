package Type

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Packet) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Head":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "Head")
				return
			}
			for zb0002 > 0 {
				zb0002--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					err = msgp.WrapError(err, "Head")
					return
				}
				switch msgp.UnsafeString(field) {
				case "P_Type":
					z.Head.P_Type, err = dc.ReadInt()
					if err != nil {
						err = msgp.WrapError(err, "Head", "P_Type")
						return
					}
				case "Op_Type":
					z.Head.Op_Type, err = dc.ReadInt()
					if err != nil {
						err = msgp.WrapError(err, "Head", "Op_Type")
						return
					}
				case "Spare":
					z.Head.Spare, err = dc.ReadString()
					if err != nil {
						err = msgp.WrapError(err, "Head", "Spare")
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						err = msgp.WrapError(err, "Head")
						return
					}
				}
			}
		case "Signal":
			z.Signal, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "Signal")
				return
			}
		case "Payload":
			z.Payload, err = dc.ReadBytes(z.Payload)
			if err != nil {
				err = msgp.WrapError(err, "Payload")
				return
			}
		case "IPResult":
			z.IPResult, err = dc.ReadBytes(z.IPResult)
			if err != nil {
				err = msgp.WrapError(err, "IPResult")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Packet) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "Head"
	err = en.Append(0x84, 0xa4, 0x48, 0x65, 0x61, 0x64)
	if err != nil {
		return
	}
	// map header, size 3
	// write "P_Type"
	err = en.Append(0x83, 0xa6, 0x50, 0x5f, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Head.P_Type)
	if err != nil {
		err = msgp.WrapError(err, "Head", "P_Type")
		return
	}
	// write "Op_Type"
	err = en.Append(0xa7, 0x4f, 0x70, 0x5f, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Head.Op_Type)
	if err != nil {
		err = msgp.WrapError(err, "Head", "Op_Type")
		return
	}
	// write "Spare"
	err = en.Append(0xa5, 0x53, 0x70, 0x61, 0x72, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Head.Spare)
	if err != nil {
		err = msgp.WrapError(err, "Head", "Spare")
		return
	}
	// write "Signal"
	err = en.Append(0xa6, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteBool(z.Signal)
	if err != nil {
		err = msgp.WrapError(err, "Signal")
		return
	}
	// write "Payload"
	err = en.Append(0xa7, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Payload)
	if err != nil {
		err = msgp.WrapError(err, "Payload")
		return
	}
	// write "IPResult"
	err = en.Append(0xa8, 0x49, 0x50, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.IPResult)
	if err != nil {
		err = msgp.WrapError(err, "IPResult")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Packet) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Head"
	o = append(o, 0x84, 0xa4, 0x48, 0x65, 0x61, 0x64)
	// map header, size 3
	// string "P_Type"
	o = append(o, 0x83, 0xa6, 0x50, 0x5f, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendInt(o, z.Head.P_Type)
	// string "Op_Type"
	o = append(o, 0xa7, 0x4f, 0x70, 0x5f, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendInt(o, z.Head.Op_Type)
	// string "Spare"
	o = append(o, 0xa5, 0x53, 0x70, 0x61, 0x72, 0x65)
	o = msgp.AppendString(o, z.Head.Spare)
	// string "Signal"
	o = append(o, 0xa6, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c)
	o = msgp.AppendBool(o, z.Signal)
	// string "Payload"
	o = append(o, 0xa7, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64)
	o = msgp.AppendBytes(o, z.Payload)
	// string "IPResult"
	o = append(o, 0xa8, 0x49, 0x50, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74)
	o = msgp.AppendBytes(o, z.IPResult)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Packet) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Head":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Head")
				return
			}
			for zb0002 > 0 {
				zb0002--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					err = msgp.WrapError(err, "Head")
					return
				}
				switch msgp.UnsafeString(field) {
				case "P_Type":
					z.Head.P_Type, bts, err = msgp.ReadIntBytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "Head", "P_Type")
						return
					}
				case "Op_Type":
					z.Head.Op_Type, bts, err = msgp.ReadIntBytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "Head", "Op_Type")
						return
					}
				case "Spare":
					z.Head.Spare, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "Head", "Spare")
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						err = msgp.WrapError(err, "Head")
						return
					}
				}
			}
		case "Signal":
			z.Signal, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Signal")
				return
			}
		case "Payload":
			z.Payload, bts, err = msgp.ReadBytesBytes(bts, z.Payload)
			if err != nil {
				err = msgp.WrapError(err, "Payload")
				return
			}
		case "IPResult":
			z.IPResult, bts, err = msgp.ReadBytesBytes(bts, z.IPResult)
			if err != nil {
				err = msgp.WrapError(err, "IPResult")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Packet) Msgsize() (s int) {
	s = 1 + 5 + 1 + 7 + msgp.IntSize + 8 + msgp.IntSize + 6 + msgp.StringPrefixSize + len(z.Head.Spare) + 7 + msgp.BoolSize + 8 + msgp.BytesPrefixSize + len(z.Payload) + 9 + msgp.BytesPrefixSize + len(z.IPResult)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *PacketHead) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "P_Type":
			z.P_Type, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "P_Type")
				return
			}
		case "Op_Type":
			z.Op_Type, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "Op_Type")
				return
			}
		case "Spare":
			z.Spare, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Spare")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z PacketHead) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "P_Type"
	err = en.Append(0x83, 0xa6, 0x50, 0x5f, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt(z.P_Type)
	if err != nil {
		err = msgp.WrapError(err, "P_Type")
		return
	}
	// write "Op_Type"
	err = en.Append(0xa7, 0x4f, 0x70, 0x5f, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Op_Type)
	if err != nil {
		err = msgp.WrapError(err, "Op_Type")
		return
	}
	// write "Spare"
	err = en.Append(0xa5, 0x53, 0x70, 0x61, 0x72, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Spare)
	if err != nil {
		err = msgp.WrapError(err, "Spare")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z PacketHead) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "P_Type"
	o = append(o, 0x83, 0xa6, 0x50, 0x5f, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendInt(o, z.P_Type)
	// string "Op_Type"
	o = append(o, 0xa7, 0x4f, 0x70, 0x5f, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendInt(o, z.Op_Type)
	// string "Spare"
	o = append(o, 0xa5, 0x53, 0x70, 0x61, 0x72, 0x65)
	o = msgp.AppendString(o, z.Spare)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PacketHead) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "P_Type":
			z.P_Type, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "P_Type")
				return
			}
		case "Op_Type":
			z.Op_Type, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Op_Type")
				return
			}
		case "Spare":
			z.Spare, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Spare")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z PacketHead) Msgsize() (s int) {
	s = 1 + 7 + msgp.IntSize + 8 + msgp.IntSize + 6 + msgp.StringPrefixSize + len(z.Spare)
	return
}
