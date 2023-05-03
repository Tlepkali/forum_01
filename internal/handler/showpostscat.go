package handler

import (
	"net/http"

	"forum/internal/render"
)

func (h *Handler) showPostsByCategory(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/showposts" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	category := r.URL.Query().Get("category")

	categories, err := h.service.CategoryService.GetAllCategories()
	if err != nil {
		h.logger.PrintError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	posts, err := h.service.PostService.GetPostsByCategory(category)
	if err != nil {
		h.logger.PrintError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.templateCache.Render(w, r, "home.page.html", &render.PageData{
		Posts:             posts,
		Categories:        categories,
		AuthenticatedUser: h.getUserFromContext(r),
	})
}
