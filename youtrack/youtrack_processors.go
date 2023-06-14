package youtrack

import (
	"log"
	"os"
)

func UpdateIssueStatus(mrAction string, issue string, youTrackToken string, youTrackBasePath string) {
	log.SetOutput(os.Stdout)
	var statusUpdate string
	switch mrAction {
	case "open", "reopen":
		statusUpdate = CodeReview
	case "approved":
		statusUpdate = CodeReviewApproved
	case "unapproved":
		statusUpdate = DevInProgress
	default:
		log.Printf("Wrong mr action %s", mrAction)
	}

	if len(statusUpdate) > 0 {
		if err := applyQueryToIssue(issue, statusUpdate, youTrackToken, youTrackBasePath); err != nil {
			log.Printf("Error when update issue %s status to %s", issue, statusUpdate)
		}
	}

}

const (
	CodeReview         string = "code review"
	CodeReviewApproved        = "code review approved"
	DevInProgress             = "dev: in progress"
)
