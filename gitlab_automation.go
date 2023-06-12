package main

import (
	"gitlab-automation/gitlab"
	"net/http"
	"os"
)

var GitLabToken string

func main() {
	GitLabToken = os.Getenv("GITLAB_TOKEN")
	if len(GitLabToken) == 0 {
		panic("GITLAB_TOKEN not has been set!")
	}
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if result, err := gitlab.ParseMergeRequestAction(r); err == nil {
		go gitlab.ProcessMergeRequestAction(result, GitLabToken)
	}
	_, _ = w.Write(nil)
	return
}
