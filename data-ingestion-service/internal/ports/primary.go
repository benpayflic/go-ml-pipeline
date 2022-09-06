package ports

import (
	jp "github.com/benpayflic/go-ml-pipelined/data-ingestion-service/internal/application/domain/job_posting"
)

type APIPort interface {
	CreateJobPostings(*[]jp.JobPosting) error
	StreamJobPostings(jp.SearchFilterParams, int, int) ([]jp.JobPosting, error)
}
