package connect

import (
	"context"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func ApiClient() *github.Client {
	value := os.Getenv("token")
	context := context.Background()

	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: value},
	)

	tokenClient := oauth2.NewClient(context, tokenService)
	client := github.NewClient(tokenClient)
	return client
}
