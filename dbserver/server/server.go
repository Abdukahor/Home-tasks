package main

import (
	"bufio"
	"fmt"
	"net"
)
const (
	Select = "select"
	Create = "create"
	Insert = "insert"
	Update = "update"
	Delete = "delete"
	Exit   = "exit"
)

var (
	connections []net.Conn
	db = map[string]*[]Student{}
	arrayOfStudents []Student
	//tempArr []Student
	y int
	outputArr =[]string{"|  ID|               Fname|Age|  IsStudent|  IsWorker|  IsTeacher|Average| Experience|",
		"|  ID|               Fname|Age|"}
)

func main() {
	ln, err := net.Listen("tcp", ":8111")

	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println(err)
		}
		user := logIn(conn)

		go handleConn(conn, user)
	}
}

func handleConn(conn net.Conn, userName string) {

	connections = append(connections, conn)

	_, err := conn.Write([]byte("Welcome to SQL Mr(s) " + userName + ";"))

	if err != nil {
		fmt.Println(err)
	}

	for {
		text, err := bufio.NewReader(conn).ReadString(';')

		if err != nil {
			conn.Close()
			removeConn(conn)
			//broadCastMsg(userName+" is offline\n", conn)
			break
		}

		text = text[:len(text) - 1]



		res := handleCmd(text)
		fmt.Println(res)
		_, err = conn.Write([]byte(res))

		if err != nil {
			fmt.Println(err)
			return
		}
		// broadCastMsg(userName+":"+text, conn)
	}

}

func removeConn(conn net.Conn) {
	var i int

	for i = range connections {
		if connections[i] == conn {
			break
		}
	}

	fmt.Println(i)

	if len(connections) > 1 {
		connections = append(connections[:i], connections[i+1:]...)
	} else {
		connections = nil
	}
}

func broadCastMsg(msg string, sourceConn net.Conn) {

	for _, conn := range connections {
		if sourceConn != conn {
			_, err := conn.Write([]byte(msg))

			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}

	msg = msg[:len(msg)-1]
	fmt.Println(msg)
}
