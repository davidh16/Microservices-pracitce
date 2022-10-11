package main

import (
	"auth-service/cmd/initializers"
	"auth-service/cmd/logger"
	"auth-service/cmd/models"
	pb "auth-service/cmd/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func init() {
	initializers.ConnectToDB()
}

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
}

func (a *AuthServiceServer) AuthenticateUser(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	user := models.User{}
	result := initializers.DB.Where("email=?", req.Email).First(&user)
	if result.Error != nil {
		response := pb.AuthResponse{
			Authenticated: false,
			Message:       "Wrong credentials",
		}

		return &response, result.Error
	}

	auth, err := user.PasswordMatches(req.Password)
	if err != nil {
		log.Panic(err)
	}
	if !auth {
		response := pb.AuthResponse{
			Authenticated: false,
			Message:       "Wrong credentials",
		}

		return &response, nil
	}

	response := pb.AuthResponse{
		Authenticated: true,
		Message:       "Authentication successfull",
	}

	logger.LogIt(user)

	return &response, nil
}

const webPort = "80"

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", webPort))
	if err != nil {
		log.Fatalf("could not start listening: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterAuthServiceServer(s, &AuthServiceServer{})

	reflection.Register(s)

	log.Printf("Server running on port %v", webPort)
	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("server could not start running: %v", err)
	}
}
