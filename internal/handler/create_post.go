package handler

import (
	"fmt"
	"net/http"

	"forum/internal/models"
	"forum/internal/render"
	"forum/pkg/forms"
)

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post/create" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			h.logger.PrintError(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		form := forms.New(r.PostForm)
		form.Required("title", "content")
		form.MaxLength("title", 100)
		form.MaxLength("content", 10000)

		if !form.Valid() {

			categories, err := h.service.CategoryService.GetAllCategories()
			if err != nil {
				h.logger.PrintError(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			form.Errors.Add("generic", "Form is not valid")
			form.Categories = append(form.Categories, categories...)

			w.WriteHeader(http.StatusBadRequest)
			h.templateCache.Render(w, r, "create.page.html", &render.PageData{
				Form:              form,
				AuthenticatedUser: h.getUserFromContext(r),
			})
			return
		}

		categoriesS := r.PostForm["categories"]

		if len(categoriesS) == 0 {

			categories, err := h.service.CategoryService.GetAllCategories()
			if err != nil {
				h.logger.PrintError(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			form.Errors.Add("generic", "You must select at least one category")
			form.Categories = append(form.Categories, categories...)

			w.WriteHeader(http.StatusBadRequest)
			h.templateCache.Render(w, r, "create.page.html", &render.PageData{
				Form:              form,
				AuthenticatedUser: h.getUserFromContext(r),
			})
			return
		}

		categories := make([]*models.Category, 0, len(categoriesS))

		for _, category := range categoriesS {
			c, err := h.service.CategoryService.GetCategoryByName(category)
			if err != nil {
				switch err {
				case models.ErrSqlNoRows:
					form.Errors.Add("generic", fmt.Sprintf("Category %s does not exist", category))
					w.WriteHeader(http.StatusBadRequest)
					h.templateCache.Render(w, r, "create.page.html", &render.PageData{
						Form:              form,
						AuthenticatedUser: h.getUserFromContext(r),
					})
				default:
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
			categories = append(categories, c)
		}

		author := h.getUserFromContext(r)

		post := &models.CreatePostDTO{
			Title:      form.Get("title"),
			Content:    form.Get("content"),
			Author:     author.ID,
			AuthorName: author.Username,
			Categories: categories,
		}

		id, err := h.service.PostService.CreatePost(post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/post/%d", id), http.StatusSeeOther)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	categories, err := h.service.CategoryService.GetAllCategories()
	if err != nil {
		h.logger.PrintError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	form := forms.New(nil)

	form.Categories = append(form.Categories, categories...)

	h.templateCache.Render(w, r, "create.page.html", &render.PageData{
		Form:              form,
		AuthenticatedUser: h.getUserFromContext(r),
	})
}
