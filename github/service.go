package github

import "github.com/google/go-github/github"

type GithubService struct {
	RepoUri               string
	RepoUsername          string
	RepoSecret            string
	LocalCheckoutLocation string
	client                github.Client
}

func NewGithubService(token string) *GithubService {

	s := new(GithubService)

	return s
}
