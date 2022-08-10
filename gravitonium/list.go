package gravitonium

import (
	"github.com/go-zoox/fetch"
	"github.com/go-zoox/storage"
)

func (g *Gravitonium) List(path string, page, pageSize int) []storage.StorageEntity {
	g.setup()

	fileptah := g.getFilePath(path)
	url := g.getAPIURL(APIs.List) + "/" + fileptah

	response, err := fetch.Stream(url, &fetch.Config{
		Method: "GET",
		Headers: map[string]string{
			"Authorization": "Bearer " + g.accessToken,
		},
	})
	if err != nil {
		return nil
	}

	data := []storage.StorageEntity{}

	// total := response.Get("total").Int()
	dataX := response.Get("data").Array()
	for _, one := range dataX {
		name := one.Get("name").String()
		path := one.Get("path").String()
		dir := one.Get("dir").String()
		size := one.Get("size").Int()
		typ := one.Get("type").String()
		hash := one.Get("hash").String()
		url := one.Get("url").String()

		data = append(data, storage.StorageEntity{
			Name:  name,
			Path:  path,
			Dir:   dir,
			Size:  size,
			IsDir: typ == "directory",
			Hash:  hash,
			URL:   url,
		})
	}

	return data
}
