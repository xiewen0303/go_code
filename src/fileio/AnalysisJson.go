package fileio

import (
	"encoding/json"
	"fmt"
	"log"
)

type MongoConfig struct {
	MongoAddr      string
	MongoPoolLimit int
	MongoDb        string
	MongoCol       string
}

/**
 *  "addr":"27017",
 *	"mongo":{
 */
type CfgData struct {
	Addr string
	Mongo MongoConfig
}

/**
 * 数据库配置文件
 */
type DBConfig struct {
	DbIp string
	DbPort int
	DbName string
	TableName string
	UserName string
	Password string
	PackageName string
}

func LoadDBConfig() DBConfig {
	var dbConfig DBConfig

	fileioUtil := FileUtil{}
	content := fileioUtil.ReadFile("./tpl/config.json")

	err := json.Unmarshal([]byte(content),&dbConfig)
	if err != nil {
		fmt.Printf("解析json错误,jsonStr=%s\n",content)
		log.Fatal(err)
	}
	return dbConfig
}

func TestLoadJsonCfg(){
	var mongoConfig CfgData

	fileioUtil := FileUtil{}
	content := fileioUtil.ReadFile("./tpl/config.json")

	err := json.Unmarshal([]byte(content),&mongoConfig)
	if err != nil {
		fmt.Printf("解析json错误,jsonStr=%s\n",content)
		log.Fatal(err)
	}

	fmt.Println(mongoConfig.Addr)
}

func OutJsonStr(){
	mongoConfig := CfgData{Addr:"192.168.0.158",Mongo:MongoConfig{MongoAddr:"192.168.0.158",MongoDb:"db",MongoPoolLimit:30}}
	b, err := json.Marshal(mongoConfig)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}



func LoadJson(path string,data map[string]string) {
	fileioUtil := FileUtil{}
	content := fileioUtil.ReadFile(path)

	err := json.Unmarshal([]byte(content),&data)
	if err != nil {
		fmt.Printf("解析json错误,jsonStr=%s\n",content)
		log.Fatal(err)
	}
}

