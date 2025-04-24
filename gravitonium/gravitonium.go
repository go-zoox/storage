package gravitonium

import (
	"sync"
	"time"

	"github.com/go-zoox/fetch"
	"github.com/go-zoox/storage"
)

const HOST = "https://gcdn.zcorky.com"

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
	ClientID     string
	ClientSecret string
	Bucket       string
	//
	accessToken string
	//
	authLock sync.Mutex
	//
	cfg *Config
}

type Config struct {
	//
	Server string
}

func New(clientID string, clientSecret string, Bucket string, opts ...func(cfg *Config)) storage.Storage {
	cfg := &Config{
		Server: HOST,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	return &Gravitonium{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Bucket:       Bucket,
		//
		cfg: cfg,
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
			"appId":     g.ClientID,
			"appSecret": g.ClientSecret,
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
