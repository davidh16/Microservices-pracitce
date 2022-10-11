package logger

import (
	"context"
	"createuser-service/cmd/models"
	pb "createuser-service/cmd/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func LogIt(model models.User) {
	conn, err := grpc.Dial("logger-service:80", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Printf("could not connect to logger service: %v", err)
		return
	}
	defer conn.Close()

	c := pb.NewLoggerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.LogEvent(ctx, &pb.LogRequest{
		Name: "user registration",
		Data: model.Name + model.Surname,
	})
	if err != nil {
		log.Printf("could not log event: %v", err)
		return
	}
}
