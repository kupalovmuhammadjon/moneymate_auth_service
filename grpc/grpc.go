package grpc

import (
	"auth_service/pkg/logger"
	"auth_service/service"
	"auth_service/storage"
	pb "auth_service/genproto/users"

	"google.golang.org/grpc"
)

func SetUpServer(storage *storage.IStorage, log logger.ILogger) *grpc.Server {
	grpcServer := grpc.NewServer()

	userService := service.NewUsersService(*storage, log)

	pb.RegisterUsersServiceServer(grpcServer, userService)

	return grpcServer
}
