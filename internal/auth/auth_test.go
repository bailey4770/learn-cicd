package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("no auth header", func(t *testing.T) {
		got, err := GetAPIKey(http.Header{})
		assert.ErrorIs(t, err, ErrNoAuthHeaderIncluded)
		assert.Equal(t, "", got)
	})

	t.Run("malformed header", func(t *testing.T) {
		header := http.Header{}
		header.Set("Authorization", "Bearer sometoken") // wrong scheme
		got, err := GetAPIKey(header)
		assert.EqualError(t, err, "malformed authorization header")
		assert.Equal(t, "", got)
	})

	t.Run("valid header", func(t *testing.T) {
		header := http.Header{}
		header.Set("Authorization", "ApiKey 12345abcde")
		got, err := GetAPIKey(header)
		assert.NoError(t, err)
		assert.Equal(t, "12345abcde", got)
	})

	t.Run("broken test", func(t *testing.T) {
		header := http.Header{}
		header.Set("Authorization", "ApiKey 12345abcde")
		got, err := GetAPIKey(header)
		assert.NoError(t, err)
		assert.Equal(t, "", got)
	})
}
