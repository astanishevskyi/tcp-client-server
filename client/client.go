package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Readers and writers are created here only once.
	inputReader := bufio.NewReader(os.Stdin)
	connReader := bufio.NewReader(conn)
	connWriter := bufio.NewWriter(conn)

	for {
		input, err := inputReader.ReadBytes('\n') // ReadBytes eliminates the need of conversion to bytes.
		if err != nil {
			fmt.Println(err)
			return
		}

		// Using a buffered writer here is optional. Just for example.
		_, err = connWriter.Write(input)
		if err != nil {
			fmt.Println(err)
			return
		}

		if err := connWriter.WriteByte('\n'); err != nil {
			fmt.Println(err)
			return
		}
		if err := connWriter.Flush(); err != nil {
			fmt.Println(err)
			return
		}

		message, err := connReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("->", message)
	}
}
