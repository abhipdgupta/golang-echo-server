package server

import (
	"echo-server/config"
	"fmt"
	"io"
	"net"
	"strings"
)

func readMessage(conn net.Conn) (string, error) {
	buff := make([]byte, 1024)
	i, err := conn.Read(buff)

	if err != nil {
		return "", err
	}

	return string(buff[:i]), nil

}

func writeMessage(conn net.Conn, message string) error {
	if _, err := conn.Write([]byte(fmt.Sprintln(message))); err != nil {
		return err
	}

	return nil
}

func RunServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		panic(err)
	}

	defer lis.Close()

	var numberOfConnections int = 0
	fmt.Println("Started synchronous echo server")

	for {
		// blocking call : waits for a connection
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println(err)
		}

		numberOfConnections += 1
		fmt.Printf("Accepted connection from %s with id %d\n", conn.RemoteAddr(), numberOfConnections)

		for {
			message, err := readMessage(conn)

			if err != nil {
				conn.Close()
				numberOfConnections -= 1
				fmt.Printf("Connection closed by %s with id %d\n", conn.RemoteAddr(), numberOfConnections)
				if err == io.EOF {
					fmt.Printf("Connection closed by %s with id %d\n", conn.RemoteAddr(), numberOfConnections)
					break
				}

				fmt.Println(err)

			}
			// trime message
			message = strings.TrimSpace(message)

			if message == "exit()" {
				fmt.Printf("Connection closed by %s with id %d\n", conn.RemoteAddr(), numberOfConnections)
				numberOfConnections -= 1
				conn.Close()
				break
			}

			fmt.Println(message)

			if err := writeMessage(conn, message); err != nil {
				fmt.Println(err)
			}

		}
	}

}
