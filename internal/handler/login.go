package handler

import (
	"net/http"
	"time"

	"forum/internal/helpers/cookies"
	"forum/internal/models"
	"forum/internal/render"
	"forum/pkg/forms"
)

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/user/login" {
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
	form.Required("email", "password")
	form.MaxLength("email", 50)
	form.MatchesPattern("email", forms.EmailRX)
	form.MaxLength("password", 50)
	form.MinLength("password", 8)

	if !form.Valid() {
		w.WriteHeader(http.StatusBadRequest)
		h.templateCache.Render(w, r, "login.page.html", &render.PageData{
			Form: form,
		})
		return
	}

	user := &models.LoginUserDTO{
		Email:    form.Get("email"),
		Password: form.Get("password"),
	}

	userID, err := h.service.UserService.LoginUser(user)
	if err != nil {
		if err == models.ErrInvalidCredentials {
			form.Errors.Add("generic", "Email or password is incorrect")
			h.templateCache.Render(w, r, "login.page.html", &render.PageData{
				Form: form,
			})
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := h.service.SessionService.CreateSession(userID)
	if err != nil {
		switch err {
		case models.ErrSessionAlreadyExists:
			form.Errors.Add("generic", "Session already exists")
			w.WriteHeader(http.StatusBadRequest)
			h.templateCache.Render(w, r, "login.page.html", &render.PageData{
				Form: form,
			})
			return
		}
	}

	cookies.SetCookie(w, session.UUID, int(time.Until(session.ExpireAt).Seconds()))

	http.Redirect(w, r, "/", http.StatusFound)
}
