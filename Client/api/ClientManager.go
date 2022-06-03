package main

import (
	// "DiniSQL/MiniSQL/src/API"
	"DiniSQL-client/Client/Interpreter/parser"
	"DiniSQL-client/Client/Interpreter/types"
	"DiniSQL-client/Client/Utils/Error"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/peterh/liner"
)

const historyCommmandFile = "~/.minisql_history"

func expandPath(path string) (string, error) {
	if strings.HasPrefix(path, "~/") {
		parts := strings.SplitN(path, "/", 2)
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, parts[1]), nil
	}
	return path, nil
}
func loadHistoryCommand() (*os.File, error) {
	var file *os.File
	path, err := expandPath(historyCommmandFile)
	if err != nil {
		return nil, err
	}
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		file, err = os.Create(path)
		if err != nil {
			return nil, err
		}
	} else {
		file, err = os.OpenFile(path, os.O_RDWR, 0666)
		if err != nil {
			return nil, err
		}
	}
	return file, err

}
func runShell(r chan<- error) {
	ll := liner.NewLiner()
	defer ll.Close()
	ll.SetCtrlCAborts(true)
	file, err := loadHistoryCommand()
	if err != nil {
		panic(err)
	}
	defer func() {
		_, err := ll.WriteHistory(file)
		if err != nil {
			panic(err)
		}
		_ = file.Close()
	}()
	s := bufio.NewScanner(file)
	for s.Scan() {
		//fmt.Println(s.Text())
		ll.AppendHistory(s.Text())
	}

	StatementChannel := make(chan types.DStatements, 500) //用于传输操作指令通道
	FinishChannel := make(chan Error.Error, 500)          //用于api执行完成反馈通道
	SQLChannel := make(chan string, 500)
	// FlushChannel := make(chan struct{})                    //用于每条指令结束后协程flush
	go HandleOneParse(StatementChannel, FinishChannel, SQLChannel) //begin the runtime for exec
	var beginSQLParse = false
	var sqlText = make([]byte, 0, 100)
	for { //each sql
	LOOP:
		beginSQLParse = false
		sqlText = sqlText[:0]
		var input string
		var err error
		for { //each line
			if beginSQLParse {
				input, err = ll.Prompt("      -> ")
			} else {
				input, err = ll.Prompt("DiniSQL> ")
			}
			if err != nil {
				if err == liner.ErrPromptAborted {
					goto LOOP
				}
			}
			trimInput := strings.TrimSpace(input) //get the input without front and backend space
			if len(trimInput) != 0 {
				ll.AppendHistory(input)
				if !beginSQLParse && (trimInput == "quit" || strings.HasPrefix(trimInput, "quit;")) {
					close(StatementChannel)
					close(SQLChannel)
					for _ = range FinishChannel {

					}
					r <- err
					return
				}
				sqlText = append(sqlText, append([]byte{' '}, []byte(trimInput)[0:]...)...)
				if !beginSQLParse {
					beginSQLParse = true
				}
				if strings.Contains(trimInput, ";") {
					break
				}
			}
		}
		beginTime := time.Now()
		// fmt.Println(string(sqlText))
		SQLChannel <- string(sqlText)
		err = parser.Parse(strings.NewReader(string(sqlText)), StatementChannel)
		//fmt.Println(string(sqlText))
		if err != nil {
			fmt.Println(err)
			sql := <-SQLChannel
			sql = sql + ""
			continue
		}
		<-FinishChannel //等待指令执行完成
		durationTime := time.Since(beginTime)
		fmt.Println("Finish operation at: ", durationTime)
		// FlushChannel <- struct{}{} //开始刷新cache
	}
}
func main() {
	fmt.Println()
	fmt.Println("_|_|_|    _|            _|    _|_|_|    _|_|      _|")
	fmt.Println("_|    _|      _|_|_|        _|        _|    _|    _|")
	fmt.Println("_|    _|      _|_|_|        _|        _|    _|    _|")
	fmt.Println("_|    _|  _|  _|    _|  _|    _|_|    _|  _|_|    _|")
	fmt.Println("_|    _|  _|  _|    _|  _|        _|  _|    _|    _|")
	fmt.Println("_|_|_|    _|  _|    _|  _|  _|_|_|      _|_|  _|  _|_|_|_|")
	fmt.Println()
	fmt.Println("Welcome to the DiniSQL monitor, a compact distributed SQL engine. Commands end with ;.")
	fmt.Println("Copyright (c) 2022, 2022, Ou Yixin, Fan Zhaoyu, Zheng Bowen, Liu Shuhan and Xiang Kerong.")
	fmt.Println()

	newCache()
	errChan := make(chan error)
	go runShell(errChan) //开启shell协程
	err := <-errChan
	fmt.Println("bye")
	if err != nil {
		fmt.Println(err)
	}
}
