package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const updateMrUrl = "https://gitlab.com/api/v4/projects/%d/merge_requests/%d"
const getApprovalRulesForMrUrl = "https://gitlab.com/api/v4/projects/%d/merge_requests/%d/approval_rules"
const getAllProjectMrUrl = "https://gitlab.com/api/v4/merge_requests?scope=all&state=opened&source_branch=%s"

func ParseMergeRequestAction(r *http.Request) (*GitlabMergeRequestAction, error) {
	log.SetOutput(os.Stdout)
	log.Println("try parse incoming body...")
	b, err := io.ReadAll(r.Body)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	var mrAction GitlabMergeRequestAction
	err = json.Unmarshal(b, &mrAction)
	if err != nil {
		log.Println("error when parse incoming body:(...")
		log.Println(err)
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

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func getApproversIdsForMr(projectID int, mrIID int, token string) ([]int, error) {
	log.SetOutput(os.Stdout)
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
		log.Println("Error reading response body:", err)
		return nil, err
	}

	var approvalRules []Rule
	err = json.Unmarshal(body, &approvalRules)
	if err != nil || approvalRules == nil {
		log.Println("Error parsing JSON:", err)
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

func getAllProjectMrBySourceName(sourceBranch string, token string) ([]GitlabMergeRequest, error) {
	log.SetOutput(os.Stdout)
	url := fmt.Sprintf(getAllProjectMrUrl,
		sourceBranch,
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
		log.Println("Error reading response body:", err)
		return nil, err
	}

	var mrs []GitlabMergeRequest
	err = json.Unmarshal(body, &mrs)
	if err != nil {
		log.Println("Error parsing JSON:", err)
		return nil, err
	}

	return mrs, nil
}

func isMrHasRelatedNotApproved(currentIId int, mrSourceBranch string, token string) bool {
	mrs, err := getAllProjectMrBySourceName(mrSourceBranch, token)
	if len(mrs) == 0 || err != nil {
		return false
	}

	for _, mr := range mrs {
		if mr.DetailedMergeStatus == "not_approved" && mr.IID != currentIId {
			return true
		}
	}

	return false
}

func generateRequestWithTokenHeader(method string, url string, token string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("PRIVATE-TOKEN", token)
	return req, nil
}
