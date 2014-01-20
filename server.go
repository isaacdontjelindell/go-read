package main

import (
    "os"
    "net"
    "bufio"
)

func main() {
    //s := string(line[:])
    println("Starting...")

    listener, err := net.Listen("tcp", "0.0.0.0:1337")

    if err != nil {
        println("Error starting listener", err.Error())
        os.Exit(1)
    }


    for {
        conn, err := listener.Accept()
        if err != nil {
            println("Error accepting connection", err.Error())
            return
        }

        go printData(conn)
    }

}

func printData(conn net.Conn) {
    b := bufio.NewReader(conn)

    for {
        line, err := b.ReadBytes('\n')
        if err != nil {
            break
        }
        print(string(line[:]))
        conn.Write(line)
    }
}
