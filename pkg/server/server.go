package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/kristiansantos/ms_first/pkg/env"
	"github.com/kristiansantos/ms_first/pkg/logger"
	"github.com/kristiansantos/ms_first/pkg/middleware"
	"github.com/kristiansantos/ms_first/pkg/mongodb"
	"github.com/kristiansantos/ms_first/pkg/routes"
)

type server struct {
	Addr        string
	Port        int
	MongodbConn mongodb.Storage
	HttpServer  *fiber.App
}

func New(addr string, port int) *server {
	return &server{
		Addr: addr,
		Port: port,
	}
}

func (s *server) Run(app env.Application, log *logger.Logger) error {
	log.Info("server.main.Run", fmt.Sprintf("Server running on port :%d", s.Port))
	log.Info("server.main.Run", fmt.Sprintf("Environment: %s", app.Environment))
	log.Info("server.main.Run", fmt.Sprintf("Version: %s", app.Version))

	s.mongodbStart()

	if s.MongodbConn.Error != nil {
		panic(fmt.Sprintf("Error connecting to mongodb: %v", s.MongodbConn.Error))
	}

	s.serverConfig(app)

	go s.startServerHttp()

	return nil
}

func (s *server) serverConfig(app env.Application) {
	config := fiber.Config{
		ReadTimeout:  app.ReadTimeout * 2,
		WriteTimeout: app.WriteTimeout * 2,
	}

	s.HttpServer = fiber.New(config)

	middleware.FiberMiddleware(s.HttpServer)

	routes.UserRoutes(s.HttpServer)
	routes.NotFoundRoute(s.HttpServer)
}

func (s *server) startServerHttp() error {
	serverHost := fmt.Sprintf("%s:%d", s.Addr, s.Port)

	if err := s.HttpServer.Listen(serverHost); err != nil {
		if err == http.ErrServerClosed {
			fmt.Println("Finish server")
		} else {
			fmt.Println(err)
		}

		process, err := os.FindProcess(os.Getpid())
		if err != nil {
			fmt.Println("Couldn't find process to exit")
			os.Exit(1)
		}

		if err := process.Signal(os.Interrupt); err != nil {
			fmt.Println("Couldn't find process to exit")
			os.Exit(1)
		}

	}
	return nil
}

func (s *server) mongodbStart() (mongoConnection mongodb.Storage) {
	ctx := context.TODO()
	mongoConnection = mongodb.New(ctx)
	return
}
