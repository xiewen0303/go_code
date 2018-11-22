package main

import (
	"netclient"
	"message"
	"github.com/golang/protobuf/proto"
	"fmt"
	//"template"
	"fileio"
	"template/service"
	"template/dbcontent"
	"os"
	"log"
	"ms"
	"analysis/jsonutil"
)

func main() {
	//testTemplate()
	//testWrite()
	//testdbBeanGen()
	//testTemplate()
	//testSql()
	//testJson()
	testToJsonStr()
}
func testToJsonStr(){
	jsonutil.OutJsonStr()
}

func testJson(){
	jsonutil.TestLoadJsonCfg()
}


func testSql(){
	connectionUtil := ms.DBConnectionUtil{}
	connectionUtil.GetConn()
}


func testdbBeanGen(){

	fileUtil := &fileio.FileUtil{}
	templateUtil  := service.TemplateUtil{}
	contentData  := &dbcontent.DBBeanContent{PackageName:"dbbean",TableName:"student"}

	beanProps := make([]dbcontent.BeanProp,0)
	benProp1 := dbcontent.BeanProp{PropName:"age",TypeClazz:"int"}
	benProp2 := dbcontent.BeanProp{PropName:"name",TypeClazz:"String"}

	contentData.BeanProps = append(beanProps,benProp1,benProp2)

	contentStr := fileUtil.ReadFile("./tpl/dbbean.tpl")

	name := "Student.java"
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

func testTemplate(){
	service.TestGen()
}