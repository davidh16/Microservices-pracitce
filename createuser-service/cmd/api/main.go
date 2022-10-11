package main

import (
	"context"
	"createuser-service/cmd/initializers"
	"createuser-service/cmd/logger"
	"createuser-service/cmd/mapper"
	pb "createuser-service/cmd/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func init() {
	initializers.ConnectToDB()
	initializers.Migrate()
}

type RegistrationServiceServer struct {
	pb.UnimplementedUserRegistrationServer
}

func (r *RegistrationServiceServer) Registration(ctx context.Context, req *pb.UserRegistrationRequest) (*pb.User, error) {
	registeredUser := mapper.ProtoToModel(req)
	err := registeredUser.Register(initializers.DB)
	if err != nil {
		log.Printf("user could not be registered: %v", err)
		return nil, err
	}
	log.Println("User successfully registered")
	response := mapper.ModelToProto(&registeredUser)

	logger.LogIt(registeredUser)

	return &response, nil
}

const webPort = "80"

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", webPort))
	if err != nil {
		log.Fatalf("could not start listening: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterUserRegistrationServer(s, &RegistrationServiceServer{})

	reflection.Register(s)

	log.Printf("Server running on port %v", webPort)
	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("server could not start running: %v", err)

	}

}
