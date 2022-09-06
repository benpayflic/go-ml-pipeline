package api

import (
	"errors"
	"log"
	"testing"

	rightDB "github.com/benpayflic/go-ml-pipelined/data-ingestion-service/internal/adapters/framework/secondary/database"
	c "github.com/benpayflic/go-ml-pipelined/data-ingestion-service/pkg/config"

	jp "github.com/benpayflic/go-ml-pipelined/data-ingestion-service/internal/application/domain/job_posting"
)

func TestCreateJobPostings(t *testing.T) {
	// config, err := c.LoadConfig("./pkg/config/")
	config, err := c.LoadConfig("/Users/bengoodenough/Documents/Portfolio-Projects/go-ml-pipeline/data-ingestion-service/pkg/config/")
	if err != nil {
		log.Fatalln("Failed to load application configurations : ", err)
	}

	dbDrivenAdapter, err := rightDB.NewAdapter(&config)
	if err != nil {
		log.Fatalf("failed to initialize db driven adapter: %v", err)
	}
	dbDrivenAdapter.Connect()

	api := NewApplication(dbDrivenAdapter, config)

	jobPostings := []jp.JobPosting{
		{
			JobID:              1,
			Title:              "Senior Go Engineer (TEST - DELETE ME)",
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
			Title:              "Senior Python Engineer (TEST - DELETE ME)",
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
			JobID:              3,
			Title:              "Senior C# Engineer (TEST - DELETE ME)",
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

	err = api.CreateJobPostings(&jobPostings)
	if err != nil {
		t.Error(err.Error())
	}

	// Clean up test DB entries.
	api.DeleteJobPostings(jobPostings)
}

func TestStreamJobPostings(t *testing.T) {
	// config, err := c.LoadConfig("./pkg/config/")
	config, err := c.LoadConfig("/Users/bengoodenough/Documents/Portfolio-Projects/go-ml-pipeline/data-ingestion-service/pkg/config/")
	if err != nil {
		log.Fatalln("Failed to load application configurations : ", err)
	}

	dbDrivenAdapter, err := rightDB.NewAdapter(&config)
	if err != nil {
		log.Fatalf("failed to initialize db driven adapter: %v", err)
	}
	dbDrivenAdapter.Connect()

	api := NewApplication(dbDrivenAdapter, config)

	// Create Job Posting
	jobPostings := []jp.JobPosting{
		{
			JobID:              2,
			Title:              "Senior Python Engineer (TEST - DELETE ME)",
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
	api.CreateJobPostings(&jobPostings)

	// Search for new job posting
	params := jp.SearchFilterParams{
		Title: "Senior Python Engineer (TEST - DELETE ME)",
	}

	postings, err := api.StreamJobPostings(params, 1, 1)

	if err != nil {
		t.Error(err.Error())
	}

	if len(postings) == 0 {
		t.Error(errors.New("no job posting found with title " + params.Title))
	}

	// Clean up test DB entries.
	api.DeleteJobPostings(jobPostings)

}
