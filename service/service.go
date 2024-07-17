package service

import (
	"api/config"
	"api/generated/auth_service"
	"api/generated/content_service"
	"api/logs"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	Auth    auth_service.AuthServiceClient
	Content content_service.ContentClient
}

func New() *Service {
	cfg := config.Load()
	logs.InitLogger()

	connAuth, err := grpc.NewClient("localhost"+cfg.AUTH_PORT,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		logs.Logger.Error("Failed to created Auth client connection","error",err.Error())
		log.Println(err)
		return nil
	}

	authService := auth_service.NewAuthServiceClient(connAuth)

	connContent,err := grpc.NewClient("localhost"+cfg.CONTENT_PORT,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		logs.Logger.Error("Failed to create Content client connection","error",err.Error())
		log.Println(err)
		return nil
	}

	contentService := content_service.NewContentClient(connContent)


	return &Service{
		Content: contentService,
		Auth:authService,
	}
}
