package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"forum/config"
	"forum/internal/models"
	"forum/internal/render"
	"forum/internal/service"
	"forum/pkg/logger"
)

type Handler struct {
	service       *service.Service
	templateCache render.TemplateCache
	logger        *logger.Logger
	googleConfig  config.GoogleConfig
	githubConfig  config.GithubConfig
}

func NewHandler(service *service.Service, tempCache render.TemplateCache, googc config.GoogleConfig, gitc config.GithubConfig) *Handler {
	return &Handler{
		service:       service,
		templateCache: tempCache,
		logger:        logger.GetLoggerInstance(),
		googleConfig:  googc,
		githubConfig:  gitc,
	}
}

func (h *Handler) getUserFromContext(r *http.Request) *models.User {
	user, ok := r.Context().Value(contextKeyUser).(*models.User)
	if !ok {
		h.logger.PrintInfo("User is not authenticated")
		return nil
	}
	return user
}

func (h *Handler) getUserInfo(accessToken string, userInfoURL string) ([]byte, error) {
	req, err := http.NewRequest("GET", userInfoURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
