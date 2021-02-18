package server

import (
	"2021_1_Noskool_team/internal/microservices/auth"
	proto "2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/protobuff"
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type server struct {
	sessionsUsecase auth.Usecase
}

func NewSessionsServerGRPC(gServer *grpc.Server, sesUsecase auth.Usecase) {
	serv:= &server{
		sessionsUsecase: sesUsecase,
	}
	proto.RegisterAuthCheckerServer(gServer, serv)
	reflection.Register(gServer)
}

func StartAuthGRPCServer(sesUsecase auth.Usecase, url string) {
	list, err := net.Listen("tcp", url)
	if err != nil {
		logrus.Error(err)
	}

	server := grpc.NewServer()

	NewSessionsServerGRPC(server, sesUsecase)

	_ = server.Serve(list)
}

func (s *server) Create(ctx context.Context, id *proto.UserID) (*proto.Result, error) {
	panic("implement me")
}

func (s *server) Check(ctx context.Context, id *proto.UserID) (*proto.Result, error) {
	panic("implement me")
}

func (s *server) Delete(ctx context.Context, id *proto.UserID) (*proto.Result, error) {
	panic("implement me")
}
