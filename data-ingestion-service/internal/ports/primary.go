package ports

import (
	jp "github.com/benpayflic/go-ml-pipelined/data-ingestion-service/internal/application/domain/job_posting"
)

type APIPort interface {
	CreateJobPostings(postings []jp.JobPosting) error
	StreamJobPostings(params jp.SearchFilterParams) (*[]jp.JobPosting, error)
}
