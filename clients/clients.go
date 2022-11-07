package TCPChat

import (
	getFunc "TCPChat/getFunctions"
	handler "TCPChat/server"
	"bufio"
	"fmt"
	"net"
)

func HandleClient(conn net.Conn) {
	// send join notification to all clients via channel

	name, err := handler.GetName(conn)
	if err != nil {
		delete(handler.Clients, name)
		handler.NumOfConn--
		return
	}

	handler.Join <- handler.Notification{Text: name + " has joined our chat...", Addr: conn}
	msg := bufio.NewScanner(conn)
	fmt.Fprint(conn, fmt.Sprintf("[%s]", getFunc.CurrentTime())+"["+name+"]: ")
	for msg.Scan() {
		text := msg.Text()
		fmt.Fprint(conn, fmt.Sprintf("[%s]", getFunc.CurrentTime())+"["+name+"]: ")
		if text == "" {
			continue
		}
		// send message to all clients via channel
		handler.Messages <- handler.Message{Time: fmt.Sprintf("[%s]", getFunc.CurrentTime()), SendrAddr: conn, Text: "[" + name + "]: " + text}
		getFunc.AddToFile(fmt.Sprint(fmt.Sprintf("[%s]", getFunc.CurrentTime()), "["+name+"]: "+text))
	}
	// send leaving notification to all clients via channel
	handler.Leave <- handler.Notification{Text: name + " has left our chat...", Addr: conn}
	delete(handler.Clients, name)
	conn.Close()
}
