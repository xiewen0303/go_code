package main

import (
	"netclient"
	"message"
	"github.com/golang/protobuf/proto"
	"fmt"
	//"template"
	"fileio"
	"template/service"
	"os"
	"log"
	"ms"
)

func main() {
	//testTemplate()
	//testWrite()

	//testTemplate()
	//testSql()
	//testJson()

	service.LoadJsonFile()
	//println(service.GetJavaType("varchar112"))
	testdbBeanGen()
}

func testSql(){
	connectionUtil := ms.DBConnectionUtil{}
	connectionUtil.GetConn()
}


func testdbBeanGen(){

	fileUtil := &fileio.FileUtil{}
	templateUtil  := service.TemplateUtil{}
	contentData := templateUtil.GetDBBeanContent()

	contentStr := fileUtil.ReadFile("./tpl/dbbean.tpl")

	name := "Email.java"
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Println(err)
	}

	templateUtil.Parse(&contentStr,f,contentData)
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
