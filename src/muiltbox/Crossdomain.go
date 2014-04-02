package muiltbox

import (
	"tbs"
	"fmt"
)

type Crossdomain struct {
	port         int
	serverSocket *tbs.ServerSocket
	dispatcher   *tbs.Dispatcher
}

var instanceCrossdomain *Crossdomain

func SharedCrossdomain() *Crossdomain {
	if instance == nil {
		instanceCrossdomain = &Crossdomain{}
		instanceCrossdomain.init()
	}
	return instanceCrossdomain
}

func (this *Crossdomain) init() {
	this.serverSocket = tbs.CreateServerSocket()
	this.dispatcher = tbs.SharedDispatcher()

	var cb1 tbs.EventCallback = this.onServerStarted
	var cb2 tbs.EventCallback = this.onAccept

	this.dispatcher.AddEventListener(tbs.ServerStarted, &cb1)
	this.dispatcher.AddEventListener(tbs.Accept, &cb2)
}

func (this *Crossdomain) Start(port int) {
	this.port = port

	this.serverSocket.Start(this.port)
}

func (this *Crossdomain) onServerStarted(event *tbs.Event) {
	fmt.Printf("%d server started.\n", this.port)
}

func (this *Crossdomain) onAccept(event *tbs.Event) {
	socket := (event.Params["socket"]).(*tbs.Socket)
	
	socket.Write([]byte("<?xml version='1.0'?><cross-domain-policy><allow-access-from domain='*' to-ports='8002'/></cross-domain-policy>"))
	socket.Flush()
	socket.Close()
	
	fmt.Printf("[#%d]%d:%s\n", socket.Sign, this.port, socket.Conn.RemoteAddr().String())
}
