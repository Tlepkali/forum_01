package handler

import (
	"net/http"

	"forum/internal/render"
	"forum/pkg/forms"
)

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/user/signup" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	h.templateCache.Render(w, r, "register.page.html", &render.PageData{
		Form:              forms.New(nil),
		AuthenticatedUser: h.getUserFromContext(r),
	})
}
