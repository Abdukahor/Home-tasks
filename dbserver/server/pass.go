package main

import (
	"crypto/sha512"
	eb64 "encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	_"time"
	"net"
	"bufio"
)

func checkHash(password string, hashStr string) bool {
	//now := time.Now()
	//timeFormat := now.Format("02.2006")
	hash := sha512.New()
	io.WriteString(hash, "root:"+password+":23.2019")
	return hashStr == eb64.RawStdEncoding.EncodeToString(hash.Sum(nil))
}

func readPasswordHash() string {
	file, err := ioutil.ReadFile("./pass.txt")

	if err != nil {
		fmt.Println("couldn't read file: " + err.Error())
		return ""
	}

	return string(file)

}



func logIn(conn net.Conn)  string{
	var logined bool

	readedPassHash := readPasswordHash()

	for !logined {

		_, err := conn.Write([]byte("Enter username: ;"))
		login, err := bufio.NewReader(conn).ReadString(';')

		if err != nil {
			fmt.Println("Some error occur: ", err.Error())
			conn.Close()
		}

		login = login[:len(login) - 1]

		_, _err := conn.Write([]byte("Enter pass: ;"))
		password, err := bufio.NewReader(conn).ReadString(';')

		if _err != nil {
			fmt.Println("Some error occur: ", err.Error())
			conn.Close()
		}

		password = password[:len(password) - 1]


		if login == "root" && checkHash(password, readedPassHash) {
			logined = true
			return login
		} else {
			s := "\nInvalid credentials\n"
			s += "-------------------\n;"
			_, err := conn.Write([]byte(s))

			if err != nil {
				fmt.Println("Some error occur: ", err.Error())
				conn.Close()
			}

			continue
		}
	}
	return ""
}


