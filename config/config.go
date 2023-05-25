package config

import (
	"encoding/json"
	"os"
	"sync"
)

type Config struct {
	Addr string `json:"addr" env-default:":8080"`
	DB   struct {
		DSN string `json:"dsn"`
	} `json:"db"`
	StaticDir    string       `json:"static_dir" env-default:"./ui/static"`
	TemplatesDir string       `json:"template_dir" env-default:"./ui/templates"`
	GoogleConfig GoogleConfig `json:"google_config"`
	GithubConfig GithubConfig `json:"github_config"`
}

type GoogleConfig struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURL  string `json:"redirect_url"`
}

type GithubConfig struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURL  string `json:"redirect_url"`
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig(path string) *Config {
	once.Do(func() {
		instance = &Config{}
		if err := instance.configParser("config.json"); err != nil {
			panic(err)
		}
	})
	return instance
}

func (c *Config) configParser(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(c)
}
