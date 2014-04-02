package core

import (
	"fmt"
	"tbs"
	"model"
	"code.google.com/p/go.net/websocket"
	"net/http"
	"log"
)

type GameMain struct {
	port         int
	serverSocket *tbs.ServerSocket
	dispatcher   *tbs.Dispatcher

	ClientMap      map[int32]*Client
}

var instanceGameMain *GameMain

func SharedGameMain() *GameMain {
	if instanceGameMain == nil {
		instanceGameMain = &GameMain{}
		instanceGameMain.init()
	}
	return instanceGameMain
}

func (this *GameMain) init() {
	model.SharedPlayerModel()

	http.Handle("/", websocket.Handler(this.acceptHandler))

	this.ClientMap = make(map[int32]*Client)

	this.dispatcher = tbs.SharedDispatcher()
}

func (this *GameMain) acceptHandler(ws *websocket.Conn) {
	socket := tbs.CreateWebSocket(ws)

	client := createClient(socket)
	this.ClientMap[client.id] = client

	fmt.Printf("[#%d]client connect on:%s\n", socket.Sign, socket.Conn.RemoteAddr().String())

	socket.Launch();
}

func (this *GameMain) Start(port int) {
	this.port = port

	fmt.Println("listen on port ", this.port)
	strPort := fmt.Sprintf(":%d", this.port)
	if err := http.ListenAndServe(strPort, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
