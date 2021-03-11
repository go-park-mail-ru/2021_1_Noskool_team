package server

import (
	"2021_1_Noskool_team/internal/microservices/auth"
	proto "2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/protobuff"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type server struct {
	sessionsUsecase auth.Usecase
}

func NewSessionsServerGRPC(gServer *grpc.Server, sesUsecase auth.Usecase) {
	serv := &server{
		sessionsUsecase: sesUsecase,
	}
	proto.RegisterAuthCheckerServer(gServer, serv)
	reflection.Register(gServer)
}

func StartSessionsGRPCServer(sesUsecase auth.Usecase, url string) {
	logrus.Info(url)
	list, err := net.Listen("tcp", url)
	if err != nil {
		logrus.Error(err)
	}

	server := grpc.NewServer()

	NewSessionsServerGRPC(server, sesUsecase)

	err = server.Serve(list)
	if err != nil {
		logrus.Error(err)
	}
}

func (s *server) Create(ctx context.Context, id *proto.UserID) (*proto.Result, error) {
	session, err := s.sessionsUsecase.CreateSession(id.ID)
	if err != nil {
		fmt.Println(err)
		result := &proto.Result{
			ID:     id,
			Status: err.Error(),
		}
		return result, err
	}

	userID := &proto.UserID{
		ID: session.UserID,
	}
	result := &proto.Result{
		ID:     userID,
		Hash:   session.Hash,
		Status: "OK",
	}
	return result, nil
}

func (s *server) Check(ctx context.Context, hash *proto.Hash) (*proto.Result, error) {
	session, err := s.sessionsUsecase.CheckSession(hash.Hash)
	if err != nil {
		fmt.Println(err)
		userID := &proto.UserID{
			ID: "-1",
		}
		result := &proto.Result{
			ID:     userID,
			Status: err.Error(),
		}
		return result, err
	}
	userID := &proto.UserID{
		ID: session.UserID,
	}
	result := &proto.Result{
		ID:     userID,
		Hash:   session.Hash,
		Status: "OK",
	}
	return result, nil
}

func (s *server) Delete(ctx context.Context, hash *proto.Hash) (*proto.Result, error) {
	err := s.sessionsUsecase.DeleteSession(hash.Hash)

	userID := &proto.UserID{
		ID: "-1",
	}
	if err != nil {
		fmt.Println(err)

		result := &proto.Result{
			ID:     userID,
			Status: err.Error(),
		}
		return result, err
	}
	result := &proto.Result{
		ID:     userID,
		Status: "OK",
	}
	return result, nil
}
