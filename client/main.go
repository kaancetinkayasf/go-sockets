package main

import (
	"fmt"
	"net"
)

func main() {

	ConnectToServer()

}

func ConnectToServer() {
	connection, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	handleConnection(connection)
}

func handleConnection(connection net.Conn) {

	go func() {
		for {
			buffer := make([]byte, 8096)
			_, err := connection.Read(buffer[:])
			if err != nil {
				fmt.Println(err)
				connection.Close()
				break
			}

			fmt.Printf("%s\n", buffer)

		}
	}()

	for {
		buffer := make([]byte, 8096)
		fmt.Scanf("%s", &buffer)
		_, err := connection.Write(buffer)
		if err != nil {
			fmt.Println(err)
		}

	}
}
