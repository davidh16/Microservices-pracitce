package mapper

import (
	"createuser-service/cmd/models"
	pb "createuser-service/cmd/proto"
)

func ProtoToModel(proto *pb.UserRegistrationRequest) models.User {
	userModel := models.User{
		Name:     proto.Name,
		Surname:  proto.Surname,
		Email:    proto.Email,
		Password: proto.Password,
	}
	return userModel
}

func ModelToProto(model *models.User) pb.User {
	proto := pb.User{
		Name:    model.Name,
		Surname: model.Surname,
		Id:      model.ID,
	}

	return proto
}
