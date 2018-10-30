package github

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// RepositoryRequest - Holds New Repository request
type RepositoryRequest struct {
	Name              string
	GithubAccessToken string
	AutoInit          bool
	GithubClient      *github.Client
	GithubRepo        *github.Repository
}

// CreateRepository creates a repository on github
func (rr *RepositoryRequest) CreateRepository() (err error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: rr.GithubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	rr.GithubClient = github.NewClient(tc)

	newrepo := github.Repository{}
	newrepo.Name = &rr.Name
	newrepo.AutoInit = &rr.AutoInit
	log.Printf("Creating a repository %v\n", newrepo)
	_, response, err := rr.GithubClient.Repositories.Create(ctx, "", &newrepo)
	if err != nil {
		//return err
	}
	fmt.Printf("Result: %v\n", response.StatusCode)

	repo, _, err := rr.GithubClient.Repositories.Get(ctx, "ubsbot", rr.Name)

	rr.GithubRepo = repo

	commit := createCommit()

	_, _, err = rr.GithubClient.Git.CreateCommit(ctx, rr.GithubRepo.Owner.GetLogin(), rr.Name, &commit)

	if err != nil {
		log.Printf("Failed to make new commit: %v\n", err)
		return err
	}

	return nil
}

// InitialCommit Makes the initial commit
func createCommit() (c github.Commit) {
	log.Printf("Creating Initial Commit")
	Message := "Initial Commit"
	Content := "YQo="
	te := github.TreeEntry{
		Content: &Content,
	}
	tes := []github.TreeEntry{te}
	commit := github.Commit{}
	commit.Message = &Message
	gt := github.Tree{}
	gt.Entries = tes
	commit.Tree = &gt

	return commit
}
