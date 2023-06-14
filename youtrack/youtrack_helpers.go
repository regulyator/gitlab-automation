package youtrack

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const youTrackUrl = "https://%s/api/commands"

func applyQueryToIssue(issue string, query string, token string, basePath string) error {
	url := fmt.Sprintf(youTrackUrl,
		basePath,
	)

	updateBody := []byte(fmt.Sprintf(`
			{
			"query": "%s",
			"issues": [ { "idReadable": "%s" } ] 
			}`, query, issue))

	req, err := generateRequestWithTokenHeader(http.MethodPost, url, token, bytes.NewBuffer(updateBody))
	if err != nil {
		return err
	}

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(r.StatusCode)
		return err
	}

	return nil
}

func generateRequestWithTokenHeader(method string, url string, token string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
