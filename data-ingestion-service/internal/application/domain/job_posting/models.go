package job_posting

import (
	"encoding/json"
	"time"
)

type JobPosting struct {
	JobID              int    `json:"job_id" bson:"job_id"`
	Title              string `json:"title" bson:"title"`
	Location           string `json:"location" bson:"location"`
	Department         string `json:"department" bson:"department"`
	SalaryRange        string `json:"salary_range" bson:"salary_range"`
	CompanyProfile     string `json:"company_profile" bson:"company_profile"`
	Description        string `json:"description" bson:"description"`
	Requirements       string `json:"requirements" bson:"requirements"`
	Benefits           string `json:"benefits" bson:"benefits"`
	Telecommuting      bool   `json:"telecommuting" bson:"telecommuting"`
	HasCompanyLogo     bool   `json:"has_company_logo" bson:"has_company_logo"`
	HasQuestions       bool   `json:"has_questions" bson:"has_questions"`
	EmployeeType       string `json:"employee_type" bson:"employee_type"`
	RequiredExperience string `json:"required_experience" bson:"required_experience"`
	RequiredEducation  string `json:"required_education" bson:"required_education"`
	Industry           string `json:"industry" bson:"industry"`
	Function           string `json:"function" bson:"function"`
	Fraudulent         bool   `json:"fraudulent" bson:"fraudulent"`
	CreatedAt          int64  `json:"created_at" bson:"created_at"`
}

func UnmarshalJobPostings(data []byte) ([]JobPosting, error) {
	var r []JobPosting
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *JobPosting) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SearchFilterParams struct {
	CreatedBefore time.Time `json:"created_before"`
	CreatedAfter  time.Time `json:"created_after"`
}

func UnmarshalSearchFilterParams(data []byte) (SearchFilterParams, error) {
	var r SearchFilterParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchFilterParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
