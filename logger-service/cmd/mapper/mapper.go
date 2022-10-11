package mapper

import (
	"logger-service/cmd/models"
	pb "logger-service/cmd/proto"
)

func ProtoToModel(proto *pb.LogRequest) models.Log {
	log := models.Log{
		Name: proto.Name,
		Data: proto.Data,
	}
	return log
}

func ModelToProto(model *models.Log) pb.LogResponse {
	proto := pb.LogResponse{
		Name: model.Name,
		Data: model.Data,
		Id:   model.ID,
	}

	return proto
}
