package service

import (
	"text/template"
	"log"
	"template/dbcontent"
	"io"
	//"github.com/go-sql-driver/mysql"
	"ms"
	"strings"
)

type TemplateUtil struct {

}

func (self *TemplateUtil) Parse(tmpContent *string,wr io.Writer,contentData dbcontent.IContentData) *string {

	tmpl, error := template.New("dbBeanTemp").Parse(*tmpContent)
	if error != nil {
		log.Println(error)
	}
	error = tmpl.Execute(wr, contentData)
	if error != nil {
		panic(error)
	}
	result := "success"
	return &result
}

func (self *TemplateUtil) GetDbContentData() dbcontent.IContentData {
	dbConnectionUtil :=  &ms.DBConnectionUtil{}
	dbConnectionUtil.GetConn()
	return  nil
}

const MaxASCII = '\u007F'
func toUpper(r rune) rune {
	if r <= MaxASCII {
		if 'a' <= r && r <= 'z' {
			r -= 'a' - 'A'
		}
		return r
	}
	return r
}

/**
 *  转换成驼峰命名规则（遇到下划线就大写）
 */
func (self *TemplateUtil) ConverParamName(source string) string {

	var result string
	//首字母
	if source == "" {
		return source
	}

	//去空格
	source = strings.Replace(source," ","",-1)

	cells := strings.Split(source,"_")
	result += cells[0]
	for index:=1; index < len(cells); index++  {
		cell := cells[index]
		data := string(cell[0:1])
		result += string(toUpper(([]rune(data))[0]))

		if len(cell) > 1 {
			result += string(cell[1:])
		}
	}
	return result
}

/**
 * 转换成首字母大写，遇到下划线就大写
 */
func (self *TemplateUtil) ConverAllUpName(source string) string {

	var result string
	//首字母
	if source == "" {
		return source
	}

	//去空格
	source = strings.Replace(source," ","",-1)

	cells := strings.Split(source,"_")

	for _,cell := range cells {

		data := string(cell[0:1])
		result += string(toUpper(([]rune(data))[0]))

		if len(cell) > 1 {
			result +=string(cell[1:])
		}
	}
	return result
}
//GetDBContent

func (self *TemplateUtil) GetDBBeanContent() *dbcontent.DBBeanContent {
	dbBeanContent := &dbcontent.DBBeanContent{}

	dbConnectionUtil := &ms.DBConnectionUtil{}
	columnContent := dbConnectionUtil.GetDBContent()

	dbBeanContent.TableName = ms.TABLE_NAME
	dbBeanContent.ClazzName = self.ConverAllUpName(dbBeanContent.TableName)
	dbBeanContent.PackageName = "com.junyou.dbbean"

	beanProps := make([]dbcontent.BeanProp,0)
	for _,val := range columnContent  {

		benProp1 := dbcontent.BeanProp{}
		benProp1.ColumnName = val.Column_name

		//去除扩后之后的数据
		index := strings.Index(val.Column_type,"(")
		if index > 0 {
			val.Column_type = string(val.Column_type[0:index])
		}

		benProp1.TypeClazz =  GetJavaType(val.Column_type)
		benProp1.PropName = self.ConverParamName(val.Column_name)
		benProp1.ClazzName = self.ConverAllUpName(benProp1.PropName)
		benProp1.FlagPRI = val.Column_key == "PRI"

		beanProps = append(beanProps,benProp1)


	}

	dbBeanContent.BeanProps = beanProps
	return dbBeanContent
}