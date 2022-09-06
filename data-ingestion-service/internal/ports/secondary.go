package ports

import (
	jp "github.com/benpayflic/go-ml-pipelined/data-ingestion-service/internal/application/domain/job_posting"
)

// Driven ports

type DBPort interface {
	Connect()
	CreateJobPostings([]jp.JobPosting) error
	RetrieveJobPostings(jp.SearchFilterParams, int, int) ([]jp.JobPosting, error)
	DeleteJobPostings([]jp.JobPosting) error
}
