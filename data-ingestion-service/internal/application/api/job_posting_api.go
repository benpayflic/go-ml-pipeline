package api

import (
	jp "github.com/benpayflic/go-ml-pipeline/data-ingestion-service/internal/application/domain/job_posting"
)

func (api Application) CreateJobPostings(postings *[]jp.JobPosting) error {

	errC := make(chan pipelineErr)

	go jobPostingConsumer(api)(
		errC, jobPostingGenerator(postings)())

	res := <-errC
	if res.err != nil {
		return res.err
	}

	return nil
}

func (api Application) StreamJobPostings(params jp.SearchFilterParams, page int, limit int) ([]jp.JobPosting, error) {
	jobPostings, err := api.db.RetrieveJobPostings(params, page, limit)
	if err != nil {
		return nil, err
	}
	return jobPostings, nil
}

// Used for test clean up
func (api Application) DeleteJobPostings(postings []jp.JobPosting) error {
	err := api.db.DeleteJobPostings(postings)
	if err != nil {
		return err
	}
	return nil
}
