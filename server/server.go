package main

import (
	"bufio"
	"net"
	"os"
)

func main() {
	println("Starting...")

	listener, err := net.Listen("tcp", "0.0.0.0:1337")
	if err != nil {
		println("Error starting listener", err.Error())
		os.Exit(1)
	}

    c := make(chan []byte)

    // accept connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			println("Error accepting connection", err.Error())
			return
		}

        go acceptData(conn, c)
        go sendData(conn, c)
	}
}

func acceptData(conn net.Conn, c chan []byte) {
	b := bufio.NewReader(conn)

	for {
		line, err := b.ReadBytes('\n')
		if err != nil {
			break
		}
		print(string(line[:]))
        c <- line
		//conn.Write(line)
	}
}

func sendData(conn net.Conn, c chan []byte) {
    for {
        val := <-c
        conn.Write(val)
    }
}
