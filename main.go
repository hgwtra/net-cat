package main

import (
	client "TCPChat/clients"
	getFunc "TCPChat/getFunctions"
	handler "TCPChat/server"
	"fmt"
	"net"
	"os"
)

func main() {
	getFunc.EmptyHistory()

	// Step 1: Check validity of port
	port, valid := getFunc.CheckPort()
	if !valid {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	// Step 2: Listen with valid port
	go getFunc.Exit()
	port = ":" + port
	li, err := net.Listen("tcp", port)
	getFunc.CheckError(err, "ERROR: unable to listen to port "+port)
	defer li.Close()
	fmt.Println("Listening on the port " + port)

	for {
		// Step 3: Accept and handle up to 10 connections with go routine
		conn, err := li.Accept()
		if err != nil {
			os.Exit(0)
		}
		go handNoOfConn(conn)
	}
}

func handNoOfConn(conn net.Conn) {
	if handler.NumOfConn >= 10 {
		fmt.Fprintln(conn, getFunc.ReadFile("./log/sorry.txt"))
		conn.Close()
	} else {
		fmt.Fprintln(conn, getFunc.ReadFile("./log/welcome.txt"))
		handler.NumOfConn++
		go client.HandleClient(conn)
		go handler.Broadcast(conn)
	}

	//List maximum number of connections
	//fmt.Println(" conn no: ", handler.NumOfConn, " ", conn)
}
