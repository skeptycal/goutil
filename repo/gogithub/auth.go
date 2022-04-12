package gogithub

import (
	"context"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

var client githubv4.Client

func init() {
	client = authenticate()
}

func authenticate() githubv4.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	return *githubv4.NewClient(httpClient)
}
