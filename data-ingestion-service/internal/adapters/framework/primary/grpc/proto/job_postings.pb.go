// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: job_postings.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type JobPosting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId              int32  `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Title              string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Location           string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Department         string `protobuf:"bytes,4,opt,name=department,proto3" json:"department,omitempty"`
	SalaryRange        string `protobuf:"bytes,5,opt,name=salary_range,json=salaryRange,proto3" json:"salary_range,omitempty"`
	CompanyProfile     string `protobuf:"bytes,6,opt,name=company_profile,json=companyProfile,proto3" json:"company_profile,omitempty"`
	Description        string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Requirements       string `protobuf:"bytes,8,opt,name=requirements,proto3" json:"requirements,omitempty"`
	Benefits           string `protobuf:"bytes,9,opt,name=benefits,proto3" json:"benefits,omitempty"`
	Telecommuting      bool   `protobuf:"varint,10,opt,name=telecommuting,proto3" json:"telecommuting,omitempty"`
	HasCompanyLogo     bool   `protobuf:"varint,11,opt,name=has_company_logo,json=hasCompanyLogo,proto3" json:"has_company_logo,omitempty"`
	HasQuestions       bool   `protobuf:"varint,12,opt,name=has_questions,json=hasQuestions,proto3" json:"has_questions,omitempty"`
	EmployeeType       string `protobuf:"bytes,13,opt,name=employee_type,json=employeeType,proto3" json:"employee_type,omitempty"`
	RequiredExperience string `protobuf:"bytes,14,opt,name=required_experience,json=requiredExperience,proto3" json:"required_experience,omitempty"`
	RequiredEducation  string `protobuf:"bytes,15,opt,name=required_education,json=requiredEducation,proto3" json:"required_education,omitempty"`
	Industry           string `protobuf:"bytes,16,opt,name=industry,proto3" json:"industry,omitempty"`
	Function           string `protobuf:"bytes,17,opt,name=function,proto3" json:"function,omitempty"`
	Fraudulent         bool   `protobuf:"varint,18,opt,name=fraudulent,proto3" json:"fraudulent,omitempty"`
}

func (x *JobPosting) Reset() {
	*x = JobPosting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_postings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobPosting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobPosting) ProtoMessage() {}

func (x *JobPosting) ProtoReflect() protoreflect.Message {
	mi := &file_job_postings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobPosting.ProtoReflect.Descriptor instead.
func (*JobPosting) Descriptor() ([]byte, []int) {
	return file_job_postings_proto_rawDescGZIP(), []int{0}
}

func (x *JobPosting) GetJobId() int32 {
	if x != nil {
		return x.JobId
	}
	return 0
}

func (x *JobPosting) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *JobPosting) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *JobPosting) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *JobPosting) GetSalaryRange() string {
	if x != nil {
		return x.SalaryRange
	}
	return ""
}

func (x *JobPosting) GetCompanyProfile() string {
	if x != nil {
		return x.CompanyProfile
	}
	return ""
}

func (x *JobPosting) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *JobPosting) GetRequirements() string {
	if x != nil {
		return x.Requirements
	}
	return ""
}

func (x *JobPosting) GetBenefits() string {
	if x != nil {
		return x.Benefits
	}
	return ""
}

func (x *JobPosting) GetTelecommuting() bool {
	if x != nil {
		return x.Telecommuting
	}
	return false
}

func (x *JobPosting) GetHasCompanyLogo() bool {
	if x != nil {
		return x.HasCompanyLogo
	}
	return false
}

func (x *JobPosting) GetHasQuestions() bool {
	if x != nil {
		return x.HasQuestions
	}
	return false
}

func (x *JobPosting) GetEmployeeType() string {
	if x != nil {
		return x.EmployeeType
	}
	return ""
}

func (x *JobPosting) GetRequiredExperience() string {
	if x != nil {
		return x.RequiredExperience
	}
	return ""
}

func (x *JobPosting) GetRequiredEducation() string {
	if x != nil {
		return x.RequiredEducation
	}
	return ""
}

func (x *JobPosting) GetIndustry() string {
	if x != nil {
		return x.Industry
	}
	return ""
}

func (x *JobPosting) GetFunction() string {
	if x != nil {
		return x.Function
	}
	return ""
}

func (x *JobPosting) GetFraudulent() bool {
	if x != nil {
		return x.Fraudulent
	}
	return false
}

type CreateJobPostingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobPosting *JobPosting `protobuf:"bytes,1,opt,name=job_posting,json=jobPosting,proto3" json:"job_posting,omitempty"`
}

func (x *CreateJobPostingsRequest) Reset() {
	*x = CreateJobPostingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_postings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobPostingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobPostingsRequest) ProtoMessage() {}

func (x *CreateJobPostingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_postings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobPostingsRequest.ProtoReflect.Descriptor instead.
func (*CreateJobPostingsRequest) Descriptor() ([]byte, []int) {
	return file_job_postings_proto_rawDescGZIP(), []int{1}
}

func (x *CreateJobPostingsRequest) GetJobPosting() *JobPosting {
	if x != nil {
		return x.JobPosting
	}
	return nil
}

type StreamJobPostingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedBefore *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=created_before,json=createdBefore,proto3,oneof" json:"created_before,omitempty"`
	CreatedAfter  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_after,json=createdAfter,proto3,oneof" json:"created_after,omitempty"`
}

