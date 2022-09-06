package api

import (
	jp "github.com/benpayflic/go-ml-pipelined/data-ingestion-service/internal/application/domain/job_posting"
)

type pipelineErr struct {
	err error
}
type generator func() <-chan jp.JobPosting
type consumer func(chan pipelineErr, <-chan jp.JobPosting)

// Push raw job postings into the pipeline
func jobPostingGenerator(postings *[]jp.JobPosting) generator {
	return func() <-chan jp.JobPosting {
		out := make(chan jp.JobPosting)
		go func() {
			defer close(out)
			for _, p := range *postings {
				out <- p
			}
		}()
		return out
	}
}

// Push processed job postings to database
func jobPostingConsumer(api Application) consumer {
	return func(errC chan pipelineErr, in <-chan jp.JobPosting) {
		batch := make([]jp.JobPosting, 0)
		for {
			p, ok := <-in
			if ok {
				batch = append(batch, p)
			} else {
				err := api.db.CreateJobPostings(batch)
				if err != nil {
					errC <- pipelineErr{err: err}
					break
				}
				errC <- pipelineErr{err: nil} //signal no error occurred
				break
			}
		}
	}
}
