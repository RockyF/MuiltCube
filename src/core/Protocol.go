package core

import (
	"tbs"
	"fmt"
	"model"
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

func (this *Protocol) Deal(client *Client, obj tbs.Message) {
	cmd := obj["cmd"];
	fmt.Printf("cmd[%d]\n", int(cmd.(float64)))
	body := obj["body"];
	_, err := tbs.Call(this, fmt.Sprintf("Deal%d", int(cmd.(float64))), client, body)
	if err != nil {
		fmt.Println(err)
	}
}

func (this *Protocol) Deal1001(client *Client, body map[string]interface {}) {
	id := body["id"]
	pwd := body["pwd"]

	result := model.SharedPlayerModel().Login(int64(id.(float64)), pwd.(string))

	var responseBody tbs.Message = make(tbs.Message)

	var resultCode int
	if result{
		resultCode = 0
	}else{
		resultCode = 1
	}
	responseBody["result"] = resultCode;

	client.SendPack(1001, responseBody)
}
