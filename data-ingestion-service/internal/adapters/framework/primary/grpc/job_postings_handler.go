package grpc

import (
	"fmt"
	"io"
	"log"

	pb "github.com/benpayflic/go-ml-pipeline/data-ingestion-service/internal/adapters/framework/primary/grpc/proto"
	jp "github.com/benpayflic/go-ml-pipeline/data-ingestion-service/internal/application/domain/job_posting"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateJobPostings(stream pb.JobPostings_CreateJobPostingsServer) error {

	for {

		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("EOF from client. No more incoming requests")
			return stream.SendAndClose(&pb.CreateJobPostingsResponse{Message: "Job postings successfully created!"})
		}
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while reading client stream: %v", err),
			)
		}

		posting := jp.NewJobPosting()
		posting.JobID = int(req.JobPosting.JobId)
		posting.Title = req.JobPosting.Title
		posting.Location = req.JobPosting.Location
		posting.Department = req.JobPosting.Department
		posting.SalaryRange = req.JobPosting.SalaryRange
		posting.CompanyProfile = req.JobPosting.CompanyProfile
		posting.Description = req.JobPosting.Description
		posting.Requirements = req.JobPosting.Requirements
		posting.Benefits = req.JobPosting.Benefits
		posting.Telecommuting = req.JobPosting.Telecommuting
		posting.HasCompanyLogo = req.JobPosting.HasCompanyLogo
		posting.HasQuestions = req.JobPosting.HasQuestions
		posting.EmployeeType = req.JobPosting.EmployeeType
		posting.RequiredExperience = req.JobPosting.RequiredExperience
		posting.RequiredEducation = req.JobPosting.RequiredEducation
		posting.Industry = req.JobPosting.Industry
		posting.Function = req.JobPosting.Function
		posting.Fraudulent = req.JobPosting.Fraudulent

		postings := jp.NewJobPostings()
		*postings = append(*postings, *posting)

		err = s.adapter.api.CreateJobPostings(postings)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Internal Error: %v", err),
			)
		}
	}
}

func (s *Server) StreamJobPostings(in *pb.StreamJobPostingsRequest, stream pb.JobPostings_StreamJobPostingsServer) error {

	params := jp.NewSearchFilterParams()
	if in.CreatedAfter.AsTime().Unix() != 0 {
		params.CreatedAfter = in.CreatedAfter.AsTime()
	}
	if in.CreatedBefore.AsTime().Unix() != 0 {
		params.CreatedBefore = in.CreatedBefore.AsTime()
	}

	page := 1
	limit := 100
	for {
		res, err := s.adapter.api.StreamJobPostings(*params, page, limit)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Internal Error: %v", err),
			)
		}
		if len(res) < 1 {
			log.Println("No results found to return with provided params. Cursor exhausted.")
			break
		}

		log.Printf("%v results found. Streaming ...", len(res))

		for _, p := range res {
			stream.Send(&pb.JobPosting{
				JobId:              int32(p.JobID),
				Title:              p.Title,
				Location:           p.Location,
				Department:         p.Department,
				SalaryRange:        p.SalaryRange,
				CompanyProfile:     p.CompanyProfile,
				Description:        p.Description,
				Requirements:       p.Requirements,
				Benefits:           p.Benefits,
				Telecommuting:      p.Telecommuting,
				HasCompanyLogo:     p.HasCompanyLogo,
				HasQuestions:       p.HasQuestions,
				EmployeeType:       p.EmployeeType,
				RequiredExperience: p.RequiredExperience,
				RequiredEducation:  p.RequiredEducation,
				Industry:           p.Industry,
				Function:           p.Function,
				Fraudulent:         p.Fraudulent,
			})
		}
		page += 1
	}
	return nil
}
