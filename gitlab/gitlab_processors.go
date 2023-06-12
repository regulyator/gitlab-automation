package gitlab

import (
	"log"
)

func ProcessMergeRequestAction(mrAction *GitlabMergeRequestAction, token string) {
	if isMrAcceptable(mrAction) {
		log.Println(mrAction.EventType)
		log.Println(mrAction.ObjectAttributes.Action)

		reviewerIds, err := getApproversIdsForMr(mrAction.Project.ID, mrAction.ObjectAttributes.IID, token)
		if err == nil && len(reviewerIds) > 0 {
			_ = setMrReviewers(mrAction.Project.ID, mrAction.ObjectAttributes.IID, token, reviewerIds)
		}

	}

}

func isMrAcceptable(mrAction *GitlabMergeRequestAction) bool {
	return mrAction.EventType == "merge_request" &&
		isMrActionIsAcceptable(mrAction.ObjectAttributes.Action) &&
		len(mrAction.ObjectAttributes.ReviewerIDs) == 0
}

func isMrActionIsAcceptable(action string) bool {
	arr := []string{"open", "reopen"}

	for _, str := range arr {
		if str == action {
			return true
		}
	}

	return false
}