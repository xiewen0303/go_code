package dbcontent

import "template"

type IContentData interface {
	getType() template.TemplateType
}