package main

import (
	// "DiniSQL/MiniSQL/src/Interpreter/parser"
	"DiniSQL-client/Client/Interpreter/types"
	"DiniSQL-client/Client/Type"
	"DiniSQL-client/Client/clientSocket"

	// "DiniSQL/MiniSQL/src/Interpreter/value"
	// "DiniSQL/MiniSQL/src/Utils"
	"DiniSQL-client/Client/Utils/Error"
	// "errors"
	"fmt"
	"strings"
	// "os"
	// "sync"
)

// import(
// 	"fmt"
// )
var MasterIP string = "192.168.84.13"
var MasterPort int = 8006

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
			fmt.Println("CreateDatabase")
			// p := Type.Packet{Head: Type.PacketHead{P_Type: Type.Ask, Op_Type: Type.CreateDatabase},
			// 	Payload: []byte(sql)}
			// result := clientSocket.ConnectToRegion(MasterIP, MasterPort, p)
			result := Type.MasterClientPacket{Head: Type.PacketHead{P_Type: Type.Answer, Op_Type: Type.CreateDatabase}, Signal: true,
				SQLResult: []byte("Successhhh"), IPResult: []byte("10.1.1.2:2020;10.2.2.1:1000")}
			printMasterResult(result)
			// err = CreateDatabaseAPI(statement.(types.CreateDatabaseStatement))
			// if err.Status != true {
			// 	fmt.Println(err.ErrorHint)
			// } else {
			// 	fmt.Println("create datbase success.")
			// }

		case types.UseDatabase:
			fmt.Println("UseDatabase")
			p := Type.Packet{Head: Type.PacketHead{P_Type: Type.Ask, Op_Type: Type.UseDatabase},
				Payload: []byte(sql)}
			clientSocket.ConnectToRegion(MasterIP, MasterPort, p)
			// result := clientSocket.ConnectToRegion(MasterIP, MasterPort, p)

			// err = UseDatabaseAPI(statement.(types.UseDatabaseStatement))
			// if err.Status != true {
			// 	fmt.Println(err.ErrorHint)
			// } else {
			// 	fmt.Printf("now you are using database.\n")
			// }
		case types.CreateTable: //M
			fmt.Println("CreateTable")
			// var Head Type.PacketHead
			// Head.P_Type = Type.Ask
			// Head.Op_Type = Type.CreateTable
			// var Packet Type.Packet
			// Packet.Head = Head
			// var sqlByte []byte = []byte(sql)
			// Packet.Payload = sqlByte
			p := Type.Packet{Head: Type.PacketHead{P_Type: Type.Ask, Op_Type: Type.CreateTable},
				Payload: []byte(sql)}
			clientSocket.ConnectToRegion(MasterIP, MasterPort, p)
			// err = CreateTableAPI(statement.(types.CreateTableStatement))
			// if err.Status != true {
			// 	fmt.Println(err.ErrorHint)
			// } else {
			// 	fmt.Printf("create table succes.\n")
			// }

		case types.CreateIndex: //M
			fmt.Println("CreateIndex")
			// var Head Type.PacketHead
			// Head.P_Type = Type.Ask
			// Head.Op_Type = Type.CreateIndex
			// var Packet Type.Packet
			// Packet.Head = Head
			// var sqlByte []byte = []byte(sql)
			// // fmt.Println("sql:"+sql)
			// Packet.Payload = sqlByte
			p := Type.Packet{Head: Type.PacketHead{P_Type: Type.Ask, Op_Type: Type.CreateIndex},
				Payload: []byte(sql)}
			clientSocket.ConnectToRegion(MasterIP, MasterPort, p)
			// err = CreateIndexAPI(statement.(types.CreateIndexStatement))
			// if err.Status != true {
			// 	fmt.Println(err.ErrorHint)
			// } else {
			// 	fmt.Printf("create index succes.\n")
			// }
		case types.DropTable: //M
			fmt.Println("DropTable")
			// var Head Type.PacketHead
			// Head.P_Type = Type.Ask
			// Head.Op_Type = Type.DropTable
			// var Packet Type.Packet
			// Packet.Head = Head
			// var sqlByte []byte = []byte(sql)
			// // fmt.Println("sql:"+sql)
			// Packet.Payload = sqlByte
			p := Type.Packet{Head: Type.PacketHead{P_Type: Type.Ask, Op_Type: Type.DropTable},
				Payload: []byte(sql)}
			clientSocket.ConnectToRegion(MasterIP, MasterPort, p)
			// err = DropTableAPI(statement.(types.DropTableStatement))
			// if err.Status != true {
			// 	fmt.Println(err.ErrorHint)
			// } else {
			// 	fmt.Printf("drop table succes.\n")
			// }

		case types.DropIndex: //M
			fmt.Println("DropIndex")
			// var Head Type.PacketHead
			// Head.P_Type = Type.Ask
			// Head.Op_Type = Type.DropIndex
			// var Packet Type.Packet
			// Packet.Head = Head
			// var sqlByte []byte = []byte(sql)
			// // fmt.Println("sql:"+sql)
			// Packet.Payload = sqlByte
			p := Type.Packet{Head: Type.PacketHead{P_Type: Type.Ask, Op_Type: Type.DropIndex},
				Payload: []byte(sql)}
			clientSocket.ConnectToRegion(MasterIP, MasterPort, p)
			// err = DropIndexAPI(statement.(types.DropIndexStatement))
			// if err.Status != true {
			// 	fmt.Println(err.ErrorHint)
			// } else {
			// 	fmt.Printf("drop index succes.\n")
			// }

		case types.Insert: //M
			fmt.Println("Insert")
			// var Head Type.PacketHead
			// Head.P_Type = Type.Ask
			// Head.Op_Type = Type.Insert
			// var Packet Type.Packet
			// Packet.Head = Head
			// var sqlByte []byte = []byte(sql)
			// // fmt.Println("sql:"+sql)
			// Packet.Payload = sqlByte
			p := Type.Packet{Head: Type.PacketHead{P_Type: Type.Ask, Op_Type: Type.Insert},
				Payload: []byte(sql)}
			clientSocket.ConnectToRegion(MasterIP, MasterPort, p)
			// err = InsertAPI(statement.(types.InsertStament))
			// if err.Status != true {
			// 	fmt.Println(err.ErrorHint)
			// } else {
			// 	fmt.Printf("insert success, 1 row affected.\n")
			// }

		case types.Update: //M
			fmt.Println("Update")
			// var Head Type.PacketHead
			// Head.P_Type = Type.Ask
			// Head.Op_Type = Type.Update
			// var Packet Type.Packet
			// Packet.Head = Head
			// var sqlByte []byte = []byte(sql)
			// // fmt.Println("sql:"+sql)
			// Packet.Payload = sqlByte
			p := Type.Packet{Head: Type.PacketHead{P_Type: Type.Ask, Op_Type: Type.Update},
				Payload: []byte(sql)}
			clientSocket.ConnectToRegion(MasterIP, MasterPort, p)
			// err = UpdateAPI(statement.(types.UpdateStament))
			// if err.Status != true {
			// 	fmt.Println(err.ErrorHint)
			// } else {
			// 	fmt.Printf("update success, %d rows are updated.\n", err.Rows)
			// }
		case types.Delete: //M
			fmt.Println("Delete")
			// var Head Type.PacketHead
			// Head.P_Type = Type.Ask
			// Head.Op_Type = Type.Delete
			// var Packet Type.Packet
			// Packet.Head = Head
			// var sqlByte []byte = []byte(sql)
			// // fmt.Println("sql:"+sql)
			// Packet.Payload = sqlByte
			p := Type.Packet{Head: Type.PacketHead{P_Type: Type.Ask, Op_Type: Type.Delete},
				Payload: []byte(sql)}
			clientSocket.ConnectToRegion(MasterIP, MasterPort, p)
			// err = DeleteAPI(statement.(types.DeleteStatement))
			// if err.Status != true {
			// 	fmt.Println(err.ErrorHint)
			// } else {
			// 	fmt.Printf("delete success, %d rows are deleted.\n", err.Rows)
			// }

		case types.Select: //R或M
			fmt.Println("Select")
			// statement2:=statement.(types.SelectStatement)
			// tableNames:=statement2.TableNames
			// var Head Type.PacketHead
			// Head.P_Type = Type.Ask
			// Head.Op_Type = Type.Select
			// var Packet Type.Packet
			// Packet.Head = Head
			// var sqlByte []byte = []byte(sql)
			// // fmt.Println("sql:"+sql)
			// Packet.Payload = sqlByte
			p := Type.Packet{Head: Type.PacketHead{P_Type: Type.Ask, Op_Type: Type.Select},
				Payload: []byte(sql)}
			clientSocket.ConnectToRegion(MasterIP, MasterPort, p)
			// err = SelectAPI(statement.(types.SelectStatement))
			// if err.Status != true {
			// 	fmt.Println(err.ErrorHint)
			// } else {
			// 	PrintTable(statement.(types.SelectStatement).TableNames[0], err.Data[err.Rows], err.Data[0:err.Rows]) //very dirty  but I have no other choose
			// }
			// case types.ExecFile:
			// 	err = ExecFileAPI(statement.(types.ExecFileStatement))
		}
		//fmt.Println(err)
		stopChannel <- err
	}
	close(stopChannel)
}

func printMasterResult(result Type.MasterClientPacket) {
	if result.Signal == true {
		if len(result.SQLResult) > 0 {
			fmt.Println(string(result.SQLResult))
		}
		IPs := strings.Split(string(result.IPResult), ";")
		fmt.Println(IPs)
	} else {
		fmt.Println(string(result.SQLResult))
	}
}
