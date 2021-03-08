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
	"strconv"
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

	_ = server.Serve(list)
}

func (s *server) Create(ctx context.Context, id *proto.UserID) (*proto.Result, error) {
	_, err := s.sessionsUsecase.CreateSession(strconv.Itoa(int(id.ID)))
	if err != nil {
		fmt.Println(err)
		result := &proto.Result{
			ID:     id,
			Status: err.Error(),
		}
		return result, err
	}
	result := &proto.Result{
		ID:     id,
		Status: "OK",
	}
	return result, nil
}

func (s *server) Check(ctx context.Context, id *proto.UserID) (*proto.Result, error) {
	_, err := s.sessionsUsecase.CheckSession(strconv.Itoa(int(id.ID)))
	if err != nil {
		fmt.Println(err)
		result := &proto.Result{
			ID:     id,
			Status: err.Error(),
		}
		return result, err
	}
	result := &proto.Result{
		ID:     id,
		Status: "OK",
	}
	return result, nil
}

func (s *server) Delete(ctx context.Context, id *proto.UserID) (*proto.Result, error) {
	err := s.sessionsUsecase.DeleteSession(strconv.Itoa(int(id.ID)))
	if err != nil {
		fmt.Println(err)
		result := &proto.Result{
			ID:     id,
			Status: err.Error(),
		}
		return result, err
	}
	result := &proto.Result{
		ID:     id,
		Status: "OK",
	}
	return result, nil
}
