package main

import (
	"DiniSQL-client/Client/Type"
	"fmt"
	// "io/ioutil"
	"log"
	"net"
	// "strconv"
	// "strings"
)

var ch = make(chan []byte, 5)
var ClientIP = "127.0.0.1"
var ClientPort = 8005

// initial connection to region server and send message
func ConnectToRegion(regionIP string, regionPort int, packet Type.Packet) (recPacket Type.Packet) {

	address := net.TCPAddr{
		IP:   net.ParseIP(regionIP),
		Port: regionPort,
	}
	conn, err := net.DialTCP("tcp4", nil, &address)

	defer conn.Close()
	if err != nil {
		log.Fatal(err) // Println + os.Exit(1)
		return
	}
	var packetBuf = make([]byte, 0)
	fmt.Printf("packet.Head.P_Type:%d\n", packet.Head.P_Type)
	fmt.Printf("packet.Head.Op_Type:%d\n", packet.Head.Op_Type)
	fmt.Printf("packet.Payload:%s\n", packet.Payload)
	packetBuf, err = packet.MarshalMsg(packetBuf)
	if err != nil {
		log.Fatal(err) // Println + os.Exit(1)
		return
	}
	_, err1 := conn.Write(packetBuf)
	fmt.Println(packetBuf)
	if err1 != nil {
		log.Println(err)
		conn.Close()
		return
	}
	print("listening...\n")
	var buf = make([]byte, 1024)
	conn.Read(buf)
	_, err = recPacket.UnmarshalMsg(buf)
	// res, err1 := ioutil.ReadAll(conn)
	// res, err1 = recPacket.UnmarshalMsg(res)
	// // ch <- p
	// if err1 != nil {
	// 	log.Println(err)
	// 	conn.Close()
	// }

	fmt.Printf("p.Head.P_Type:%d\n", recPacket.Head.P_Type)
	fmt.Printf("p.Head.Op_Type:%d\n", recPacket.Head.Op_Type)
	fmt.Printf("p.Payload:%s\n", recPacket.Payload)

	return

}

// listen
// input : IP and Port of client
// func KeepListening(conn net.Conn) (receivedPacket Type.Packet) {
// 	fmt.Printf("listening in %s...\n", conn.LocalAddr())
// 	ClientPort := strings.Split(conn.LocalAddr().String(),":")[1]
// 	listener, err := net.Listen("tcp", ":"+ClientPort)
// 	defer listener.Close()
// 	if err != nil {
// 		log.Fatal(err) // Println + os.Exit(1)
// 	}

// 	conn, err := listener.Accept()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("remote address:", conn.RemoteAddr())
// 	res, err1 := ioutil.ReadAll(conn)
// 	if err1 != nil {
// 		log.Println(err)
// 		conn.Close()
// 	}
	// res, err1 = receivedPacket.UnmarshalMsg(res)
	// // ch <- p
	// if err1 != nil {
	// 	log.Println(err)
	// 	conn.Close()
	// }

	// fmt.Printf("p.Head.P_Type:%d\n", receivedPacket.Head.P_Type)
	// fmt.Printf("p.Head.Op_Type:%d\n", receivedPacket.Head.Op_Type)
	// fmt.Printf("p.Payload:%s\n", receivedPacket.Payload)
// 	return

// }

func main() {
	// var sql string
	// go KeepListening("127.0.0.1",8006)  // test locally
	// for true {
		// fmt.Scanln(&sql)
		// print(sql)
		sql := "insert into student2 values(1080100001,'name1',99);"
		p := Type.Packet{Head: Type.PacketHead{P_Type: Type.SQLOperation , Op_Type: Type.Insert},
			Payload: []byte(sql)}

		ConnectToRegion("192.168.84.36", 3037, p) //RegionIP + RegionPort
	// }
}

//client 8005 send -> listen
//region 8006
