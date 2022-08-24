package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	// "log"
	"net/http"
	"path/filepath"

	model "github.com/Precious025/go-web_app/models"
	"github.com/Precious025/go-web_app/pkg/config"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}
func AddDefaultData(td *model.TemplateData) *model.TemplateData {
	return td
}

func RenderTemplates(w http.ResponseWriter, tmpl string, td *model.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplate()
	}

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Template not found: ", tmpl)
	}

	buff := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buff, td)

	_, err := buff.WriteTo(w)
	if err != nil {
		fmt.Println("Failed to write to the browser", err)
	}
}

func CreateTemplate() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err := ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
			myCache[name] = ts
		}
	}
	return myCache, nil
}
