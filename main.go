package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("error: missing URL arg")
	}
    workflowFileName := os.Args[1]

    owner := "skovranek"
    repo := "learn-cicd-starter"

    jobsLogs, err := getWorkflowRunJobsLogs(owner, repo, workflowFileName)	
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(jobsLogs)
}
// first thing that gets the list of workflows by filename and givhes the needed run ID
// ListWorkflowRunsByFileName lists all workflow runs by workflow file name.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/workflow-runs#list-workflow-runs
//	u := fmt.Sprintf("repos/%s/%s/actions/workflows/%v/runs", owner, repo, workflowFileName)
