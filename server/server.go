package TCPChat

import (
	getFunc "TCPChat/getFunctions"
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	Name string
	Conn net.Conn
}

type Notification struct {
	Text string
	Addr net.Conn
}

type Message struct {
	Time      string
	SendrAddr net.Conn
	Text      string
}

//create a map of client with name as key and connection as value
var Clients = make(map[string]Client)

// create channels for communicating messages
var Join = make(chan Notification)
var Leave = make(chan Notification)
var Messages = make(chan Message)
var NumOfConn int

// combining the broadcast into one function
func Broadcast(conn net.Conn) {
	for {
		select {
		case msg := <-Join:
			Status("<-join", msg)
		case msg := <-Messages:
			for _, cl := range Clients {
				if msg.SendrAddr != cl.Conn {
					fmt.Fprintln(cl.Conn, "\n"+"["+getFunc.CurrentTime()+"]"+msg.Text)
					fmt.Fprint(cl.Conn, "["+getFunc.CurrentTime()+"]"+"["+cl.Name+"]: ")
				}
			}
		case msg := <-Leave:
			NumOfConn--
			Status("<-leave", msg)
		}
	}
}

func Status(s string, msg Notification) {
	for _, cl := range Clients {
		if msg.Addr != cl.Conn {
			fmt.Fprintln(cl.Conn, "\n"+msg.Text)
			fmt.Fprint(cl.Conn, "["+getFunc.CurrentTime()+"]"+"["+cl.Name+"]: ")
		} else {
			fmt.Fprint(cl.Conn)
		}
	}
}

// function to get the name of the client
func GetName(conn net.Conn) (string, error) {
	for {
		fmt.Fprint(conn, "Enter your name: ")
		name, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return name, err
		}
		name = strings.Trim(name, "\r\n")
		if name == "" {
			continue
		}

		Clients[name] = Client{name, conn}
		if len(getFunc.ReadFile("./log/history.txt")) > 0 {
			fmt.Fprintln(conn, getFunc.ReadFile("./log/history.txt"))
		}
		return name, nil
	}
}
