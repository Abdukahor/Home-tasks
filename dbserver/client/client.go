package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8111")

	if err != nil {
		panic(err)
	}
	
	
	go writeMsg(conn)
	readMsg(conn)
}

func writeMsg(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		_, err := conn.Write([]byte(text))

		if err != nil {
			panic(err)
		}
	}
}

func readMsg(conn net.Conn) {

	for {
		msg, err := bufio.NewReader(conn).ReadString(';')

		if err == io.EOF {
			fmt.Println("Connection close, Bye!")
			conn.Close()
			panic(err)
		} else if err != nil {
			fmt.Println(err.Error())
			conn.Close()
			panic(err)
		}

		msg = msg[:len(msg)-1]
		fmt.Println(string(msg))
	}
}
