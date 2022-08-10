package gravitonium

import "strings"

func (g *Gravitonium) getAPIURL(path string) string {
	return HOST + path
}

func (g *Gravitonium) getFilePath(filepath string) string {
	if !strings.HasPrefix(filepath, "/") {
		panic("filepath must not start with /")
	}

	return filepath[1:]
}
