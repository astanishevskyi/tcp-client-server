package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func handle(conn net.Conn) {
	defer conn.Close()
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("Message received:", message)
		splitedMessage := strings.Fields(message)
		command := splitedMessage[0]
		var text string

		if len(splitedMessage) > 1 {
			text = strings.Join(splitedMessage[1:], " ")
		} else {
			text = ""
		}
		switch command {
		case "TIME":
			conn.Write([]byte(time.Now().String() + "\n"))
		case "REVERSE":
			conn.Write([]byte(reverseString(text) + "\n"))
		default:
			conn.Write([]byte(strings.ToUpper(message) + "\n"))
		}

		conn.SetDeadline(time.Now().Add(time.Minute)) // update timeout
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		conn.SetDeadline(time.Now().Add(time.Minute)) // set timeout
		go handle(conn)
	}
}
