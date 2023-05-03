package handler

import (
	"net/http"

	"forum/internal/render"
)

func (h *Handler) showLikedPosts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/likedposts" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user := h.getUserFromContext(r)

	posts, err := h.service.PostService.GetLikedPosts(user.ID)
	if err != nil {
		h.logger.PrintError(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	h.templateCache.Render(w, r, "home.page.html", &render.PageData{
		Posts:             posts,
		AuthenticatedUser: user,
	})
}
