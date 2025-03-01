package gravitonium

import (
	"strings"
	"time"

	"github.com/go-zoox/jwt"
)

func (g *Gravitonium) getAPIURL(path string) string {
	return HOST + path
}

func (g *Gravitonium) getFilePath(filepath string) string {
	if !strings.HasPrefix(filepath, "/") {
		panic("filepath must not start with /")
	}

	return filepath[1:]
}

func IsAccessTokenValid(accessToken string) bool {
	// no accessToken
	if accessToken == "" {
		return false
	}

	// @TODO check jwt expired
	_, payload, _, _, _, err := jwt.Parse(accessToken)
	if err == nil {
		expiredAt := payload.Get("exp").Int64()
		if expiredAt < time.Now().Unix() {
			return false
		}
	}

	return true
}
