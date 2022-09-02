package api

import (
	"fmt"

	jp "github.com/benpayflic/go-ml-pipelined/data-ingestion-service/internal/application/domain/job_posting"
)

type pipelineErr struct {
	err error
}
type generator func() <-chan jp.JobPosting
type consumer func(chan pipelineErr, <-chan jp.JobPosting)

// Push raw job postings into the pipeline
func jobPostingGenerator(postings []jp.JobPosting) generator {
	return func() <-chan jp.JobPosting {
		out := make(chan jp.JobPosting)
		go func() {
			defer close(out)
			for _, p := range postings {
				out <- p
			}
		}()
		return out
	}
}

// Push processed job postings to database
func jobPostingConsumer() consumer {
	return func(errC chan pipelineErr, in <-chan jp.JobPosting) {
		for {
			p, ok := <-in
			if ok {
				fmt.Println(p.Location)
				// TODO: Push to DB

			} else {
				// processed all data
				errC <- pipelineErr{err: nil} //signal no error occurred
				return
			}
		}
	}
}
