package main

import (
	"fmt"
	"net"
	"testing"
)

func ConnectionTest(t *testing.T) {

	StartServer()

	for i := 0; i < 10000; i++ {

		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			t.Error(err)
		}

		fmt.Println(conn.RemoteAddr().String())
	}

}
