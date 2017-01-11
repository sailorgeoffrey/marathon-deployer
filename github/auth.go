package github

import (
	"encoding/json"
	"log"
	"net/http"
)

type AuthHandler struct {
}

func (AuthHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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

}
