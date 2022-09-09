package grpc

import (
	"log"
	"net"
	"os"

	pb "github.com/benpayflic/go-ml-pipeline/data-ingestion-service/internal/adapters/framework/primary/grpc/proto"

	"github.com/benpayflic/go-ml-pipeline/data-ingestion-service/internal/ports"
	"github.com/benpayflic/go-ml-pipeline/data-ingestion-service/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api    ports.APIPort
	config *config.Config
}

func NewAdapter(api ports.APIPort, config *config.Config) *Adapter {
	return &Adapter{api: api, config: config}
}

var (
	ErrLog = log.New(os.Stderr, "[GRPC][ERROR] ", log.LstdFlags|log.Lmsgprefix)
	Log    = log.New(os.Stdout, "[GRPC][INFO] ", log.LstdFlags|log.Lmsgprefix)
)

type Server struct {
	pb.JobPostingsServer

	adapter Adapter
}

func (adapter Adapter) Start() {

	lis, err := net.Listen("tcp", adapter.config.GRPCAddress)
	if err != nil {
		ErrLog.Fatal("Failed to listen: ", adapter.config.GRPCAddress)
	}

	Log.Println("Listening on ", adapter.config.GRPCAddress)

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterJobPostingsServer(s, &Server{adapter: adapter})

	if err = s.Serve(lis); err != nil {
		log.Fatal("Failed to serve", err)
	}

}
