package muiltbox

import (
	"tbs"
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"model"
)

type GameMain struct {
	port         int
	serverSocket *tbs.ServerSocket
	dispatcher   *tbs.Dispatcher

	ClientMap map[int32]*Client
}

var instance *GameMain

func SharedGameMain() *GameMain {
	if instance == nil {
		instance = &GameMain{}
		instance.init()
	}
	return instance
}

func (this *GameMain) init() {
	m := md5.New()
	m.Write([]byte("jsj092"))
	println(hex.EncodeToString(m.Sum(nil)))

	model.SharedPlayerModel()

	this.serverSocket = tbs.CreateServerSocket()

	this.ClientMap = make(map[int32]*Client)

	this.dispatcher = tbs.SharedDispatcher()

	var cb1 tbs.EventCallback = this.onServerStarted
	var cb2 tbs.EventCallback = this.onAccept

	this.dispatcher.AddEventListener(tbs.ServerStarted, &cb1)
	this.dispatcher.AddEventListener(tbs.Accept, &cb2)
}

func (this *GameMain) Start(port int) {
	this.port = port

	this.serverSocket.Start(this.port)
}

func (this *GameMain) onServerStarted(event *tbs.Event) {
	fmt.Println("server started.")
}

func (this *GameMain) onAccept(event *tbs.Event) {
	socket := (event.Params["socket"]).(*tbs.Socket)

	client := createClient(socket)
	this.ClientMap[client.id] = client

	fmt.Printf("[#%d]client connect on:%s\n", socket.Sign, socket.Conn.RemoteAddr().String())
}
