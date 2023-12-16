package admin

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

type DummyTestTokenSource struct {
	oauth2.TokenSource
}

func (d DummyTestTokenSource) Token() (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken: "abc",
	}, nil
}

func TestNewTokenSource(t *testing.T) {
	tokenSource := DummyTestTokenSource{}
	nebulaTokenSource := NewCustomHeaderTokenSource(tokenSource, true, "test")
	metadata, err := nebulaTokenSource.GetRequestMetadata(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, "Bearer abc", metadata["test"])
}
