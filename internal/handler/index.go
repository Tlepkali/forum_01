package handler

import (
	"net/http"
	"strconv"

	"forum/internal/render"
)

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}

	posts, err := h.service.PostService.GetAllPosts(offset, limit)
	if err != nil {
		h.logger.PrintError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := h.service.CategoryService.GetAllCategories()
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
