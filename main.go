package main

import (
	"certification/internal/config"
	"certification/internal/handler/rest"
	"certification/internal/repository"
	server2 "certification/internal/server"
	"certification/internal/service"
	"certification/internal/storage/postgres"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("error initializating etc %s", err.Error())
	}

	db, err := postgres.InitConnect(cfg.Postgres)
	if err != nil {
		log.Fatalf("error initializing datatbase: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	transport := rest.NewHandler(service)
	router := server2.NewRouter(transport)

	//REST SERVER
	go func() {
		server := new(server2.Server)
		if err := server.Run(cfg.Server.PortREST, router.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
