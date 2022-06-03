package clientSocket

import (
	"DiniSQL-client/Client/Type"
	"fmt"
	"io/ioutil"
	"net"
	"strconv"
	// "strings"
)

var ch = make(chan []byte, 5)
var ClientIP = "127.0.0.1"
var ClientPort = 8005

// initial connection to region server and send message
func ConnectToRegion(regionIP string, regionPort int, packet Type.Packet) (recPacket Type.Packet) {
	// fmt.Println("Entering")
	address := net.TCPAddr{
		IP:   net.ParseIP(regionIP),
		Port: regionPort,
	}
	conn, err := net.DialTCP("tcp4", nil, &address)
	// print("dial...\n")

	if err != nil {
		ret := fmt.Sprintln(err)
		recPacket = Type.Packet{Head: Type.PacketHead{P_Type: Type.Result, Op_Type: -1},
			Payload: []byte(ret)}
		return recPacket
	}
	defer conn.Close()
	// print("dial1...\n")
	var packetBuf = make([]byte, 0)
	// fmt.Printf("packet.Head.P_Type:%d\n", packet.Head.P_Type)
	// fmt.Printf("packet.Head.Op_Type:%d\n", packet.Head.Op_Type)
	// fmt.Printf("packet.Payload:%s\n", packet.Payload)
	packetBuf, err = packet.MarshalMsg(packetBuf)
	if err != nil {
		ret := fmt.Sprintln(err)
		recPacket = Type.Packet{Head: Type.PacketHead{P_Type: Type.Result, Op_Type: -1},
			Payload: []byte(ret)}
		return recPacket
	}
	// print("dial2...\n")
	_, err = conn.Write(packetBuf)
	// fmt.Println(packetBuf)
	if err != nil {
		ret := fmt.Sprintln(err)
		recPacket = Type.Packet{Head: Type.PacketHead{P_Type: Type.Result, Op_Type: -1},
			Payload: []byte(ret)}
		return recPacket
	}
	// print("dial3...\n")
	// print("listening...\n")
	var buf = make([]byte, 1024)
	conn.Read(buf)
	// print("read\n")
	_, err = recPacket.UnmarshalMsg(buf)
	if err != nil {
		ret := fmt.Sprintln(err)
		recPacket = Type.Packet{Head: Type.PacketHead{P_Type: Type.Result, Op_Type: -1},
			Payload: []byte(ret)}
		return recPacket
	}

	// fmt.Printf("p.Head.P_Type:%d\n", recPacket.Head.P_Type)
	// fmt.Printf("p.Head.Op_Type:%d\n", recPacket.Head.Op_Type)
	// fmt.Printf("p.Payload:%s\n", recPacket.Payload)
	// fmt.Print("hhhhhhhhhhh!!!")
	// fmt.Println(recPacket.Head.Op_Type)
	return recPacket

}

// listen
// input : IP and Port of client
func KeepListening(ClientIP string, ClientPort int) (receivedPacket Type.Packet) {
	// fmt.Printf("listening in %s...\n", ClientIP)
	// ClientPort := strings.Split(conn.LocalAddr().String(),":")[1]
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(ClientPort))
	defer listener.Close()
	if err != nil {
		ret := fmt.Sprintln(err)
		receivedPacket = Type.Packet{Head: Type.PacketHead{P_Type: Type.Result, Op_Type: -1},
			Payload: []byte(ret)}
		return receivedPacket
	}
	// for{
	conn, err := listener.Accept()
	if err != nil {
		ret := fmt.Sprintln(err)
		receivedPacket = Type.Packet{Head: Type.PacketHead{P_Type: Type.Result, Op_Type: -1},
			Payload: []byte(ret)}
		return receivedPacket
	}
	// fmt.Println("remote address:", conn.RemoteAddr())
	res, err := ioutil.ReadAll(conn)
	if err != nil {
		ret := fmt.Sprintln(err)
		receivedPacket = Type.Packet{Head: Type.PacketHead{P_Type: Type.Result, Op_Type: -1},
			Payload: []byte(ret)}
		return receivedPacket
	}
	res, err = receivedPacket.UnmarshalMsg(res)
	// ch <- p
	if err != nil {
		ret := fmt.Sprintln(err)
		receivedPacket = Type.Packet{Head: Type.PacketHead{P_Type: Type.Result, Op_Type: -1},
			Payload: []byte(ret)}
		return receivedPacket
	}

	// fmt.Printf("p.Head.P_Type:%d\n", receivedPacket.Head.P_Type)
	// fmt.Printf("p.Head.Op_Type:%d\n", receivedPacket.Head.Op_Type)
	// fmt.Printf("p.Payload:%s\n", receivedPacket.Payload)
	// }
	return receivedPacket

}

// func main() {
// 	var sql string
// 	// go KeepListening("127.0.0.1", 8006) // test locally
// 	for true {
// 		fmt.Scanln(&sql)
// 		// print(sql)
// 		sql := "insert into student2 values(1080100001,'name1',99);"
// 		p := Type.Packet{Head: Type.PacketHead{P_Type: Type.SQLOperation, Op_Type: Type.Insert},
// 			Payload: []byte(sql)}

// 		ConnectToRegion("127.0.0.1", 8006, p) //RegionIP + RegionPort
// 	}
// }

//client 8005 send -> listen
//region 8006
