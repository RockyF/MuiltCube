package core

import (
	"tbs"
	"fmt"
	"model"
)

type Client struct {
	id            int32
	playerVo      *model.PlayerInfoVo
	socket        *tbs.WebSocket
	dispatcher    *tbs.Dispatcher
	protocol      *Protocol
	cbReceiveData *tbs.EventCallback
	cbClosed      *tbs.EventCallback
}

func createClient(socket *tbs.WebSocket) *Client {
	client := &Client{id: int32(1001), socket: socket}
	client.init()
	return client
}

func (this *Client) init() {
	this.protocol = SharedProtocol()

	this.dispatcher = tbs.SharedDispatcher()

	var cb1 tbs.EventCallback = this.onData
	this.cbReceiveData = &cb1
	var cb2 tbs.EventCallback = this.onClosed
	this.cbClosed = &cb2

	this.dispatcher.AddEventListener(tbs.ReceiveData, this.cbReceiveData)
	this.dispatcher.AddEventListener(tbs.Closed, this.cbClosed)
}

func (this *Client) destroy() {
	this.dispatcher.RemoveEventListener(tbs.Closed, this.cbClosed)
	this.dispatcher.RemoveEventListener(tbs.ReceiveData, this.cbReceiveData)

	delete(SharedGameMain().ClientMap, this.id)
	fmt.Println(SharedGameMain().ClientMap)
}

func (this *Client) SendPack(cmd int, body tbs.Message) {
	var msg tbs.Message = make(tbs.Message)
	msg["cmd"] = cmd
	msg["body"] = body

	this.Send(msg)
}

func (this *Client) Send(msg interface{}) {
	this.socket.Send(msg)
}

func (this *Client) Broadcast(cmd int, body tbs.Message) {
	for _, client := range SharedGameMain().ClientMap {
		client.SendPack(cmd, body)
	}
}

func (this *Client) onClosed(event *tbs.Event) {
	sign := (event.Params["sign"]).(int)

	if sign != this.socket.Sign {
		return
	}

	fmt.Printf("[#%d] closed\n", this.socket.Sign)

	this.destroy()
}

func (this *Client) onData(event *tbs.Event) {
	sign := (event.Params["sign"]).(int)

	if sign != this.socket.Sign {
		return
	}

	fmt.Printf("[#%d]:get data\n", this.socket.Sign)

	msg := event.Params["data"].(tbs.Message)

	//fmt.Println(msg)
	this.protocol.Deal(this, msg)
}
