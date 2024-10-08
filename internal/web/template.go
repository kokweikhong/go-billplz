package web

import (
	"embed"
	"html/template"
	"net/http"
)

type WebTemplate struct {
	Template map[string]*template.Template
}

//go:embed views/*.html
var views embed.FS

func (a *app) InitTemplate() {
	a.template = new(WebTemplate)
	tpl := make(map[string]*template.Template)

	tpl["index"] = template.Must(template.ParseFS(views, "views/index.html"))

	a.template.Template = tpl
}

func (a *app) RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl, ok := a.template.Template[name]
	if !ok {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
