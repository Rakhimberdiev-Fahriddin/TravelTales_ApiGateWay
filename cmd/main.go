package main

import (
	"api/api/handler"
	"api/config"
	"api/logs"
	"api/service"
	"log"
	"api/api"
)

func main() {
	cfg := config.Load()
	logs.InitLogger()
	logs.Logger.Info("Starting the application ...")

	srv := service.New()
	h := handler.NewHandler(srv.Auth,srv.Content,logs.Logger)

	r := api.Router(h)
	logs.Logger.Info("Server is running", "PORT", cfg.HTTP_PORT)
	log.Fatal(r.Run(cfg.HTTP_PORT))
}
