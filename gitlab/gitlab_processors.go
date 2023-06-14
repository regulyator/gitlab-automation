package gitlab

import (
	"fmt"
	"gitlab-automation/youtrack"
	"regexp"
)

func ProcessMergeRequestAction(mrAction *GitlabMergeRequestAction, gitlabToken string, youTrackToken string, youTrackBasePath string) {
	if isMrAcceptable(mrAction) {
		switch mrActionStatus := mrAction.ObjectAttributes.Action; mrActionStatus {
		case "open", "reopen", "update":
			if len(mrAction.ObjectAttributes.ReviewerIDs) == 0 {
				err := setReviewers(mrAction, gitlabToken)
				if err == nil {
					youtrack.UpdateIssueStatus(mrActionStatus, mrAction.ObjectAttributes.SourceBranch, youTrackToken, youTrackBasePath)
				}
			}
		case "approved":
			if !isMrHasRelatedNotApproved(mrAction.ObjectAttributes.IID, mrAction.ObjectAttributes.SourceBranch, gitlabToken) {
				youtrack.UpdateIssueStatus(mrActionStatus, mrAction.ObjectAttributes.SourceBranch, youTrackToken, youTrackBasePath)
			}
		case "unapproved":
			youtrack.UpdateIssueStatus(mrActionStatus, mrAction.ObjectAttributes.SourceBranch, youTrackToken, youTrackBasePath)
		}

	}
}

func setReviewers(mrAction *GitlabMergeRequestAction, gitlabToken string) error {
	reviewerIds, err := getApproversIdsForMr(mrAction.Project.ID, mrAction.ObjectAttributes.IID, gitlabToken)
	if err != nil {
		return err
	}
	if len(reviewerIds) == 0 {
		return fmt.Errorf("DEFAULT APPROVERS FOR %d NOT FOUND", mrAction.Project.ID)
	}

	err = setMrReviewers(mrAction.Project.ID, mrAction.ObjectAttributes.IID, gitlabToken, reviewerIds)
	if err != nil {
		return err
	}

	return nil

}

func isMrAcceptable(mrAction *GitlabMergeRequestAction) bool {
	mrSourceBranchRegExp := regexp.MustCompile(`\bREV-\w+`)

	return mrAction.EventType == "merge_request" &&
		isMrActionIsAcceptable(mrAction.ObjectAttributes.Action) &&
		mrSourceBranchRegExp.MatchString(mrAction.ObjectAttributes.SourceBranch)
}

func isMrActionIsAcceptable(action string) bool {
	arr := []string{"open", "reopen", "update", "approved", "unapproved"}

	for _, str := range arr {
		if str == action {
			return true
		}
	}

	return false
}
