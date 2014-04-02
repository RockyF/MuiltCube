package tbs

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
)

type Message map[string]interface {}

var signWebSocket int = 0
type WebSocket struct {
	Conn       *websocket.Conn
	Sign       int
	dispatcher *Dispatcher
}

func CreateWebSocket(conn *websocket.Conn) *WebSocket {
	socket := &WebSocket{Conn: conn, Sign: signWebSocket}
	signWebSocket ++
	return socket
}

func (this *WebSocket) Launch() {
	this.dispatcher = SharedDispatcher()

	this.handle()
}

func (this *WebSocket) handle() {
	var err error
	for {
		var msg Message
		if err = websocket.JSON.Receive(this.Conn, &msg); err != nil {
			fmt.Println("Can't receive:", err)
			break
		}

		params := make(map[string]interface{})
		params["sign"] = this.Sign
		params["data"] = msg
		this.dispatcher.DispatchEvent(CreateEvent(ReceiveData, params))
	}
}

func (this *WebSocket) Send(msg interface {}) {
	var err error
	if err = websocket.JSON.Send(this.Conn, &msg); err != nil {
		fmt.Println("Can't send");
	}
}

func (this *WebSocket) onClosed() {
	params := make(map[string]interface{})
	params["sign"] = this.Sign
	this.Conn.Close()
	this.dispatcher.DispatchEvent(CreateEvent(Closed, params))
}

func (this *WebSocket) Close() {
	this.Conn.Close()
}
