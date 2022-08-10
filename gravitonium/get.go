package gravitonium

import (
	"io"

	"github.com/go-zoox/fetch"
)

func (g *Gravitonium) Get(path string) (io.ReadCloser, error) {
	g.setup()

	fileptah := g.getFilePath(path)
	url := g.getAPIURL(APIs.File) + "/" + fileptah

	// fmt.Println("[gravitonium.Get] request:", url, "accessToken:", g.accessToken)

	response, err := fetch.Stream(url, &fetch.Config{
		Method: "GET",
		Headers: map[string]string{
			"Authorization": "Bearer " + g.accessToken,
		},
	})
	if err != nil {
		return nil, err
	}

	// fmt.PrintJSON("[gravitonium.Get] request2:", response.Request)

	return response.Stream, nil
}
