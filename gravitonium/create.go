package gravitonium

import (
	"io"

	"github.com/go-zoox/fetch"
)

func (g *Gravitonium) Create(path string, stream io.ReadCloser) error {
	// // @TODO force reauthenticate to avoid stream has been read
	// //
	// g.accessToken = ""

	g.checkAuth()

	if stream == nil {
		panic("[gravitonium.create] stream is nil")
	}

	// fileptah := g.getFilePath(path)
	url := g.getAPIURL(APIs.Upload) + "?filepath=" + path

	// fmt.Println("[gravitonium.Create] request:", url, "accessToken:", g.accessToken)

	response, err := fetch.Post(url, &fetch.Config{
		Headers: map[string]string{
			"Authorization": "Bearer " + g.accessToken,
			"Content-Type":  "multipart/form-data",
		},
		Body: map[string]any{
			"file": stream,
		},
	})
	if err != nil {
		return err
	}

	if !response.Ok() {
		if response.Status == 401 || response.Status == 403 {
			g.reauthenticate()

			return g.Create(path, stream)
		}

		return response.Error()
	}

	return nil
}
