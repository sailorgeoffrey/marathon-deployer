package github

import (
	"encoding/json"
	"log"
	"net/http"
)

const key = "1Nw2A7Zx0TH3f76PCv15wl4UuRm275hr"

type user struct {
	Name     string
	Email    string
	Username string
}

type commit struct {
	Author    user
	Committer user
	Message   string
	Added     []string
	Removed   []string
	Modified  []string
}

type config struct {
	Secret string
}

type hook struct {
	Type   string
	Id     int
	Name   string
	Events []string
	Config config
}

type repository struct {
	Name     string
	FullName string
}

type sender struct {
	Login string
}

type push struct {
	Commits    []commit
	Hook       hook
	Repository repository
	Sender     sender
}

type WebHookHandler struct {
}

func (WebHookHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var p push
	err := decoder.Decode(&p)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	// Make sure the key matches
	//if p.Hook.Config.Secret != key {
	//	w.WriteHeader(403)
	//	return
	//}

	log.Printf("push from %s", p.Sender.Login)

	servicesToLaunch := map[string]bool{}
	servicesToShutdown := map[string]bool{}
	l := 0

	for _, commitInstance := range p.Commits {
		log.Printf(commitInstance.Message)
		log.Print("Added:")
		for _, item := range commitInstance.Added {
			log.Printf("\t%s", item)
			servicesToLaunch[item] = true
			l++
		}
		log.Print("Removed:")
		for _, item := range commitInstance.Removed {
			log.Printf("\t%s", item)
			servicesToShutdown[item] = true

		}
		log.Print("Modified:")
		for _, item := range commitInstance.Modified {
			log.Printf("\t%s", item)
			servicesToLaunch[item] = true
			l++
		}
	}

}
