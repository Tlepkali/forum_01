package render

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

func (cache *TemplateCache) Render(w http.ResponseWriter, r *http.Request, name string, data *PageData) {
	ts, ok := map[string]*template.Template(*cache)[name]
	if !ok {
		fmt.Println("Template not found")
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, data)
	if err != nil {
		fmt.Println("Error executing template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}
