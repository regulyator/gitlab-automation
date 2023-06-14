package main

import (
	"gitlab-automation/gitlab"
	"net/http"
	"os"
)

var GitLabToken string
var GitLabHeaderSecret string
var YouTrackToken string
var YouTrackBasePath string

func main() {
	GitLabToken = os.Getenv("GITLAB_TOKEN")
	YouTrackToken = os.Getenv("YOU_TRACK_TOKEN")
	YouTrackBasePath = os.Getenv("YOU_TRACK_BASE_PATH")
	GitLabHeaderSecret = os.Getenv("GITLAB_SECRET_HEADER")
	if len(GitLabToken) == 0 || len(YouTrackToken) == 0 || len(YouTrackBasePath) == 0 {
		panic("some variables not has been set!")
	}
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/", handler)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if r.Header["X-Gitlab-Token"][0] == GitLabHeaderSecret {
		if result, err := gitlab.ParseMergeRequestAction(r); err == nil {
			go gitlab.ProcessMergeRequestAction(result, GitLabToken, YouTrackToken, YouTrackBasePath)
		}
	}
	_, _ = w.Write(nil)
	return
}
