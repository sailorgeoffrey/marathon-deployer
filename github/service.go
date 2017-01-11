package github

import "log"
import "github.com/google/go-github/github"
import "golang.org/x/oauth2"

type GithubService struct {
	RepoUri               string
	RepoUsername          string
	RepoSecret            string
	LocalCheckoutLocation string
	client                github.Client
}

func NewGithubService(token string) *GithubService {

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	s := new(GithubService)
	s.client = github.NewClient(tc)

	return s
}

// Fetch the content from a GitHub URI
func (service *GithubService) FetchContent(serviceUris *[]string) *[]string {
	content := [len(*serviceUris)]string{}
	for i, item := range *serviceUris {
		log.Printf("\t%s", item)
		content[i] = nil
	}
	return &content
}
