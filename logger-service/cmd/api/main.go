package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"logger-service/cmd/initializers"
	"logger-service/cmd/mapper"
	pb "logger-service/cmd/proto"
	"net"
)

func init() {
	initializers.ConnectToDB()
	initializers.Migrate()
}

type LoggerServer struct {
	pb.UnimplementedLoggerServiceServer
}

func (l *LoggerServer) LogEvent(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error) {
	logModel := mapper.ProtoToModel(req)
	err := logModel.CreateLog(initializers.DB)
	if err != nil {
		log.Fatalf("could not log: %v", err)
		return nil, err
	}

	response := mapper.ModelToProto(&logModel)

	return &response, nil

}

const webPort = "80"

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", webPort))

	if err != nil {
		log.Printf(webPort)
		log.Fatalf("could not start listening: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterLoggerServiceServer(s, &LoggerServer{})

	reflection.Register(s)

	log.Printf("Server running on port %v", webPort)
	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("server could not start running: %v", err)

	}

}
