package auth

import (
	"net/http"
	"testing"
)

func TestEmptyHeader(t *testing.T) {
	headers := http.Header{}
	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatalf("should fail on empty auth header")
	}
}
