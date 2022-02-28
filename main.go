package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

var connectionList []net.Conn

var messageBuf []byte

func main() {
	StartServer()
}

func StartServer() {
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

		header := make([]byte, 8)
		_, err := connection.Read(header[:])
		if err != nil {
			fmt.Println(err)
			connection.Close()
			break
		}

		mlen := binary.LittleEndian.Uint32(header[4:])
		dataBuffer := make([]byte, mlen)
		_, err = connection.Read(dataBuffer[:])
		if err != nil {
			fmt.Println(err)
			connection.Close()
			break
		}

		messageBuf = append(messageBuf, header...)
		messageBuf = append(messageBuf, dataBuffer...)

		mtype, mlen, msg := readMessage(messageBuf)
		messageBuf = nil
		fmt.Printf("%d %d %s\n", mtype, mlen, msg)

		for _, client := range connectionList {
			msg := createMessage(int(mtype), msg)
			_, err := client.Write(msg)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}

}

func createMessage(mtype int, data string) []byte {
	buf := make([]byte, 4+4+len(data))
	binary.LittleEndian.PutUint32(buf[0:], uint32(mtype))
	binary.LittleEndian.PutUint32(buf[4:], uint32(len(data)))
	copy(buf[8:], []byte(data))

	return buf

}

func readMessage(data []byte) (mtype, mlen uint32, msg string) {
	mtype = binary.LittleEndian.Uint32(data[0:])
	mlen = binary.LittleEndian.Uint32(data[4:])
	msg = string(data[8:])

	return mtype, mlen, msg
}
