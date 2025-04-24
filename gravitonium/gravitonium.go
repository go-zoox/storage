package gravitonium

import (
	"os"
	"sync"
	"time"

	"github.com/go-zoox/fetch"
	"github.com/go-zoox/storage"
)

const HOST = "https://gcdn.zcorky.com"

const ENV_GRAVITONIUM_CLIENT_ID_KEY = "GRAVITONIUM_CLIENT_ID"
const ENV_GRAVITONIUM_CLIENT_SECRET_KEY = "GRAVITONIUM_CLIENT_SECRET"
const ENV_GRAVITONIUM_BUCKET_KEY = "GRAVITONIUM_BUCKET"
const ENV_GRAVITONIUM_SERVER_KEY = "GRAVITONIUM_SERVER"

var APIs = struct {
	Token  string
	Upload string
	File   string
	// Create  string
	Remove  string
	Inspect string
	List    string
}{
	Token:  "/sdk/token",
	Upload: "/sdk/upload",
	File:   "/sdk/file",
	// Create:  "/sdk/create",
	Remove:  "/sdk/remove",
	Inspect: "/sdk/inspect",
	List:    "/sdk/list",
}

type Gravitonium struct {
	//
	accessToken string
	//
	authLock sync.Mutex
	//
	cfg *Config
}

type Config struct {
	ClientID     string
	ClientSecret string
	Bucket       string
	//
	Server string
}

func New(clientID string, clientSecret string, Bucket string, opts ...func(cfg *Config)) storage.Storage {
	cfg := &Config{
		Server: HOST,
		//
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Bucket:       Bucket,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	return &Gravitonium{
		//
		cfg: cfg,
	}
}

// WithServer ...
func WithServer(server string) func(cfg *Config) {
	return func(cfg *Config) {
		cfg.Server = server
	}
}

// WithEnv ...
func WithEnv() func(cfg *Config) {
	return func(cfg *Config) {
		if v := os.Getenv(ENV_GRAVITONIUM_SERVER_KEY); v != "" {
			cfg.Server = v
		}

		if v := os.Getenv(ENV_GRAVITONIUM_CLIENT_ID_KEY); v != "" {
			cfg.ClientID = v
		}

		if v := os.Getenv(ENV_GRAVITONIUM_CLIENT_SECRET_KEY); v != "" {
			cfg.ClientSecret = v
		}

		if v := os.Getenv(ENV_GRAVITONIUM_BUCKET_KEY); v != "" {
			cfg.Bucket = v
		}
	}
}

func (g *Gravitonium) isAccessTokenValid() bool {
	return IsAccessTokenValid(g.accessToken)
}

func (g *Gravitonium) checkAuth() error {
	g.authLock.Lock()
	defer g.authLock.Unlock()

	if g.isAccessTokenValid() {
		return nil
	}

	url := g.getAPIURL(APIs.Token)
	response, err := fetch.Post(url, &fetch.Config{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: map[string]string{
			"appId":     g.cfg.ClientID,
			"appSecret": g.cfg.ClientSecret,
		},
	})
	if err != nil {
		return err
	}

	if !response.Ok() {
		return response.Error()
	}

	// fmt.Println("[gravitonium.setup] response:", response.String())

	g.accessToken = response.Get("result.access_token").String()
	return nil
}

func (g *Gravitonium) reauthenticate() error {
	// fmt.Println("[gravitonium.reauthenticate] reauthenticate")
	time.Sleep(3 * time.Second)
	g.accessToken = ""
	return g.checkAuth()
}
