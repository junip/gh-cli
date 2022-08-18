/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"

	"github-cli/connect"

	"github.com/kyokomi/emoji/v2"
	"github.com/spf13/cobra"
)

type Repo struct {
	ID          int64
	NodeID      string
	Name        string
	FullName    string
	Description string
}

var RepoName string

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fetchRepoDetails(RepoName)

	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
	repoCmd.PersistentFlags().StringVarP(&RepoName, "name", "n", "", "Name of the repository")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func validateArgs(args []string) {

}

func fetchRepoDetails(repoName string) {
	context := context.Background()
	repo, _, err := connect.ApiClient().Repositories.Get(context, "junip", repoName)

	if err != nil {
		fmt.Printf("Could not fetch the repo %v", err)
	}

	repoDetail := Repo{
		ID:          *repo.ID,
		Name:        *repo.Name,
		FullName:    *repo.FullName,
		Description: *repo.Description,
	}

	emoji.Println(repoDetail.Name)
	emoji.Println(repoDetail.Description)

}
