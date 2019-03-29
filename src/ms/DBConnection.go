package ms

import (
	"database/sql"
	"log"
	"template/dbcontent"
	_ "github.com/go-sql-driver/mysql"
	"fileio"
	"fmt"
)

type DBConnectionUtil struct {
	DBConfig fileio.DBConfig
}

func (self *DBConnectionUtil) GetConn() *sql.DB {

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",self.DBConfig.UserName,self.DBConfig.Password,self.DBConfig.DbIp,self.DBConfig.DbPort,self.DBConfig.DbName)

	conn,err := sql.Open("mysql",connStr)
	if err != nil {
		log.Println(err)
	}
	return conn
}

//const TABLE_NAME = "email"

/**
 * 获取数据表信息
 */
func (self *DBConnectionUtil) GetDBContent() []dbcontent.ColumnContent {
	result := make([]dbcontent.ColumnContent,0)

	conn := self.GetConn()

	//select TABLE_NAME,TABLE_COMMENT from information_schema.tables where table_schema='shenzhuan_server'   and TABLE_NAME='email'
	sql := "select column_name,COLUMN_TYPE,COLUMN_KEY,COLUMN_COMMENT from information_schema.columns where table_schema=? and table_name=?"

	rowDatas,err := conn.Query(sql,"shenzhuan_server","email")
	//rowDatas,err := conn.Query(sql,self.DBConfig.DbName,self.DBConfig.TableName)
	if err != nil {
		log.Println(err)
	}

	for rowDatas.Next() {
		column := dbcontent.ColumnContent{}
		err = rowDatas.Scan(&column.Column_name,&column.Column_type,&column.Column_key,&column.Column_comment)
		if err != nil {
			log.Print(err)
			continue
		}
		//fmt.Printf("字段名：%s,\t\t\t字段类型：%s,\t\t\t是否为主键：%s,\t\t\t字段注释：%s",column.Column_name,column.Column_type,column.Column_key,column.Column_comment)
		//fmt.Println("")
		result = append(result,column)
	}

	self.CloseConn(conn)
	return result
}

func (self *DBConnectionUtil) CloseConn(conn *sql.DB){
	if conn != nil {
		conn.Close()
	}
}