package gravitonium

import (
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
}

func New(clientID string, clientSecret string, Bucket string) storage.Storage {
	return &Gravitonium{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Bucket:       Bucket,
	}
}

func (g *Gravitonium) setup() error {
	if g.accessToken != "" {
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
