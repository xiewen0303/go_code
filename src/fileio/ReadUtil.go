package fileio

import (
	"io/ioutil"
	"log"
	"fmt"
)

/**
 * 读取文件的工具类
 */
type FileUtil struct {

	//需要读取的文件名称
	filePaths []*string
}

func (self *FileUtil) setFilePaths(filePaths []*string){
	self.filePaths = filePaths
}

func (self *FileUtil) addFilePath(filePath *string){
	self.filePaths = append(self.filePaths,filePath)
}

func (self *FileUtil) ReadFile(filePath string) string {

	data,err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("file is not exits,filePath=%s",filePath)
		return ""
	}
	return string(data[:])
}

func (self *FileUtil) WriteFile() bool {
	name := "a.txt"
	content := "will write content"

	data :=[]byte(content)


	if ioutil.WriteFile(name,data,0644) == nil {
		fmt.Println("写入文件成功:",content)
	}

	return true
}



