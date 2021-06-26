package githubgo

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func GetTags(gitHubUser string, accessToken string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	// orgs, _, err := client.Organizations.List(context.Background(), gitHubUser, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%v", orgs)

	// list public repositories for org "github"
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "github", opt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", repos)
}
