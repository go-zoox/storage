package gravitonium

import (
	"testing"
)

func TestCheckAuth(t *testing.T) {
	accessToken := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhcHBJZCI6ImFiYyIsInVzZXJJZCI6ImFiYyIsImFwcE5hbWUiOiJ4eHgiLCJpYXQiOjE3MzA2MDkwNTQsImV4cCI6MTczMDYxNjI1NH0.Q7_CMFgnl4wrgjxfQVrSEgIdt2bk_YAaS3zQA4M0sYQjiGMSu5L3nToHuHdiJYDiG-AuKHLs7gwlIgbZ_EoVeA"

	if IsAccessTokenValid(accessToken) {
		t.Fatalf("IsAccessTokenValid(%s) got true, want false", accessToken)
	}
}
