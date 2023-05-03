package handler

import (
	"net/http"
	"strconv"
	"strings"

	"forum/internal/render"
)

func (h *Handler) showPost(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/post/") {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	pathID := r.URL.Path[len("/post/"):]
	id, err := strconv.Atoi(pathID)
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	post, err := h.service.PostService.GetPostByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	comments, err := h.service.CommentService.GetAllByPostID(post.ID)
	if err != nil {
		h.logger.PrintError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, comment := range comments {
		comment.Likes, comment.Dislikes, err = h.service.CommentVoteService.GetLikesAndDislikes(comment.ID)
		if err != nil {
			h.logger.PrintError(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	likes, dislikes, err := h.service.PostVoteService.GetLikesAndDislikes(post.ID)
	if err != nil {
		h.logger.PrintError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	post.Likes = likes
	post.Dislikes = dislikes

	h.templateCache.Render(w, r, "post.page.html", &render.PageData{
		Post:              post,
		Comments:          comments,
		AuthenticatedUser: h.getUserFromContext(r),
	})
}
