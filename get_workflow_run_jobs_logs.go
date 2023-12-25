package main

import (
    "encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
    "strings"
)

func getWorkflowRunJobsLogs(owner, repo, workflowFileName string) ([]string, error) {
    listWorkflowRunsURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/workflows/%v/runs", owner, repo, workflowFileName)
    // 1) make request to get list of workflow runs
    fmt.Printf("Making request to get list of workflow runs:\n'%s'\n", listWorkflowRunsURL)

	res, err := http.Get(listWorkflowRunsURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return nil, errors.New("res failed")
	}

    body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	workflowRuns := WorkflowRuns{}
	err = json.Unmarshal(body, &workflowRuns)
	if err != nil {
        return nil, err
	}

    // 2) get workflow run JobsURL
    workflowRunJobsURL := *workflowRuns.WorkflowRuns[0].JobsURL
    
    //workflowRunJobsURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/runs/%s/jobs", owner, repo, workflowID)

    // 3) get list of workflow run jobs from above URL
	fmt.Printf("Making request to get list of workflow run jobs:\n'%s'\n", workflowRunJobsURL)

	res, err = http.Get(workflowRunJobsURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return nil, errors.New("res failed")
	}

    body, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	jobs := Jobs{}
	err = json.Unmarshal(body, &jobs)
	if err != nil {
        return nil, err
	}

	jobsContainAll := map[string][]string{
        "Tests": []string{
            "test",
            "go test",
        },
        "Style": []string{
            "staticcheck",
            "go fmt",
        },
        "Deploy": []string{
            "install",
        },
	}

    result := []string{}
    bearerToken := "Bearer "

	for _, job := range jobs.Jobs {
		for jobName, contains := range jobsContainAll {
			if *job.Name == jobName {
                // 4) get job ID
                jobID := *job.ID
                workflowRunJobLogURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/jobs/%v/logs", owner, repo, jobID)
                // 5) get job log
                fmt.Printf("Making request to get workflow run job log:\n'%s'\n", workflowRunJobLogURL)

                req, err := http.NewRequest("GET", workflowRunJobLogURL, nil)
                if err != nil {
                    return nil, err
                }

                req.Header.Set("Authorization", bearerToken)

                client := &http.Client{}
                res, err := client.Do(req)
                if err != nil {
                    return nil, err
                }
                defer res.Body.Close()

                if res.StatusCode > 302 {
                    return nil, errors.New("res failed")
                }

                jobLog, err := io.ReadAll(res.Body)
                if err != nil {
                    return nil, err
                }
          
                // 6) check job log contains terms
                for _, contain := range contains {
                    if !strings.Contains(string(jobLog), contain) {
                        return nil, fmt.Errorf("'%s' job log does not contain '%s'", jobName, contain)
                    }
                    fmt.Printf("'%s' found\n", contain)
                }

                jobResult := fmt.Sprintf("%s: %v",jobName, jobID)
                result = append(result, jobResult)
                break
            }
        }
    }

    return result, nil
}
