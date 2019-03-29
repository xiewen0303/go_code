package main

import (
	"netclient"
	"message"
	"github.com/golang/protobuf/proto"
	"fmt"
	//"template"
	"fileio"
	//"template/service"
	"ms"
)

func main() {
	//testTemplate()
	//testWrite()
	//testTemplate()
	//testSql()
	//testJson()
	//fmt.Printf("%f",3.4*4-4+0.8)
	////println(service.GetJavaType("varchar112"))


	//service.LoadJsonFile()
	//templateUtil := service.TemplateUtil{}
	//templateUtil.GenCode()

	fmt.Print(2860+800+748)
}

func testSql(){
	connectionUtil := ms.DBConnectionUtil{}
	connectionUtil.GetConn()
}


func testWrite() {
	readUtil := &fileio.FileUtil{}
	readUtil.WriteFile()
}

func testWebSocket() {
	var client2 netclient.Test1
	client2.Test1()
	loginReq := message.Login_C{AccountId:proto.String("zhanshang")}
	fmt.Println(loginReq.GetAccountId())
}
