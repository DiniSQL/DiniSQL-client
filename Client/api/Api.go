package main

import (
	// "DiniSQL/MiniSQL/src/Interpreter/parser"
	"DiniSQL/MiniSQL/src/Interpreter/types"
	// "DiniSQL/MiniSQL/src/Interpreter/value"
	// "DiniSQL/MiniSQL/src/Utils"
	"DiniSQL/MiniSQL/src/Utils/Error"
	// "errors"
	"fmt"
	"DiniSQL/Client/type"
	"DiniSQL/Client/clientSocket"
	// "os"
	// "sync"
)
// import(
// 	"fmt"
// )


//HandleOneParse 用来处理parse处理完的DStatement类型  dataChannel是接收Statement的通道,整个mysql运行过程中不会关闭，但是quit后就会关闭
//stopChannel 用来发送同步信号，每次处理完一个后就发送一个信号用来同步两协程，主协程需要接收到stopChannel的发送后才能继续下一条指令，当dataChannel
//关闭后，stopChannel才会关闭
func HandleOneParse(dataChannel <-chan types.DStatements, stopChannel chan<- Error.Error,sqlChannel <-chan string) {
	var err Error.Error
	for statement := range dataChannel {
		var sql string
		sql=<-sqlChannel
		switch statement.GetOperationType() {
		case types.CreateDatabase:
			fmt.Println("CreateDatabase")
			var Head Type.PacketHead
			Head.P_Type=Type.Ask
			Head.Op_Type=Type.CreateDatabase
			var Packet Type.Packet
			Packet.Head=Head
			var sqlByte []byte = []byte(sql)
			// fmt.Println("sql:"+sql)
			Packet.Payload=sqlByte
			Socket.ConnectToRegion("10.192.182.120",8004,Packet);
			// err = CreateDatabaseAPI(statement.(types.CreateDatabaseStatement))
			// if err.Status != true {
			// 	fmt.Println(err.ErrorHint)
			// } else {
			// 	fmt.Println("create datbase success.")
			// }

		case types.UseDatabase:
			fmt.Println("UseDatabase")
			var Head Type.PacketHead
			Head.P_Type=Type.Ask
			Head.Op_Type=Type.UseDatabase
			var Packet Type.Packet
			Packet.Head=Head
			var sqlByte []byte = []byte(sql)
			// fmt.Println("sql:"+sql)
			Packet.Payload=sqlByte
			// err = UseDatabaseAPI(statement.(types.UseDatabaseStatement))
			// if err.Status != true {
			// 	fmt.Println(err.ErrorHint)
			// } else {
			// 	fmt.Printf("now you are using database.\n")
			// }
		case types.CreateTable:  //M
			fmt.Println("CreateTable")
			var Head Type.PacketHead
			Head.P_Type=Type.Ask
			Head.Op_Type=Type.CreateTable
			var Packet Type.Packet
			Packet.Head=Head
			var sqlByte []byte = []byte(sql)
			Packet.Payload=sqlByte
			// err = CreateTableAPI(statement.(types.CreateTableStatement))
			// if err.Status != true {
			// 	fmt.Println(err.ErrorHint)
			// } else {
			// 	fmt.Printf("create table succes.\n")
			// }

		case types.CreateIndex:  //M
			fmt.Println("CreateIndex")
			var Head Type.PacketHead
			Head.P_Type=Type.Ask
			Head.Op_Type=Type.CreateIndex
			var Packet Type.Packet
			Packet.Head=Head
			var sqlByte []byte = []byte(sql)
			// fmt.Println("sql:"+sql)
			Packet.Payload=sqlByte
			// err = CreateIndexAPI(statement.(types.CreateIndexStatement))
			// if err.Status != true {
			// 	fmt.Println(err.ErrorHint)
			// } else {
			// 	fmt.Printf("create index succes.\n")
			// }
		case types.DropTable:    //M
			fmt.Println("DropTable")
			var Head Type.PacketHead
			Head.P_Type=Type.Ask
			Head.Op_Type=Type.DropTable
			var Packet Type.Packet
			Packet.Head=Head
			var sqlByte []byte = []byte(sql)
			// fmt.Println("sql:"+sql)
			Packet.Payload=sqlByte
			// err = DropTableAPI(statement.(types.DropTableStatement))
			// if err.Status != true {
			// 	fmt.Println(err.ErrorHint)
			// } else {
			// 	fmt.Printf("drop table succes.\n")
			// }

		case types.DropIndex:    //M
			fmt.Println("DropIndex")
			var Head Type.PacketHead
			Head.P_Type=Type.Ask
			Head.Op_Type=Type.DropIndex
			var Packet Type.Packet
			Packet.Head=Head
			var sqlByte []byte = []byte(sql)
			// fmt.Println("sql:"+sql)
			Packet.Payload=sqlByte
			// err = DropIndexAPI(statement.(types.DropIndexStatement))
			// if err.Status != true {
			// 	fmt.Println(err.ErrorHint)
			// } else {
			// 	fmt.Printf("drop index succes.\n")
			// }

		case types.Insert:       //M
			fmt.Println("Insert")
			var Head Type.PacketHead
			Head.P_Type=Type.Ask
			Head.Op_Type=Type.Insert
			var Packet Type.Packet
			Packet.Head=Head
			var sqlByte []byte = []byte(sql)
			// fmt.Println("sql:"+sql)
			Packet.Payload=sqlByte
			// err = InsertAPI(statement.(types.InsertStament))
			// if err.Status != true {
			// 	fmt.Println(err.ErrorHint)
			// } else {
			// 	fmt.Printf("insert success, 1 row affected.\n")
			// }
		case types.Update:       //M
			fmt.Println("Update")
			var Head Type.PacketHead
			Head.P_Type=Type.Ask
			Head.Op_Type=Type.Update
			var Packet Type.Packet
			Packet.Head=Head
			var sqlByte []byte = []byte(sql)
			// fmt.Println("sql:"+sql)
			Packet.Payload=sqlByte
			// err = UpdateAPI(statement.(types.UpdateStament))
			// if err.Status != true {
			// 	fmt.Println(err.ErrorHint)
			// } else {
			// 	fmt.Printf("update success, %d rows are updated.\n", err.Rows)
			// }
		case types.Delete:       //M
			fmt.Println("Delete")
			var Head Type.PacketHead
			Head.P_Type=Type.Ask
			Head.Op_Type=Type.Delete
			var Packet Type.Packet
			Packet.Head=Head
			var sqlByte []byte = []byte(sql)
			// fmt.Println("sql:"+sql)
			Packet.Payload=sqlByte
			// err = DeleteAPI(statement.(types.DeleteStatement))
			// if err.Status != true {
			// 	fmt.Println(err.ErrorHint)
			// } else {
			// 	fmt.Printf("delete success, %d rows are deleted.\n", err.Rows)
			// }
		case types.Select:      //R或M
			fmt.Println("Select")
			// statement2:=statement.(types.SelectStatement)
			// tableNames:=statement2.TableNames
			var Head Type.PacketHead
			Head.P_Type=Type.Ask
			Head.Op_Type=Type.Select
			var Packet Type.Packet
			Packet.Head=Head
			var sqlByte []byte = []byte(sql)
			// fmt.Println("sql:"+sql)
			Packet.Payload=sqlByte
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

// //CreateDatabaseAPI 只调用CM，和IM、RM无关
// func CreateDatabaseAPI(statement types.CreateDatabaseStatement) Error.Error {

// 	err := CatalogManager.CreateDatabase(statement.DatabaseId)
// 	if err != nil {
// 		return Error.CreateFailError(err)

// 	}
// 	return Error.CreateSuccessError()
// }

// //UseDatabaseAPI 只调用CM，和IM、RM无关
// func UseDatabaseAPI(statement types.UseDatabaseStatement) Error.Error {
// 	err := CatalogManager.UseDatabase(statement.DatabaseId)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	return Error.CreateSuccessError()
// }

// //DropDatabaseAPI  先CM的check，和IM、RM无关 ，再调用RM的drop ， 再在CM中删除并flush
// func DropDatabaseAPI(statement types.DropDatabaseStatement) Error.Error {
// 	err := CatalogManager.DropDatabaseCheck(statement.DatabaseId)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	err = RecordManager.DropDatabase(statement.DatabaseId)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	err = CatalogManager.DropDatabase(statement.DatabaseId)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	return Error.CreateSuccessError()
// }

// //CreateTableAPI CM进行检查，index检查 语法检查  之后调用RM中的CreateTable创建表， 之后使用RM中的CreateIndex建索引
// func CreateTableAPI(statement types.CreateTableStatement) Error.Error {
// 	err, indexs := CatalogManager.CreateTableCheck(statement)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	err = RecordManager.CreateTable(statement.TableName)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	for _, item := range indexs {
// 		err = RecordManager.CreateIndex(CatalogManager.GetTableCatalogUnsafe(statement.TableName), item)
// 		if err != nil {
// 			return Error.CreateFailError(err)
// 		}
// 	}
// 	err = CatalogManager.FlushDatabaseMeta(CatalogManager.UsingDatabase.DatabaseId)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	return Error.CreateSuccessError()
// }

// //CreateIndexAPI CM进行检查，index语法检查 之后使用RM中的CreateIndex建索引
// func CreateIndexAPI(statement types.CreateIndexStatement) Error.Error {
// 	err, indexCatalog := CatalogManager.CreateIndexCheck(statement)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	err = RecordManager.CreateIndex(CatalogManager.GetTableCatalogUnsafe(statement.TableName), *indexCatalog)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	err = CatalogManager.FlushDatabaseMeta(CatalogManager.UsingDatabase.DatabaseId)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	return Error.CreateSuccessError()
// }

// //DropTableAPI CM进行检查，注意这个时候并不真正删除CM中的记录， 之后RM的DropTable删除table文件以及index文件， 之后让CM删除map中的记录同时flush
// func DropTableAPI(statement types.DropTableStatement) Error.Error {
// 	err := CatalogManager.DropTableCheck(statement)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	err = RecordManager.DropTable(statement.TableName)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	err = CatalogManager.DropTable(statement)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	err = CatalogManager.FlushDatabaseMeta(CatalogManager.UsingDatabase.DatabaseId)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	return Error.CreateSuccessError()
// }

// //DropIndexAPI CM进行检查， RM中删除index 之后CM中再删除并flush
// func DropIndexAPI(statement types.DropIndexStatement) Error.Error {
// 	err := CatalogManager.DropIndexCheck(statement)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	err = RecordManager.DropIndex(CatalogManager.GetTableCatalogUnsafe(statement.TableName), statement.IndexName)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	err = CatalogManager.DropIndex(statement)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	err = CatalogManager.FlushDatabaseMeta(CatalogManager.UsingDatabase.DatabaseId)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	return Error.CreateSuccessError()
// }

// //InsertAPI nothing to explain
// func InsertAPI(statement types.InsertStament) Error.Error {
// 	err, colPos, startBytePos, uniquescolumns := CatalogManager.InsertCheck(statement)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	err = RecordManager.InsertRecord(CatalogManager.GetTableCatalogUnsafe(statement.TableName), colPos, startBytePos, statement.Values, uniquescolumns)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	return Error.CreateRowsError(1)
// }

// //UpdateAPI nothing to explain
// func UpdateAPI(statement types.UpdateStament) Error.Error {
// 	err, setColumns, values, exprLSRV := CatalogManager.UpdateCheck(statement)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	var rowNum int
// 	if exprLSRV == nil {
// 		err, rowNum = RecordManager.UpdateRecord(CatalogManager.GetTableCatalogUnsafe(statement.TableName), setColumns, values, statement.Where)
// 	} else {
// 		err, rowNum = RecordManager.UpdateRecordWithIndex(CatalogManager.GetTableCatalogUnsafe(statement.TableName), setColumns, values, statement.Where, *exprLSRV)
// 	}
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	return Error.CreateRowsError(rowNum)
// }

// //DeleteAPI nothing to explain
// func DeleteAPI(statement types.DeleteStatement) Error.Error {
// 	err, exprLSRV := CatalogManager.DeleteCheck(statement)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	var rowNum int
// 	if exprLSRV == nil {
// 		err, rowNum = RecordManager.DeleteRecord(CatalogManager.GetTableCatalogUnsafe(statement.TableName), statement.Where)
// 	} else {
// 		err, rowNum = RecordManager.DeleteRecordWithIndex(CatalogManager.GetTableCatalogUnsafe(statement.TableName), statement.Where, *exprLSRV)
// 	}
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}

// 	return Error.CreateRowsError(rowNum)
// }

// //SelectAPI nothing to explain
// func SelectAPI(statement types.SelectStatement) Error.Error {
// 	err, exprLSRV := CatalogManager.SelectCheck(statement)
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	var rows []value.Row
// 	if exprLSRV == nil {
// 		err, rows = RecordManager.SelectRecord(CatalogManager.GetTableCatalogUnsafe(statement.TableNames[0]), statement.Fields.ColumnNames, statement.Where)
// 	} else {
// 		err, rows = RecordManager.SelectRecordWithIndex(CatalogManager.GetTableCatalogUnsafe(statement.TableNames[0]), statement.Fields.ColumnNames, statement.Where, *exprLSRV)
// 	}
// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	//非常dirty  data里面第一行是列明 使用value.Bytes存
// 	var colNames []string
// 	if statement.Fields.SelectAll {
// 		colNames = CatalogManager.GetTableColumnsInOrder(statement.TableNames[0])
// 	} else {
// 		colNames = statement.Fields.ColumnNames
// 	}
// 	var ColumnNameRow value.Row
// 	ColumnNameRow.Values = make([]value.Value, 0, len(colNames))
// 	for _, item := range colNames {
// 		ColumnNameRow.Values = append(ColumnNameRow.Values, value.Bytes{Val: []byte(item)})
// 	}
// 	rows = append(rows, ColumnNameRow)

// 	return Error.CreateDataError(len(rows)-1, rows)

// }

// // ExecFileAPI 执行某文件  开辟两个新协程
// func ExecFileAPI(statement types.ExecFileStatement) Error.Error {
// 	//parse协程 有缓冲信道
// 	StatementChannel := make(chan types.DStatements, 500)
// 	FinishChannel := make(chan Error.Error, 500)
// 	if !Utils.Exists(statement.FileName) {
// 		return Error.CreateFailError(errors.New("file " + statement.FileName + " don't exist"))
// 	}
// 	reader, err := os.Open(statement.FileName)
// 	defer reader.Close()
// 	if err != nil {
// 		return Error.CreateFailError(errors.New("open file " + statement.FileName + " fail"))
// 	}
// 	var wg sync.WaitGroup
// 	wg.Add(1) //等待FinishChannel关闭

// 	go HandleOneParse(StatementChannel, FinishChannel) //begin the runtime for exec
// 	go func() {
// 		defer wg.Done()
// 		for _ = range FinishChannel {
// 			//TODO 更加优雅的处理方式
// 		}
// 	}()
// 	err = parser.Parse(reader, StatementChannel) //开始解析

// 	close(StatementChannel) //关闭StatementChannel，进而关闭FinishChannel

// 	wg.Wait()

// 	if err != nil {
// 		return Error.CreateFailError(err)
// 	}
// 	return Error.CreateSuccessError()
// }
