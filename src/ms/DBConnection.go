package ms

import (
	"database/sql"
	"log"
	"template/dbcontent"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DBConnectionUtil struct {

}

func (self *DBConnectionUtil) GetConn() {
	conn,err := sql.Open("mysql","root:root@tcp(192.168.0.158:3306)/shenzhuanserver?charset=utf8")
	if err != nil {
		log.Println(err)
	}

	sql := "select column_name,COLUMN_TYPE,COLUMN_KEY,COLUMN_COMMENT from information_schema.columns where table_schema=? and table_name=?"

	//rowDatas,err := conn.Query(sql,"shenzhuan_server","email")
	rowDatas,err := conn.Query(sql,"shenzhuanserver","email")
	if err != nil {
		log.Println(err)
	}

	for rowDatas.Next() {
		column := dbcontent.ColumnContent{}
		err = rowDatas.Scan(&column.Column_name,&column.Column_type,&column.Column_key,&column.Column_comment)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("字段名：%s,\t\t\t字段类型：%s,\t\t\t是否为主键：%s,\t\t\t字段注释：%s",column.Column_name,column.Column_type,column.Column_key,column.Column_comment)
		fmt.Println("")
	}
}


func (self *DBConnectionUtil) CloseConn(conn *sql.DB){
	if conn != nil {
		conn.Close()
	}
}

