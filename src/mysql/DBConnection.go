package mysql

import (
	"database/sql"
	"log"
)

type DBConnectionUtil struct {

}

func (self *DBConnectionUtil) getConn() {
	conn,err := sql.Open("mysql","root:root@192.168.0.129:3306/shenzhuan_server")
	if err != nil {
		log.Println(err)
	}

	sql := "select column_name,COLUMN_TYPE,COLUMN_KEY,COLUMN_COMMENT from information_schema.columns where table_schema='?' and table_name='?'"

	rowDatas,err := conn.Query(sql,"shenzhuan_server","email")
	if err != nil {
		log.Println(err)
	}

	

	if rowDatas.Next() {
		rowDatas.Scan()
	}
}

