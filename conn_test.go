package main

import (
	"net"
	"testing"
)

func TestConnection(t *testing.T) {

	for i := 0; i < 100; i++ {
		conn, err := net.Dial("tcp4", "127.0.0.1:8080")
		if err != nil {
			t.Errorf("%s", err)
		}

		conn.Write([]byte("hi"))
		buf := make([]byte, 2)
		conn.Read(buf)

		response := string(buf)

		if response != "hi" {
			t.Error("Wrond response")
			break
		}
	}

}
