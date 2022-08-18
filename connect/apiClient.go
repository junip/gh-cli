package connect

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"

	"github.com/joho/godotenv"
)

func ApiClient() *github.Client {
	context := context.Background()
	envLoadErr := godotenv.Load()

	if envLoadErr != nil {
		fmt.Print("please create .env file in your root folder")
	}

	apiToken := os.Getenv("API_TOKEN")

	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: apiToken},
	)

	tokenClient := oauth2.NewClient(context, tokenService)
	client := github.NewClient(tokenClient)
	return client
}
