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
	googleAuthURL     = "https://accounts.google.com/o/oauth2/auth"
	googleTokenURL    = "https://accounts.google.com/o/oauth2/token"
	googleUserInfoURL = "https://www.googleapis.com/oauth2/v3/userinfo"
)

type googleUserInfo struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Sub   string `json:"sub"`
}

func (h *Handler) handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=code&scope=profile email", googleAuthURL, h.googleConfig.ClientID, h.googleConfig.RedirectURL)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *Handler) handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	// Exchange authorization code for access token
	data := strings.NewReader(fmt.Sprintf("code=%s&client_id=%s&client_secret=%s&redirect_uri=%s&grant_type=authorization_code", code, h.googleConfig.ClientID, h.googleConfig.ClientSecret, h.googleConfig.RedirectURL))
	resp, err := http.Post(googleTokenURL, "application/x-www-form-urlencoded", data)
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
	accessToken := auth.ExtractValueFromBody(body, "access_token")
	if accessToken == "" {
		log.Printf("Failed to extract access token from response")
		http.Error(w, "Failed to extract access token", http.StatusInternalServerError)
		return
	}

	// Get user info using the access token
	userInfo, err := h.getUserInfo(accessToken, googleUserInfoURL)
	if err != nil {
		log.Printf("Failed to get user info: %v", err)
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}

	var googleUserInfo googleUserInfo

	err = json.Unmarshal(userInfo, &googleUserInfo)
	if err != nil {
		log.Printf("Failed to unmarshal user info: %v", err)
		http.Error(w, "Failed to unmarshal user info", http.StatusInternalServerError)
		return
	}

	user, _ := h.service.UserService.GetUserByEmail(googleUserInfo.Email)
	if user == nil {
		userDTO := &models.CreateUserDTO{
			Username: googleUserInfo.Name,
			Email:    googleUserInfo.Email,
			Password: googleUserInfo.Sub,
		}
		err = h.service.UserService.CreateUser(userDTO)
		if err != nil {
			log.Printf("Failed to create user: %v", err)
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}
	}

	userLogin := &models.LoginUserDTO{
		Email:    googleUserInfo.Email,
		Password: googleUserInfo.Sub,
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
