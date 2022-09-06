package main

import (
	"log"

	rightDB "github.com/benpayflic/go-ml-pipelined/data-ingestion-service/internal/adapters/framework/secondary/database"
	"github.com/benpayflic/go-ml-pipelined/data-ingestion-service/internal/application/api"
	jp "github.com/benpayflic/go-ml-pipelined/data-ingestion-service/internal/application/domain/job_posting"
	c "github.com/benpayflic/go-ml-pipelined/data-ingestion-service/pkg/config"
)

func main() {
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
	api := api.NewApplication(dbDrivenAdapter, config)

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

	err = api.CreateJobPostings(&jobPostings)
	if err != nil {
		log.Fatal(err)
	}
}
