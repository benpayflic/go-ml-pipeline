package job_posting

import "encoding/json"

type JobPosting struct {
	JobID              int    `json:"job_id"`
	Title              string `json:"title"`
	Location           string `json:"location"`
	Department         string `json:"department"`
	SalaryRange        string `json:"salary_range"`
	CompanyProfile     string `json:"company_profile"`
	Description        string `json:"description"`
	Requirements       string `json:"requirements"`
	Benefits           string `json:"benefits"`
	Telecommuting      bool   `json:"telecommuting"`
	HasCompanyLogo     bool   `json:"has_company_logo"`
	HasQuestions       bool   `json:"has_questions"`
	EmployeeType       string `json:"employee_type"`
	RequiredExperience string `json:"required_experience"`
	RequiredEducation  string `json:"required_education"`
	Industry           string `json:"industry"`
	Function           string `json:"function"`
	Fraudulent         bool   `json:"fraudulent"`
}

func UnmarshalJobPosting(data []byte) (JobPosting, error) {
	var r JobPosting
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *JobPosting) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SearchFilterParams struct {
	Title      string `json:"title"`
	Industry   string `json:"industry"`
	Location   string `json:"location"`
	Fraudulent bool   `json:"fraudulent"`
}

func UnmarshalSearchFilterParams(data []byte) (SearchFilterParams, error) {
	var r SearchFilterParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchFilterParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
