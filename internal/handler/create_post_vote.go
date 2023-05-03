package handler

import (
	"fmt"
	"net/http"

	"forum/internal/models"
	"forum/pkg/forms"
)

func (h *Handler) createPostVote(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post/vote/create" {
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

	form.Required("post_id", "status")
	id := form.IsInt("post_id")
	status := form.IsStatus("status")

	author := h.getUserFromContext(r)

	vote := &models.PostVote{
		PostID:   id,
		Status:   status,
		AuthorID: author.ID,
	}

	if err := h.service.PostVoteService.CreateVote(vote); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/post/%d", id), http.StatusSeeOther)
}
