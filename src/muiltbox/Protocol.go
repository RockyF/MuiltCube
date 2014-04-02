package muiltbox

import (
	"tbs"
	"fmt"
)

type Protocol struct {
	clients map[int32]*Client
}

var protocolInstance *Protocol

func SharedProtocol() *Protocol {
	if protocolInstance == nil {
		protocolInstance = &Protocol{clients: SharedGameMain().ClientMap}
	}

	return protocolInstance
}

func (p *Protocol) Deal(client *Client, bytes []byte) {

	buffer := tbs.CreateByteArray(bytes)

	cmd, _ := buffer.ReadInt32()
	fmt.Printf("cmd[%d], len[%d]\n", cmd, len(bytes))
	_, err := tbs.Call(p, fmt.Sprintf("Deal%d", cmd), client, buffer)
	if err != nil {
		fmt.Println(err)
	}
}

func (this *Protocol) Deal1000(client *Client, bufferIn *tbs.ByteArray) {
	client.playerVo.Nickname, _ = bufferIn.ReadUTF()

	buffer := tbs.CreateByteArray([]byte{})

	buffer.WriteBool(true)
	buffer.WriteInt32(int32(client.id))

	client.SendPack(1000, buffer.Bytes())
}

func (this *Protocol) Deal1010(client *Client, bufferIn *tbs.ByteArray) {
	buffer := tbs.CreateByteArray([]byte{})

	buffer.WriteBool(true)

	client.SendPack(1010, buffer.Bytes())
}

func (this *Protocol) Deal1011(client *Client, bufferIn *tbs.ByteArray) {
	buffer := tbs.CreateByteArray([]byte{})

	clientMap := SharedGameMain().ClientMap
	buffer.WriteInt32(int32(len(clientMap)))
	for _, client := range clientMap {
		buffer.WriteInt32(client.id)
		buffer.WriteUTF(client.playerVo.Nickname)
		buffer.WriteInt32(0)
		buffer.WriteInt32(client.playerVo.X)
		buffer.WriteInt32(client.playerVo.Y)
	}

	client.Broadcast(1011, buffer.Bytes())
}

func (this *Protocol) Deal1012(client *Client) {
	buffer := tbs.CreateByteArray([]byte{})

	buffer.WriteInt32(int32(client.id))

	fmt.Println(client.id, "closed")

	client.Broadcast(1012, buffer.Bytes())
}

func (this *Protocol) Deal1100(client *Client, bufferIn *tbs.ByteArray) {
	bufferIn.SetReadPos(8)
	x, _ := bufferIn.ReadInt32()
	y, _ := bufferIn.ReadInt32()
	client.playerVo.X = x
	client.playerVo.Y = y

	bufferIn.SetWriteEnd()
	bufferIn.WriteInt32(int32(client.playerVo.Id))

	bufferIn.SetReadPos(4)

	client.Broadcast(1100, bufferIn.BytesAvailable())
}
