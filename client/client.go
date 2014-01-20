package main

import (
	"bufio"
	"net"
	"os"
)

func main() {
	println("Connecting to localhost, port 1337...")

	conn, err := net.Dial("tcp", "localhost:1337")

	if err != nil {
		println("Error dialing!", err.Error())
		os.Exit(1)
	}

	for {
		b := bufio.NewReader(os.Stdin)

		for {
			line, err := b.ReadBytes('\n')
			if err != nil {
				break
			}
			//print(string(line[:]))
			conn.Write(line)
		}
	}
}
