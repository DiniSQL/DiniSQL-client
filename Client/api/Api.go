package main

import (
	// "DiniSQL/MiniSQL/src/Interpreter/parser"
	"DiniSQL-client/Client/Interpreter/types"
	"DiniSQL-client/Client/Type"
	"DiniSQL-client/Client/clientSocket"
	"log"
	"strconv"
	"strings"

	// "DiniSQL/MiniSQL/src/Interpreter/value"
	// "DiniSQL/MiniSQL/src/Utils"
	"DiniSQL-client/Client/Utils/Error"
	// "errors"
	"fmt"
	// "os"
	// "sync"
)

// import(
// 	"fmt"
// )
var MasterIP string = "172.20.10.2"
var MasterPort int = 9000

//HandleOneParse 用来处理parse处理完的DStatement类型  dataChannel是接收Statement的通道,整个mysql运行过程中不会关闭，但是quit后就会关闭
//stopChannel 用来发送同步信号，每次处理完一个后就发送一个信号用来同步两协程，主协程需要接收到stopChannel的发送后才能继续下一条指令，当dataChannel
//关闭后，stopChannel才会关闭
func HandleOneParse(dataChannel <-chan types.DStatements, stopChannel chan<- Error.Error, sqlChannel <-chan string) {
	var err Error.Error
	for statement := range dataChannel {
		var sql string
		sql = <-sqlChannel
		switch statement.GetOperationType() {
		case types.CreateDatabase:
			// fmt.Println("CreateDatabase")
			p := Type.Packet{Head: Type.PacketHead{P_Type: Type.Ask, Op_Type: Type.CreateDatabase},
				Payload: []byte(sql)}
			result := clientSocket.ConnectToRegion(MasterIP, MasterPort, p)
			// result := Type.Packet{Head: Type.PacketHead{P_Type: Type.Answer, Op_Type: Type.CreateDatabase}, Signal: true,
			// 	Payload: []byte("Successhhh"), IPResult: []byte("10.1.1.2:2020;10.2.2.1:1000")}
			if result.Signal == true {
				if len(result.Payload) > 0 {
					fmt.Println(string(result.Payload))
				}
			} else {
				fmt.Println(string(result.Payload))
			}

		case types.UseDatabase:
			// fmt.Println("UseDatabase")
			p := Type.Packet{Head: Type.PacketHead{P_Type: Type.Ask, Op_Type: Type.UseDatabase},
				Payload: []byte(sql)}
			result := clientSocket.ConnectToRegion(MasterIP, MasterPort, p)
			if result.Signal == true {
				if len(result.Payload) > 0 {
					fmt.Println(string(result.Payload))
				}
			} else {
				fmt.Println(string(result.Payload))
			}

		case types.CreateTable:
			// fmt.Println("CreateTable")
			p := Type.Packet{Head: Type.PacketHead{P_Type: Type.Ask, Op_Type: Type.CreateTable},
				Payload: []byte(sql)}
			// fmt.Println(sql)
			result := clientSocket.ConnectToRegion(MasterIP, MasterPort, p)
			// result := Type.Packet{Head: Type.PacketHead{P_Type: Type.Answer, Op_Type: Type.CreateDatabase}, Signal: true,
			// 	Payload: []byte("Successhhh"), IPResult: []byte("10.1.1.2:2020;10.2.2.1:1000")}
			fmt.Print("Master发来的SQL语句结果：")
			if result.Signal == true {
				if len(result.Payload) > 0 {
					fmt.Println(string(result.Payload))
				}
				IPs := strings.Split(string(result.IPResult), ";")
				// fmt.Println(IPs)
				setCache(statement.(types.CreateTableStatement).TableName, IPs)
				fmt.Print("Master发来的table对应的IP信息：")
				fmt.Println(IPs)
			} else {
				fmt.Println(string(result.Payload))
				// fmt.Println(string(result.IPResult))
			}

		case types.CreateIndex: //M
			// fmt.Println("CreateIndex")
			p := Type.Packet{Head: Type.PacketHead{P_Type: Type.Ask, Op_Type: Type.CreateIndex},
				Payload: []byte(sql)}
			result := clientSocket.ConnectToRegion(MasterIP, MasterPort, p)
			fmt.Print("Master发来的SQL语句结果：")
			if result.Signal == true {
				if len(result.Payload) > 0 {
					fmt.Println(string(result.Payload))
				}
				IPs := strings.Split(string(result.IPResult), ";")
				setCache(statement.(types.CreateIndexStatement).TableName, IPs)
				fmt.Print("Master发来的table对应的IP信息：")
				fmt.Println(IPs)
			} else {
				fmt.Println(string(result.Payload))
			}

		case types.DropTable: //M
			// fmt.Println("DropTable")
			p := Type.Packet{Head: Type.PacketHead{P_Type: Type.Ask, Op_Type: Type.DropTable},
				Payload: []byte(sql)}
			result := clientSocket.ConnectToRegion(MasterIP, MasterPort, p)
			fmt.Print("Master发来的SQL语句结果：")
			if result.Signal == true {
				if len(result.Payload) > 0 {
					fmt.Println(string(result.Payload))
					deleteKey(statement.(types.DropTableStatement).TableName)
				}
			} else {
				fmt.Println(string(result.Payload))
			}

		case types.DropIndex: //M
			// fmt.Println("DropIndex")
			p := Type.Packet{Head: Type.PacketHead{P_Type: Type.Ask, Op_Type: Type.DropIndex},
				Payload: []byte(sql)}
			result := clientSocket.ConnectToRegion(MasterIP, MasterPort, p)
			fmt.Print("Master发来的SQL语句结果：")
			if result.Signal == true {
				if len(result.Payload) > 0 {
					fmt.Println(string(result.Payload))
				}
				IPs := strings.Split(string(result.IPResult), ";")
				setCache(statement.(types.DropIndexStatement).TableName, IPs)
				fmt.Print("Master发来的table对应的IP信息：")
				fmt.Println(IPs)
			} else {
				fmt.Println(string(result.Payload))
			}

		case types.Insert: //M
			// fmt.Println("Insert")
			p := Type.Packet{Head: Type.PacketHead{P_Type: Type.Ask, Op_Type: Type.Insert},
				Payload: []byte(sql)}
			result := clientSocket.ConnectToRegion(MasterIP, MasterPort, p)
			fmt.Print("Master发来的SQL语句结果：")
			if result.Signal == true {
				if len(result.Payload) > 0 {
					fmt.Println(string(result.Payload))
				}
				IPs := strings.Split(string(result.IPResult), ";")
				setCache(statement.(types.InsertStament).TableName, IPs)
				fmt.Print("Master发来的table对应的IP信息：")
				fmt.Println(IPs)
			} else {
				fmt.Println(string(result.Payload))
			}

		case types.Update: //M
			// fmt.Println("Update")
			p := Type.Packet{Head: Type.PacketHead{P_Type: Type.Ask, Op_Type: Type.Update},
				Payload: []byte(sql)}
			result := clientSocket.ConnectToRegion(MasterIP, MasterPort, p)
			fmt.Print("Master发来的SQL语句结果：")
			if result.Signal == true {
				if len(result.Payload) > 0 {
					fmt.Println(string(result.Payload))
				}
				IPs := strings.Split(string(result.IPResult), ";")
				setCache(statement.(types.UpdateStament).TableName, IPs)
				fmt.Print("Master发来的table对应的IP信息：")
				fmt.Println(IPs)
			} else {
				fmt.Println(string(result.Payload))
			}

		case types.Delete: //M
			// fmt.Println("Delete")
			p := Type.Packet{Head: Type.PacketHead{P_Type: Type.Ask, Op_Type: Type.Delete},
				Payload: []byte(sql)}
			result := clientSocket.ConnectToRegion(MasterIP, MasterPort, p)
			fmt.Print("Master发来的SQL语句结果：")
			if result.Signal == true {
				if len(result.Payload) > 0 {
					fmt.Println(string(result.Payload))
				}
				IPs := strings.Split(string(result.IPResult), ";")
				setCache(statement.(types.DeleteStatement).TableName, IPs)
				fmt.Print("Master发来的table对应的IP信息：")
				fmt.Println(IPs)
			} else {
				fmt.Println(string(result.Payload))
			}

		case types.Select: //R或M
			// fmt.Println("Select")
			tableName := statement.(types.SelectStatement).TableNames[0]
			fmt.Print("当前缓存中的数据：")
			fmt.Println(cache)
			if getCache(tableName) == nil {
				fmt.Println("当前缓存中不存在该表对应IP")
				p := Type.Packet{Head: Type.PacketHead{P_Type: Type.Ask, Op_Type: Type.Select},
					Payload: []byte(sql)}
				getRegionsIP(p, tableName)
			}
			// fmt.Println(getCache(tableName))
			if getCache(tableName) != nil {
				// fmt.Println("I'm here!!!")
				p := Type.Packet{Head: Type.PacketHead{P_Type: Type.SQLOperation, Op_Type: Type.Select},
					Payload: []byte(sql)}
				IPAndPort := strings.Split(getCache(tableName)[0], ":")
				IP := IPAndPort[0]
				Port, err := strconv.Atoi(IPAndPort[1])
				if err != nil {
					log.Fatal(err)
				}
				// fmt.Println("here1")
				// fmt.Println(cache)
				result := clientSocket.ConnectToRegion(IP, Port, p)
				// fmt.Println("here2")
				if result.Head.Op_Type == -1 {
					// fmt.Println("here3")
					fmt.Println("当前缓存中该表对应IP已经过期")
					p := Type.Packet{Head: Type.PacketHead{P_Type: Type.Ask, Op_Type: Type.Select},
						Payload: []byte(sql)}
					getRegionsIP(p, tableName)
					p = Type.Packet{Head: Type.PacketHead{P_Type: Type.SQLOperation, Op_Type: Type.Select},
						Payload: []byte(sql)}
					IPAndPort := strings.Split(getCache(tableName)[0], ":")
					IP := IPAndPort[0]
					Port, _ := strconv.Atoi(IPAndPort[1])
					result = clientSocket.ConnectToRegion(IP, Port, p)
				}
				fmt.Println("Region发来的select结果")
				// fmt.Println("here4")
				if result.Signal == true {
					if len(result.Payload) > 0 {
						fmt.Println(string(result.Payload))
					}
				} else {
					fmt.Println(string(result.Payload))
				}
			}
		}
		//fmt.Println(err)
		stopChannel <- err
	}
	close(stopChannel)
}
func getRegionsIP(packet Type.Packet, tableName string) {
	// fmt.Println("getRegionIP")
	result := clientSocket.ConnectToRegion(MasterIP, MasterPort, packet)
	if result.Signal == true {
		if len(result.Payload) > 0 {
			fmt.Println(string(result.Payload))
		}
		// fmt.Println("IPs" + string(result.IPResult))
		IPs := strings.Split(string(result.IPResult), ";")
		fmt.Print("Master发来的table对应的IP信息：")
		fmt.Println(IPs)
		setCache(tableName, IPs)
	} else {
		fmt.Println(string(result.Payload))
	}
	// fmt.Println(cache)
}
