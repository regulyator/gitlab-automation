package main

import (
	"gitlab-automation/gitlab"
	"log"
	"net/http"
	"os"
)

var GitLabToken string
var GitLabHeaderSecret string
var YouTrackToken string
var YouTrackBasePath string

func main() {
	log.SetOutput(os.Stdout)
	GitLabToken = os.Getenv("GITLAB_TOKEN")
	YouTrackToken = os.Getenv("YOU_TRACK_TOKEN")
	YouTrackBasePath = os.Getenv("YOU_TRACK_BASE_PATH")
	GitLabHeaderSecret = os.Getenv("GITLAB_SECRET_HEADER")
	if len(GitLabToken) == 0 || len(YouTrackToken) == 0 || len(YouTrackBasePath) == 0 {
		panic("some variables not has been set!")
	}
	log.Println("...start listening for hooks:)")
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.SetOutput(os.Stdout)
	w.WriteHeader(http.StatusOK)
	if len(r.Header["X-Gitlab-Token"]) == 0 || r.Header["X-Gitlab-Token"][0] == GitLabHeaderSecret {
		log.Println("process incoming request...")
		if result, err := gitlab.ParseMergeRequestAction(r); err == nil {
			go gitlab.ProcessMergeRequestAction(result, GitLabToken, YouTrackToken, YouTrackBasePath)
		}
	}
	_, _ = w.Write(nil)
	return
}
