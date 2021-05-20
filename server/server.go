package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
	"unicode/utf8"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, utf8.RuneCountInString(s)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func updateDeadline(conn net.Conn) {
	if err := conn.SetDeadline(time.Now().Add(20 * time.Second)); err != nil {
		fmt.Println(err)
		return
	}
}

func handle(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(conn)

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("Message received:", message)
		splitedMessage := strings.Fields(message)
		if len(splitedMessage) == 0 {
			_, err := conn.Write([]byte("Please enter some words\n"))
			if err != nil {
				fmt.Println(err)
				return
			}
			continue
		}

		command := splitedMessage[0]
		var text string

		if len(splitedMessage) > 1 {
			text = strings.Join(splitedMessage[1:], " ")
		} else {
			text = ""
		}
		switch command {
		case "TIME":
			_, err := conn.Write([]byte(time.Now().String() + "\n"))
			updateDeadline(conn)
			if err != nil {
				fmt.Println(err)
				return
			}
		case "REVERSE":
			_, err := conn.Write([]byte(reverseString(text) + "\n"))
			if err != nil {
				fmt.Println(err)
				return
			}
		default:
			_, err := conn.Write([]byte(strings.ToUpper(message)))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		updateDeadline(conn)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(listener)

	for {
		conn, _ := listener.Accept()
		updateDeadline(conn)
		go handle(conn)
	}
}
