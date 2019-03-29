package dbcontent

import "template"

/**
 * dbbean 生成内容
 */
type DBBeanBaseContent struct {

	PackageName		string   		//包名
	ClazzName		string	 		//类名
	ClazzDescribe	string   		//类名描述
	BeanProps		[]BeanProp 		//属性字段
	//PrimaryKeys 	[]string		//主键
	TableName		string			//表名字
}

type DBBeanContent struct {
	PackageName		string   		//包名
	ClazzName		string	 		//类名
	ExtendsName		string			//继承者名字
	ClazzDescribe	string   		//类名描述
}

func (self *DBBeanBaseContent) getType() template.TemplateType {
	return template.DBBeanBaseType
}

func (self *DBBeanContent) getType() template.TemplateType {
	return template.DBBeanType
}


/**
 * 属性字段
 */
type BeanProp struct {

	TypeClazz string   //类型
	ColumnName string  //数据库表字段名
	PropName string    //属性类型
	ClazzName string   //类名称
	FlagPRI bool	   //是否为主键
}


/**
 * 字段信息
 */
type ColumnContent struct {
	Column_name string		//名字
	Column_type string		//类型
	Column_key string   	//是否为主键
	Column_comment string 	//注释
}




