package handler

import (
	protobuf2 "api/api/protobuf"
	"context"
)

type MessageServer struct {
	protobuf2.UnimplementedMessageServiceServer
}

func (s *MessageServer) GetMessage(ctx context.Context, request *protobuf2.GetMessageRequest) (*protobuf2.MessageResponse, error) {
	return &protobuf2.MessageResponse{
		Id:      request.MessageId,
		Content: "Hello world!",
	}, nil
}
