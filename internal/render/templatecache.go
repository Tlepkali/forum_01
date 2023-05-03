package render

import (
	"html/template"
	"log"
	"path/filepath"
	"time"
)

type TemplateCache map[string]*template.Template

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Local().Format("2/1/2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func NewTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*page.html"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*layout.html"))
		if err != nil {
			log.Println(err)
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*partial.html"))
		if err != nil {
			log.Println(err)
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
