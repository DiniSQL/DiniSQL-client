package main

import (
    "fmt"
    "log"
    "net"
	"strconv"
	"DiniSQL/Client/type"
	"io/ioutil"
)

var ch = make(chan []byte, 5)
var ClientIP = "127.0.0.1"
var ClientPort = 8005
// initial connection to region server and send message
func ConnectToRegion(regionIP string,regionPort int, packet Type.Packet)(recPacket Type.Packet){

    address := net.TCPAddr{
        IP:   net.ParseIP(regionIP), 
        Port: regionPort,
    }
	conn, err := net.DialTCP("tcp4", nil, &address)
	// defer conn.Close()
    if err != nil {
        log.Fatal(err) // Println + os.Exit(1)
		return 
    }
	var packetBuf = make([]byte, 0)
	fmt.Printf("packet.Head.P_Type:%d\n",packet.Head.P_Type)
	fmt.Printf("packet.Head.Op_Type:%d\n",packet.Head.Op_Type)
	fmt.Printf("packet.Payload:%s\n",packet.Payload)
	packetBuf,err = packet.MarshalMsg(packetBuf)
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
	conn.Close()
	fmt.Printf("send %d to %s\n", packet.Head.P_Type, conn.RemoteAddr())
	// listen after send
	recPacket=KeepListening(ClientIP, ClientPort)
	return 

}

// listen 
// input : IP and Port of client
func KeepListening(ClientIP string, ClientPort int)(receivedPacket Type.Packet){
	fmt.Printf("listening in %s...\n",ClientIP+":"+strconv.Itoa(ClientPort))
    listener, err := net.Listen("tcp", ":"+strconv.Itoa(ClientPort)) 
	defer listener.Close()
    if err != nil {
        log.Fatal(err) // Println + os.Exit(1)
    }

        conn, err := listener.Accept()
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("remote address:", conn.RemoteAddr())
		res,err1 := ioutil.ReadAll(conn)
		if err1 != nil {
            log.Println(err)
            conn.Close()
        }
		res,err1 = receivedPacket.UnmarshalMsg(res)
		// ch <- p
		if err1 != nil {
            log.Println(err)
            conn.Close()
        }

		fmt.Printf("p.Head.P_Type:%d\n",receivedPacket.Head.P_Type)
		fmt.Printf("p.Head.Op_Type:%d\n",receivedPacket.Head.Op_Type)
		fmt.Printf("p.Payload:%s\n",receivedPacket.Payload)
	return
    
}

func main(){
	var sql string
	// go KeepListening("127.0.0.1",8006)  // test locally
	for true  {
        fmt.Scanln(&sql)
		p := Type.Packet{Head:Type.PacketHead{P_Type: Type.KeepAlive, Op_Type: Type.CreateIndex},
			Payload: []byte(sql)}

		ConnectToRegion("172.20.10.10",9000,p)  //RegionIP + RegionPort
    }
}
//client 8005 send -> listen
//region 8006 