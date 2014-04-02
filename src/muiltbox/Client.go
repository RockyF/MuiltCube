package muiltbox

import (
	"tbs"
	"fmt"
	"muiltbox/model"
)

type Client struct {
	id            int32
	playerVo      *model.PlayerVo
	socket        *tbs.Socket
	dispatcher    *tbs.Dispatcher
	packLen       int32
	protocol      *Protocol
	cbReceiveData *tbs.EventCallback
	cbClosed      *tbs.EventCallback
}

func createClient(socket *tbs.Socket) *Client {
	playerVo := SharedPlayerModel().Get(1001)
	client := &Client{id: int32(playerVo.Id), playerVo: playerVo, socket: socket}
	client.init()
	return client
}

func (this *Client) init() {
	this.packLen = 0

	this.protocol = SharedProtocol()

	this.dispatcher = tbs.SharedDispatcher()

	var cb1 tbs.EventCallback = this.onData
	this.cbReceiveData = &cb1
	var cb2 tbs.EventCallback = this.onClosed
	this.cbClosed = &cb2
	fmt.Println(this.cbReceiveData)
	fmt.Println(this.cbClosed)

	this.dispatcher.AddEventListener(tbs.ReceiveData, this.cbReceiveData)
	this.dispatcher.AddEventListener(tbs.Closed, this.cbClosed)
}

func (this *Client) destroy() {
	this.dispatcher.RemoveEventListener(tbs.Closed, this.cbClosed)
	this.dispatcher.RemoveEventListener(tbs.ReceiveData, this.cbReceiveData)

	delete(SharedGameMain().ClientMap, this.id)
	fmt.Println(SharedGameMain().ClientMap)
	this.protocol.Deal1012(this)
}

func (this *Client) SendPack(cmd int, bytes []byte) {
	buffer := tbs.CreateByteArray([]byte{})
	bufferIn := tbs.CreateByteArray([]byte{})

	bufferIn.WriteInt32(int32(cmd))
	bufferIn.WriteBytes(bytes)

	buffer.WriteInt32(int32(len(bufferIn.Bytes())))
	buffer.Write(bufferIn.Bytes())

	this.Send(buffer.Bytes())

	//fmt.Println("send:", len(bufferIn.Bytes()))
}

func (this *Client) Send(bytes []byte) {
	this.socket.Write(bytes)
	this.socket.Flush()
}

func (this *Client) Broadcast(cmd int, bytes []byte) {
	for _, client := range SharedGameMain().ClientMap {
		client.SendPack(cmd, bytes)
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

func (this *Client) GetBuffer() *tbs.ByteArray {
	return this.socket.BufferIn
}

func (this *Client) onData(event *tbs.Event) {
	sign := (event.Params["sign"]).(int)

	if sign != this.socket.Sign {
		return
	}

	length := (event.Params["length"]).(int)

	fmt.Printf("[#%d]:get data, len = %d\n", this.socket.Sign, length)

	buffer := this.socket.BufferIn

	if this.packLen == 0 && buffer.Length() >= 4 {
		this.packLen, _ = buffer.ReadInt32()
		fmt.Println("read packLen:", this.packLen)
	}

	if buffer.Length() >= int(this.packLen) {
		temp := make([]byte, this.packLen)
		copy(temp, buffer.BytesAvailable())
		this.protocol.Deal(this, temp)

		buffer.Reset()
		this.packLen = 0
	}
}
