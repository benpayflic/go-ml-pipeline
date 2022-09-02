package api

import (
	jp "github.com/benpayflic/go-ml-pipelined/data-ingestion-service/internal/application/domain/job_posting"
)

func (api Application) CreateJobPostings(postings []jp.JobPosting) error {

	errC := make(chan pipelineErr)

	go jobPostingConsumer()(
		errC, jobPostingGenerator(postings)())

	res := <-errC
	if res.err != nil {
		return res.err
	}

	return nil
}

func (api Application) StreamJobPostings(params jp.SearchFilterParams) (*[]jp.JobPosting, error) {
	// TODO: Implement function
	return nil, nil
}
