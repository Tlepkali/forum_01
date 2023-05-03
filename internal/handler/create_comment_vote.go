package handler

import (
	"fmt"
	"net/http"

	"forum/internal/models"
	"forum/pkg/forms"
)

func (h *Handler) createCommentVote(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/comment/vote/create" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	form := forms.New(r.PostForm)

	form.Required("comment_id", "status", "post_id")
	postID := form.IsInt("post_id")
	id := form.IsInt("comment_id")
	status := form.IsStatus("status")

	if !form.Valid() {
		http.Redirect(w, r, fmt.Sprintf("/post/%d", postID), http.StatusSeeOther)
		return
	}

	author := h.getUserFromContext(r)

	vote := &models.CommentVote{
		CommentID: id,
		Status:    status,
		AuthorID:  author.ID,
	}

	if err := h.service.CommentVoteService.CreateVote(vote); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/post/%d", postID), http.StatusSeeOther)
}