func (x *StreamJobPostingsRequest) Reset() {
	*x = StreamJobPostingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_postings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamJobPostingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamJobPostingsRequest) ProtoMessage() {}

func (x *StreamJobPostingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_postings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamJobPostingsRequest.ProtoReflect.Descriptor instead.
func (*StreamJobPostingsRequest) Descriptor() ([]byte, []int) {
	return file_job_postings_proto_rawDescGZIP(), []int{2}
}

func (x *StreamJobPostingsRequest) GetCreatedBefore() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedBefore
	}
	return nil
}

func (x *StreamJobPostingsRequest) GetCreatedAfter() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAfter
	}
	return nil
}

type CreateJobPostingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateJobPostingsResponse) Reset() {
	*x = CreateJobPostingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_postings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobPostingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobPostingsResponse) ProtoMessage() {}

func (x *CreateJobPostingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_postings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobPostingsResponse.ProtoReflect.Descriptor instead.
func (*CreateJobPostingsResponse) Descriptor() ([]byte, []int) {
	return file_job_postings_proto_rawDescGZIP(), []int{3}
}

func (x *CreateJobPostingsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_job_postings_proto protoreflect.FileDescriptor

var file_job_postings_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x04, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x61, 0x6c, 0x61, 0x72, 0x79, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x27,
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x65, 0x6c,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x74, 0x65, 0x6c, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x28, 0x0a, 0x10, 0x68, 0x61, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x6c,
	0x6f, 0x67, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x68, 0x61, 0x73, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x68, 0x61, 0x73,
	0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x68, 0x61, 0x73, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f,
	0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x5f, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x66,
	0x72, 0x61, 0x75, 0x64, 0x75, 0x6c, 0x65, 0x6e, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x66, 0x72, 0x61, 0x75, 0x64, 0x75, 0x6c, 0x65, 0x6e, 0x74, 0x22, 0x55, 0x0a, 0x18, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0b, 0x6a, 0x6f, 0x62, 0x5f, 0x70,
	0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6a,
	0x6f, 0x62, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x50,
	0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x6a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x22, 0xcd, 0x01, 0x0a, 0x18, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x62,
	0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x46, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x65,
	0x66, 0x6f, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x44, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x0c, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x66, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a,
	0x0f, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65,
	0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x66, 0x74,
	0x65, 0x72, 0x22, 0x35, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50,
	0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xce, 0x01, 0x0a, 0x0b, 0x4a, 0x6f,
	0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x66, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x26,
	0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x6f, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50,
	0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28,
	0x01, 0x12, 0x57, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x62, 0x50, 0x6f,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x26, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x6f, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x62, 0x50,
	0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x4a, 0x6f,
	0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x30, 0x01, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x6e, 0x70, 0x61, 0x79, 0x66,
	0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x6c, 0x2d, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_job_postings_proto_rawDescOnce sync.Once
	file_job_postings_proto_rawDescData = file_job_postings_proto_rawDesc
)

func file_job_postings_proto_rawDescGZIP() []byte {
	file_job_postings_proto_rawDescOnce.Do(func() {
		file_job_postings_proto_rawDescData = protoimpl.X.CompressGZIP(file_job_postings_proto_rawDescData)
	})
	return file_job_postings_proto_rawDescData
}

var file_job_postings_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_job_postings_proto_goTypes = []interface{}{
	(*JobPosting)(nil),                // 0: job_postings.JobPosting
	(*CreateJobPostingsRequest)(nil),  // 1: job_postings.CreateJobPostingsRequest
	(*StreamJobPostingsRequest)(nil),  // 2: job_postings.StreamJobPostingsRequest
	(*CreateJobPostingsResponse)(nil), // 3: job_postings.CreateJobPostingsResponse
	(*timestamppb.Timestamp)(nil),     // 4: google.protobuf.Timestamp
}
var file_job_postings_proto_depIdxs = []int32{
	0, // 0: job_postings.CreateJobPostingsRequest.job_posting:type_name -> job_postings.JobPosting
	4, // 1: job_postings.StreamJobPostingsRequest.created_before:type_name -> google.protobuf.Timestamp
	4, // 2: job_postings.StreamJobPostingsRequest.created_after:type_name -> google.protobuf.Timestamp
	1, // 3: job_postings.JobPostings.CreateJobPostings:input_type -> job_postings.CreateJobPostingsRequest
	2, // 4: job_postings.JobPostings.StreamJobPostings:input_type -> job_postings.StreamJobPostingsRequest
	3, // 5: job_postings.JobPostings.CreateJobPostings:output_type -> job_postings.CreateJobPostingsResponse
	0, // 6: job_postings.JobPostings.StreamJobPostings:output_type -> job_postings.JobPosting
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_job_postings_proto_init() }
func file_job_postings_proto_init() {
	if File_job_postings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_job_postings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobPosting); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_job_postings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJobPostingsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_job_postings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamJobPostingsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_job_postings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJobPostingsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_job_postings_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_job_postings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_job_postings_proto_goTypes,
		DependencyIndexes: file_job_postings_proto_depIdxs,
		MessageInfos:      file_job_postings_proto_msgTypes,
	}.Build()
	File_job_postings_proto = out.File
	file_job_postings_proto_rawDesc = nil
	file_job_postings_proto_goTypes = nil
	file_job_postings_proto_depIdxs = nil
}