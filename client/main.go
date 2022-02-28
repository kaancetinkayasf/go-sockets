package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

const (
	TextMessage = 1
	JsonMessage = 2
)

func main() {

	ConnectToServer()

}

var messageBuf []byte

func ConnectToServer() {
	connection, err := net.Dial("tcp4", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	handleConnection(connection)
}

func handleConnection(connection net.Conn) {

	go func() {
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

		}
	}()

	for {

		consoleReader := bufio.NewReader(os.Stdin)
		input, _ := consoleReader.ReadString('\n')

		msg := createMessage(TextMessage, input)
		_, err := connection.Write(msg)
		if err != nil {
			fmt.Println(err)
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
