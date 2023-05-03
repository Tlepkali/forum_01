package handler

import (
	"net/http"

	"forum/internal/models"
	"forum/internal/render"
	"forum/pkg/forms"
)

func (h *Handler) registerPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/user/registration" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		h.logger.PrintError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("name", "email", "password")
	form.MaxLength("username", 50)
	form.MaxLength("email", 50)
	form.MatchesPattern("email", forms.EmailRX)
	form.MaxLength("password", 50)
	form.MinLength("password", 8)

	if !form.Valid() {
		form.Errors.Add("generic", "Invalid credentials")
		w.WriteHeader(http.StatusBadRequest)
		h.templateCache.Render(w, r, "register.page.html", &render.PageData{
			Form: form,
		})
		return
	}

	user := &models.CreateUserDTO{
		Username: form.Get("name"),
		Email:    form.Get("email"),
		Password: form.Get("password"),
	}

	err = h.service.UserService.CreateUser(user)
	if err != nil {
		switch err {
		case models.ErrDuplicateEmail:
			form.Errors.Add("email", "Email already in use")
			w.WriteHeader(http.StatusBadRequest)
			h.templateCache.Render(w, r, "register.page.html", &render.PageData{
				Form: form,
			})
			return
		case models.ErrDuplicateUsername:
			form.Errors.Add("name", "Username already in use")
			w.WriteHeader(http.StatusBadRequest)
			h.templateCache.Render(w, r, "register.page.html", &render.PageData{
				Form: form,
			})
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
