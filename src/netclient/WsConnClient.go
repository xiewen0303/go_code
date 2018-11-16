package netclient

import (
	"fmt"
	"golang.org/x/net/websocket"
	//"log"
)

type WsConnClient struct {
	Ip string
	Port int32
}

const CONN_CONST string = "ws://%s:%d/"

func (self *WsConnClient) getConnStr() string{
	return fmt.Sprintf(CONN_CONST,self.Ip,self.Port)
}


func echoHandler(ws *websocket.Conn){
	//msg := make([]byte,512)
	//n,err := ws.Read(msg)
	//if err != nil{
	//	log.Fatal(err)
	//}
	//fo


}

func (self *WsConnClient) Conn(ip string,port int32){
	//a := websocket.Handler()
	//
	//self.getConnStr()


}