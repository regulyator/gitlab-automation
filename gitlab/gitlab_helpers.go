package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const updateMrUrl = "https://gitlab.com/api/v4/projects/%d/merge_requests/%d"
const getApprovalRulesForMrUrl = "https://gitlab.com/api/v4/projects/%d/merge_requests/%d/approval_rules"

func ParseMergeRequestAction(r *http.Request) (*GitlabMergeRequestAction, error) {
	b, err := io.ReadAll(r.Body)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	var mrAction GitlabMergeRequestAction
	err = json.Unmarshal(b, &mrAction)
	if err != nil {
		return nil, err
	}

	return &mrAction, nil
}

func setMrReviewers(projectID int, mrIID int, token string, reviewerIds []int) error {

	url := fmt.Sprintf(updateMrUrl,
		projectID,
		mrIID,
	)

	updateBody := []byte(fmt.Sprintf(`{"reviewer_ids": %v}`,
		strings.Join(strings.Fields(fmt.Sprint(reviewerIds)), ", ")))

	req, err := generateRequestWithTokenHeader(http.MethodPut, url, token, bytes.NewBuffer(updateBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(resp.StatusCode)
		return err
	}

	return nil
}

func getApproversIdsForMr(projectID int, mrIID int, token string) ([]int, error) {
	url := fmt.Sprintf(getApprovalRulesForMrUrl,
		projectID,
		mrIID,
	)

	req, err := generateRequestWithTokenHeader(http.MethodGet, url, token, nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var approvalRules []Rule
	err = json.Unmarshal(body, &approvalRules)
	if err != nil || approvalRules == nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, err
	}

	var approversIds []int
	for _, rule := range approvalRules {
		if rule.RuleType == "regular" {
			for _, user := range rule.EligibleApprovers {
				approversIds = append(approversIds, user.ID)
			}
			return approversIds, nil
		}
	}

	return approversIds, nil
}

func generateRequestWithTokenHeader(method string, url string, token string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("PRIVATE-TOKEN", token)
	return req, nil
}
