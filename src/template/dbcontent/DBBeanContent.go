package dbcontent

import "template"

/**
 * dbbean 生成内容
 */
type DBBeanContent struct {

	PackageName		string   		//包名
	ClazzName		string	 		//类名
	ClazzDescribe	string   		//类名描述
	BeanProps		[]BeanProp 		//主键名字
	PrimaryKeys 	[]string		//主键
	TableName		string			//表名字
}


func (self *DBBeanContent) getType() template.TemplateType {
	return template.DBBeanType
}


/**
 * 属性字段
 */
type BeanProp struct {

	TypeClazz string   //类型
	PropName string    //属性类型
}




