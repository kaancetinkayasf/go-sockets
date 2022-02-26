package main

import (
	"fmt"
	"net"
)

var connectionList []net.Conn

func main() {

	listener, err := net.Listen("tcp4", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Server started")

	for {

		connection, err := listener.Accept()

		connectionList = append(connectionList, connection)
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	for {

		buffer := make([]byte, 8)
		_, err := connection.Read(buffer[:])
		if err != nil {
			fmt.Println(err)
			connection.Close()
			break
		}
		fmt.Printf("%s\n", buffer)

		for _, client := range connectionList {
			_, err := client.Write(buffer)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}

}
