package tbs

import (
	//"fmt"
	"net"
	"strconv"
)

const ServerStarted = "ServerStarted"
const ServerError = "ServerError"
const ServerClosed = "ServerClosed"
const Accept = "Accept"
const Closed = "Closed"
const ReceiveData = "ReceiveData"

//==========ServerSocket
type ServerSocket struct {
	port       int
	started    bool
	listener   *net.TCPListener
	dispatcher *Dispatcher
}

func CreateServerSocket() *ServerSocket {
	instance := &ServerSocket{}
	instance.dispatcher = SharedDispatcher()
	return instance
}

func (this *ServerSocket) Start(port int) error {
	if this.started {
		return nil
	}

	this.port = port
	service := ":" + strconv.Itoa(port)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		return err
	}

	this.listener, err = net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}

	this.started = true
	this.dispatcher.DispatchEvent(CreateEvent(ServerStarted, nil))

	sign := 0

	for {
		conn, err := this.listener.Accept()
		if err != nil {
			continue
		}

		socket := CreateSocket(conn, sign)
		socket.launch()

		params := make(map[string]interface{})
		params["socket"] = socket
		this.dispatcher.DispatchEvent(CreateEvent(Accept, params))

		sign++
	}
}

func (this *ServerSocket) stop() {
	if this.started {
		this.listener.Close()
		this.dispatcher.DispatchEvent(CreateEvent(ServerClosed, nil))
	}

	this.started = false
}

//===============Socket
type Socket struct {
	Conn       net.Conn
	Sign       int
	dispatcher *Dispatcher

	BufferIn  *ByteArray
	BufferOut *ByteArray
}

func CreateSocket(conn net.Conn, sign int) *Socket {
	bufferIn := CreateByteArray([]byte{})
	bufferOut := CreateByteArray([]byte{})

	socket := &Socket{Conn: conn, Sign: sign, BufferIn: bufferIn, BufferOut: bufferOut}
	return socket
}

func (this *Socket) launch() {
	this.dispatcher = SharedDispatcher()

	go this.handle()
}

func (this *Socket) handle() {
	buffer := make([]byte, 256)

	for {
		readLen, err := this.Conn.Read(buffer)

		if err != nil {
			this.onClosed()
			break
		}

		if readLen == 0 {
			this.onClosed()
			break
		}

		bytesGet := make([]byte, readLen)
		copy(bytesGet, buffer[:readLen])

		this.BufferIn.WriteBytes(bytesGet)

		params := make(map[string]interface{})
		params["sign"] = this.Sign
		params["length"] = readLen
		this.dispatcher.DispatchEvent(CreateEvent(ReceiveData, params))
	}
}

func (this *Socket) Write(bytes []byte) {
	this.BufferOut.WriteBytes(bytes)
}

func (this *Socket) Flush() {
	this.Conn.Write(this.BufferOut.Bytes())

	this.BufferOut.Reset()
}

func (this *Socket) onClosed() {
	params := make(map[string]interface{})
	params["sign"] = this.Sign
	this.Conn.Close()
	this.dispatcher.DispatchEvent(CreateEvent(Closed, params))
}

func (this *Socket) Close() {
	this.Conn.Close()
}
