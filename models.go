package main

import "time"

// Timestamp represents a time that can be unmarshalled from a JSON string
// formatted as either an RFC3339 or Unix timestamp. This is necessary for some
// fields since the GitHub API is inconsistent in how it represents times. All
// exported methods of time.Time can be called on Timestamp.
type Timestamp struct {
	time.Time
}

// TaskStep represents a single task step from a sequence of tasks of a job.
type TaskStep struct {
	Name        *string    `json:"name,omitempty"`
	Status      *string    `json:"status,omitempty"`
	Conclusion  *string    `json:"conclusion,omitempty"`
	Number      *int64     `json:"number,omitempty"`
	StartedAt   *Timestamp `json:"started_at,omitempty"`
	CompletedAt *Timestamp `json:"completed_at,omitempty"`
}

// WorkflowJob represents a repository action workflow job.
type WorkflowJob struct {
	ID          *int64      `json:"id,omitempty"`
	RunID       *int64      `json:"run_id,omitempty"`
	RunURL      *string     `json:"run_url,omitempty"`
	NodeID      *string     `json:"node_id,omitempty"`
	HeadBranch  *string     `json:"head_branch,omitempty"`
	HeadSHA     *string     `json:"head_sha,omitempty"`
	URL         *string     `json:"url,omitempty"`
	HTMLURL     *string     `json:"html_url,omitempty"`
	Status      *string     `json:"status,omitempty"`
	Conclusion  *string     `json:"conclusion,omitempty"`
	CreatedAt   *Timestamp  `json:"created_at,omitempty"`
	StartedAt   *Timestamp  `json:"started_at,omitempty"`
	CompletedAt *Timestamp  `json:"completed_at,omitempty"`
	Name        *string     `json:"name,omitempty"`
	Steps       []*TaskStep `json:"steps,omitempty"`
	CheckRunURL *string     `json:"check_run_url,omitempty"`
	// Labels represents runner labels from the `runs-on:` key from a GitHub Actions workflow.
	Labels          []string `json:"labels,omitempty"`
	RunnerID        *int64   `json:"runner_id,omitempty"`
	RunnerName      *string  `json:"runner_name,omitempty"`
	RunnerGroupID   *int64   `json:"runner_group_id,omitempty"`
	RunnerGroupName *string  `json:"runner_group_name,omitempty"`
	RunAttempt      *int64   `json:"run_attempt,omitempty"`
	WorkflowName    *string  `json:"workflow_name,omitempty"`
}

// Jobs represents a slice of repository action workflow job.
type Jobs struct {
	TotalCount *int           `json:"total_count,omitempty"`
	Jobs       []*WorkflowJob `json:"jobs,omitempty"`
}

// WorkflowRun represents a repository action workflow run.
type WorkflowRun struct {
	ID                 *int64         `json:"id,omitempty"`
	Name               *string        `json:"name,omitempty"`
	NodeID             *string        `json:"node_id,omitempty"`
	HeadBranch         *string        `json:"head_branch,omitempty"`
	HeadSHA            *string        `json:"head_sha,omitempty"`
	RunNumber          *int           `json:"run_number,omitempty"`
	RunAttempt         *int           `json:"run_attempt,omitempty"`
	Event              *string        `json:"event,omitempty"`
	DisplayTitle       *string        `json:"display_title,omitempty"`
	Status             *string        `json:"status,omitempty"`
	Conclusion         *string        `json:"conclusion,omitempty"`
	WorkflowID         *int64         `json:"workflow_id,omitempty"`
	CheckSuiteID       *int64         `json:"check_suite_id,omitempty"`
	CheckSuiteNodeID   *string        `json:"check_suite_node_id,omitempty"`
	URL                *string        `json:"url,omitempty"`
	HTMLURL            *string        `json:"html_url,omitempty"`
	// PullRequests       []*PullRequest `json:"pull_requests,omitempty"`
	CreatedAt          *Timestamp     `json:"created_at,omitempty"`
	UpdatedAt          *Timestamp     `json:"updated_at,omitempty"`
	RunStartedAt       *Timestamp     `json:"run_started_at,omitempty"`
	JobsURL            *string        `json:"jobs_url,omitempty"`
	LogsURL            *string        `json:"logs_url,omitempty"`
	CheckSuiteURL      *string        `json:"check_suite_url,omitempty"`
	ArtifactsURL       *string        `json:"artifacts_url,omitempty"`
	CancelURL          *string        `json:"cancel_url,omitempty"`
	RerunURL           *string        `json:"rerun_url,omitempty"`
	PreviousAttemptURL *string        `json:"previous_attempt_url,omitempty"`
	// HeadCommit         *HeadCommit    `json:"head_commit,omitempty"`
	WorkflowURL        *string        `json:"workflow_url,omitempty"`
	// Repository         *Repository    `json:"repository,omitempty"`
	// HeadRepository     *Repository    `json:"head_repository,omitempty"`
	// Actor              *User          `json:"actor,omitempty"`
}

// WorkflowRuns represents a slice of repository action workflow run.
type WorkflowRuns struct {
	TotalCount   *int           `json:"total_count,omitempty"`
	WorkflowRuns []*WorkflowRun `json:"workflow_runs,omitempty"`
}
