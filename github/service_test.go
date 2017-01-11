package github

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the GitHub service being tested.
	//service GithubService

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

func TestGithubService_FetchRepo(t *testing.T) {

	service := NewGithubService();
	service.

}

//func TestGitService_GetBlob(t *testing.T) {
//	setup()
//	defer teardown()
//
//	mux.HandleFunc("/repos/o/r/git/blobs/s", func(w http.ResponseWriter, r *http.Request) {
//		if m := "GET"; m != r.Method {
//			t.Errorf("Request method = %v, want %v", r.Method, m)
//		}
//		fmt.Fprint(w, `{
//			  "sha": "s",
//			  "content": "blob content"
//			}`)
//	})
//
//	blob, _, err := client.Git.GetBlob("o", "r", "s")
//	if err != nil {
//		t.Errorf("Git.GetBlob returned error: %v", err)
//	}
//
//	want := Blob{
//		SHA:     String("s"),
//		Content: String("blob content"),
//	}
//
//	if !reflect.DeepEqual(*blob, want) {
//		t.Errorf("Blob.Get returned %+v, want %+v", *blob, want)
//	}
//}
