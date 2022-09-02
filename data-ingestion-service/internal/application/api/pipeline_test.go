package api

import (
	"testing"

	jp "github.com/benpayflic/go-ml-pipelined/data-ingestion-service/internal/application/domain/job_posting"
)

func TestJobpostingPipeline(t *testing.T) {

	api := NewApplication()

	jobPostings := []jp.JobPosting{
		{
			JobID:              1,
			Title:              "Senior Go Engineer",
			Location:           "",
			Department:         "",
			SalaryRange:        "",
			CompanyProfile:     "",
			Description:        "",
			Requirements:       "",
			Benefits:           "",
			Telecommuting:      true,
			HasCompanyLogo:     false,
			HasQuestions:       true,
			EmployeeType:       "",
			RequiredExperience: "",
			RequiredEducation:  "",
			Industry:           "industry",
			Function:           "function",
			Fraudulent:         false,
		}, {
			JobID:              2,
			Title:              "Senior Python Engineer",
			Location:           "",
			Department:         "",
			SalaryRange:        "",
			CompanyProfile:     "",
			Description:        "",
			Requirements:       "",
			Benefits:           "",
			Telecommuting:      true,
			HasCompanyLogo:     false,
			HasQuestions:       true,
			EmployeeType:       "",
			RequiredExperience: "",
			RequiredEducation:  "",
			Industry:           "industry",
			Function:           "function",
			Fraudulent:         false,
		},
	}

	err := api.CreateJobPostings(jobPostings)
	if err != nil {
		t.Error(err.Error())
	}
}
