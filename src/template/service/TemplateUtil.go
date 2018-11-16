package service

import (
	"text/template"
	"os"
	"log"
	"template/dbcontent"
	"io"
)

type TemplateUtil struct {

}

type Inventory struct {
	Material string
	Count    int
}

func TestGen() {

	sweaters := Inventory{"axe", 0}
	html := `
	"test").Parse("{{.Count}} items are made of {{.Material}}"
	{{$a := .Count}}
	{{$b := 17}}
	{{$c := 18}}	
  
	{{if eq  .Count $b}}
	oo
	{{else}}
	xx
	{{end}}
	`
	tmpl, data := template.New("test").Parse(html)
	if data != nil {
		panic(data)
	}
	data = tmpl.Execute(os.Stdout, sweaters)
	if data != nil {
		panic(data)
	}
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