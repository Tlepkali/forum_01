package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"forum/internal/helpers/auth"
	"forum/internal/helpers/cookies"
	"forum/internal/models"
)

const (
	githubAuthURL     = "https://github.com/login/oauth/authorize"
	githubTokenURL    = "https://github.com/login/oauth/access_token"
	githubUserInfoURL = "https://api.github.com/user"
)

type githubUserInfo struct {
	Name   string `json:"name"`
	NodeID string `json:"node_id"`
}

func (h *Handler) handleGithubLogin(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&scope=user:email", githubAuthURL, h.githubConfig.ClientID, h.githubConfig.RedirectURL)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *Handler) handleGithubCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	// Exchange authorization code for access token
	data := strings.NewReader(fmt.Sprintf("code=%s&client_id=%s&client_secret=%s&redirect_uri=%s", code, h.githubConfig.ClientID, h.githubConfig.ClientSecret, h.githubConfig.RedirectURL))
	resp, err := http.Post(githubTokenURL, "application/x-www-form-urlencoded", data)
	if err != nil {
		log.Printf("Failed to exchange token: %v", err)
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	// Extract access token from response
	accessToken, err := auth.ExtractAccessTokenFromResponse(string(body))
	if err != nil {
		log.Printf("Failed to extract access token from response")
		http.Error(w, "Failed to extract access token from response", http.StatusInternalServerError)
		return
	}

	// Get user info using the access token
	userInfo, err := h.getUserInfo(accessToken, githubUserInfoURL)
	if err != nil {
		log.Printf("Failed to get user info: %v", err)
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}

	var githubUserInfo githubUserInfo

	err = json.Unmarshal(userInfo, &githubUserInfo)
	if err != nil {
		log.Printf("Failed to unmarshal user info: %v", err)
		http.Error(w, "Failed to unmarshal user info", http.StatusInternalServerError)
		return
	}

	// Use the user info to create or update user in your forum
	// ...
	user, _ := h.service.UserService.GetUserByEmail(githubUserInfo.NodeID)
	if user == nil {
		userDTO := &models.CreateUserDTO{
			Username: githubUserInfo.Name,
			Email:    githubUserInfo.NodeID,
			Password: githubUserInfo.NodeID,
		}
		err = h.service.UserService.CreateUser(userDTO)
		if err != nil {
			log.Printf("Failed to create user: %v", err)
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}
	}

	userLogin := &models.LoginUserDTO{
		Email:    githubUserInfo.NodeID,
		Password: githubUserInfo.NodeID,
	}

	userID, err := h.service.UserService.LoginUser(userLogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := h.service.SessionService.CreateSession(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookies.SetCookie(w, session.UUID, int(time.Until(session.ExpireAt).Seconds()))

	http.Redirect(w, r, "/", http.StatusFound)
}
