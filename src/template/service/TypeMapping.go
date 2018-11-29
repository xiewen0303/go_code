package service

import (
	"fileio"
	"fmt"
)

var KeyMappings = make(map[string]string)

func LoadJsonFile() {
	fileio.LoadJson("./tpl/typemapping.json",KeyMappings)
}

func GetJavaType(sqlType string) string {
	data := KeyMappings[sqlType]
	if data == ""{
		fmt.Printf("数据类型不存在,sqlType=%s\n",sqlType)
	}
	return data
}
