package connect

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func ApiClient() *github.Client {
	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)

	tokenClient := oauth2.NewClient(context, tokenService)
	client := github.NewClient(tokenClient)
	return client
}
