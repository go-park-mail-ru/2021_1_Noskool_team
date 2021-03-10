package client

import (
	proto "2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/protobuff"
	"2021_1_Noskool_team/internal/microservices/auth/models"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

type SessionsClient struct {
	client proto.AuthCheckerClient
}

func NewSessionsClient(con *grpc.ClientConn) *SessionsClient {
	client := proto.NewAuthCheckerClient(con)
	return &SessionsClient{
		client: client,
	}
}

func (sesClient *SessionsClient) Create(ctx context.Context, id string) (models.Result, error) {
	UserID := &proto.UserID{ID: id}
	result, err := sesClient.client.Create(ctx, UserID, grpc.EmptyCallOption{})
	if err != nil {
		fmt.Println(err)
		return models.Result{}, err
	}
	return transformIntoResultModel(result), nil
}
func (sesClient *SessionsClient) Check(ctx context.Context, hash string) (models.Result, error) {
	Hash := &proto.Hash{Hash: hash}
	result, err := sesClient.client.Check(ctx, Hash, grpc.EmptyCallOption{})
	if err != nil {
		fmt.Println(err)
		return models.Result{}, err
	}
	return transformIntoResultModel(result), nil
}

func (sesClient *SessionsClient) Delete(ctx context.Context, hash string) (models.Result, error) {
	Hash := &proto.Hash{Hash: hash}
	result, err := sesClient.client.Delete(ctx, Hash, grpc.EmptyCallOption{})
	if err != nil {
		fmt.Println(err)
		return models.Result{}, err
	}
	return transformIntoResultModel(result), nil
}

func transformIntoResultModel(result *proto.Result) models.Result {
	if result == nil {
		return models.Result{}
	}
	resultModel := models.Result{
		ID:     result.ID.ID,
		Hash:   result.Hash,
		Status: result.Status,
	}

	return resultModel
}
