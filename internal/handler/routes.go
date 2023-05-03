package handler

import (
	"net/http"

	"forum/config"
)

func (h *Handler) Routes(cfg *config.Config) http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", h.index)

	mux.HandleFunc("/user/signup", h.register)
	mux.HandleFunc("/user/registration", h.registerPost)
	mux.HandleFunc("/login", h.loginForm)
	mux.HandleFunc("/user/login", h.login)
	mux.Handle("/user/logout", h.requireAuthentication(http.HandlerFunc(h.logout)))

	mux.Handle("/post/create", h.requireAuthentication(http.HandlerFunc(h.createPost)))
	mux.HandleFunc("/post/", h.showPost)

	mux.Handle("/comment/create", h.requireAuthentication(http.HandlerFunc(h.createComment)))

	mux.Handle("/post/vote/create", h.requireAuthentication(http.HandlerFunc(h.createPostVote)))
	mux.Handle("/comment/vote/create", h.requireAuthentication(http.HandlerFunc(h.createCommentVote)))

	mux.Handle("/myposts", h.requireAuthentication(http.HandlerFunc(h.showMyPosts)))
	mux.Handle("/likedposts", h.requireAuthentication(http.HandlerFunc(h.showLikedPosts)))
	mux.Handle("/showposts", http.HandlerFunc(h.showPostsByCategory))

	return h.authenticate(h.recoverPanic(h.logRequest(h.secureHeaders(mux))))
}
