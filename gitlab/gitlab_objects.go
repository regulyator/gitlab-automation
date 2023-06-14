package gitlab

type GitlabMergeRequest struct {
	ID                  int    `json:"id"`
	IID                 int    `json:"iid"`
	Title               string `json:"title"`
	Description         string `json:"description"`
	State               string `json:"state"`
	MergeStatus         string `json:"merge_status"`
	DetailedMergeStatus string `json:"detailed_merge_status"`
}

type GitlabMergeRequestAction struct {
	ObjectKind       string     `json:"object_kind"`
	EventType        string     `json:"event_type"`
	User             User       `json:"user"`
	Project          Project    `json:"project"`
	ObjectAttributes Attribute  `json:"object_attributes"`
	Labels           []string   `json:"labels"`
	Changes          Changes    `json:"changes"`
	Repository       Repository `json:"repository"`
	Assignees        []User     `json:"assignees"`
}

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
}

type Project struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	WebURL            string `json:"web_url"`
	AvatarURL         string `json:"avatar_url"`
	GitSSHURL         string `json:"git_ssh_url"`
	GitHTTPURL        string `json:"git_http_url"`
	Namespace         string `json:"namespace"`
	VisibilityLevel   int    `json:"visibility_level"`
	PathWithNamespace string `json:"path_with_namespace"`
	DefaultBranch     string `json:"default_branch"`
	CIConfigPath      string `json:"ci_config_path"`
	Homepage          string `json:"homepage"`
	URL               string `json:"url"`
	SSHURL            string `json:"ssh_url"`
	HTTPURL           string `json:"http_url"`
}

type Attribute struct {
	AssigneeID                  int         `json:"assignee_id"`
	AuthorID                    int         `json:"author_id"`
	CreatedAt                   string      `json:"created_at"`
	Description                 string      `json:"description"`
	HeadPipelineID              int         `json:"head_pipeline_id"`
	ID                          int         `json:"id"`
	IID                         int         `json:"iid"`
	LastEditedAt                string      `json:"last_edited_at"`
	LastEditedByID              int         `json:"last_edited_by_id"`
	MergeCommitSHA              interface{} `json:"merge_commit_sha"`
	MergeError                  interface{} `json:"merge_error"`
	MergeParams                 MergeParams `json:"merge_params"`
	MergeStatus                 string      `json:"merge_status"`
	MergeUserID                 interface{} `json:"merge_user_id"`
	MergeWhenPipelineSucceeds   bool        `json:"merge_when_pipeline_succeeds"`
	MilestoneID                 interface{} `json:"milestone_id"`
	SourceBranch                string      `json:"source_branch"`
	SourceProjectID             int         `json:"source_project_id"`
	StateID                     int         `json:"state_id"`
	TargetBranch                string      `json:"target_branch"`
	TargetProjectID             int         `json:"target_project_id"`
	TimeEstimate                int         `json:"time_estimate"`
	Title                       string      `json:"title"`
	UpdatedAt                   string      `json:"updated_at"`
	UpdatedByID                 int         `json:"updated_by_id"`
	URL                         string      `json:"url"`
	Source                      Project     `json:"source"`
	Target                      Project     `json:"target"`
	LastCommit                  LastCommit  `json:"last_commit"`
	WorkInProgress              bool        `json:"work_in_progress"`
	TotalTimeSpent              int         `json:"total_time_spent"`
	TimeChange                  int         `json:"time_change"`
	HumanTotalTimeSpent         interface{} `json:"human_total_time_spent"`
	HumanTimeChange             interface{} `json:"human_time_change"`
	HumanTimeEstimate           interface{} `json:"human_time_estimate"`
	AssigneeIDs                 []int       `json:"assignee_ids"`
	ReviewerIDs                 []int       `json:"reviewer_ids"`
	Labels                      []string    `json:"labels"`
	State                       string      `json:"state"`
	BlockingDiscussionsResolved bool        `json:"blocking_discussions_resolved"`
	FirstContribution           bool        `json:"first_contribution"`
	DetailedMergeStatus         string      `json:"detailed_merge_status"`
	Action                      string      `json:"action"`
}

type MergeParams struct {
	ForceRemoveSourceBranch string `json:"force_remove_source_branch"`
}

type LastCommit struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Title     string `json:"title"`
	Timestamp string `json:"timestamp"`
	URL       string `json:"url"`
	Author    Author `json:"author"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Changes struct {
	Description  Description  `json:"description"`
	LastEditedAt LastEditedAt `json:"last_edited_at"`
	UpdatedAt    UpdatedAt    `json:"updated_at"`
}

type Description struct {
	Previous string `json:"previous"`
	Current  string `json:"current"`
}

type LastEditedAt struct {
	Previous string `json:"previous"`
	Current  string `json:"current"`
}

type UpdatedAt struct {
	Previous string `json:"previous"`
	Current  string `json:"current"`
}

type Repository struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
}

type ApprovalUser struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	State     string `json:"state"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

type Rule struct {
	ID                int            `json:"id"`
	Name              string         `json:"name"`
	RuleType          string         `json:"rule_type"`
	EligibleApprovers []ApprovalUser `json:"eligible_approvers"`
	ApprovalsRequired int            `json:"approvals_required"`
	Groups            []string       `json:"groups"`
	ContainsHidden    bool           `json:"contains_hidden_groups"`
	Section           string         `json:"section"`
	SourceRule        *Rule          `json:"source_rule"`
	Overridden        bool           `json:"overridden"`
}
