package router

import (
	"context"
	"net/http"
	"strings"
)

type Router struct {
	mp map[string]map[string]func(http.ResponseWriter, *http.Request)
}

func NewRouter() *Router {
	return &Router{mp: make(map[string]map[string]func(http.ResponseWriter, *http.Request))}
}

func (r *Router) Handle(method string, pattern string, handler http.Handler) {
	r.HandleFunc(method, pattern, handler.ServeHTTP)
}

func (r *Router) HandleFunc(method, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	if r.mp[method] == nil {
		r.mp[method] = make(map[string]func(http.ResponseWriter, *http.Request))
	}
	r.mp[method][pattern] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Get the request method and URL path
	method := req.Method
	path := req.URL.Path

	// Check if requested path matches a file path
	if strings.HasPrefix(path, "/static/") {
		filePath := strings.TrimPrefix(path, "/static")
		http.ServeFile(w, req, filePath)
		return
	}

	// Find the handler function for the requested path and method
	handlers := r.mp[method]
	if handlers == nil {
		http.NotFound(w, req)
		return
	}
	var handler func(http.ResponseWriter, *http.Request)
	for pattern, h := range handlers {
		if match, params := matchPath(pattern, path); match {
			handler = h
			req = setPathParams(req, params)
			break
		}
	}
	if handler == nil {
		http.NotFound(w, req)
		return
	}

	// Call the handler function
	handler(w, req)
}

// matchPath checks if the given path matches the given pattern.
// If it does, it returns true and a map of path parameters.
func matchPath(pattern, path string) (bool, map[string]string) {
	patternParts := strings.Split(pattern, "/")
	pathParts := strings.Split(path, "/")
	if len(patternParts) != len(pathParts) {
		return false, nil
	}
	params := make(map[string]string)
	for i := 0; i < len(patternParts); i++ {
		if patternParts[i] == "" && pathParts[i] == "" {
			continue
		}
		if strings.HasPrefix(patternParts[i], ":") {
			paramName := strings.TrimPrefix(patternParts[i], ":")
			params[paramName] = pathParts[i]
		} else if patternParts[i] != pathParts[i] {
			return false, nil
		}
	}
	return true, params
}

// setPathParams adds the path parameters to the request context.
func setPathParams(req *http.Request, params map[string]string) *http.Request {
	return req.WithContext(context.WithValue(req.Context(), "params", params))
}

// GetPathParams retrieves the path parameters from the request context.
func GetPathParams(req *http.Request) map[string]string {
	params, _ := req.Context().Value("params").(map[string]string)
	return params
}
