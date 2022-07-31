package initializer

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kristiansantos/ms_first/pkg/logger"
)

type server struct {
	Addr       string
	Port       int
	httpServer http.Server
}

func NewServer(addr string, port int) *server {
	return &server{
		Addr: addr,
		Port: port,
	}
}

func (s *server) HttpServerBuild(app Application) {
	s.httpServer = http.Server{
		Addr: fmt.Sprintf("%s:%d", s.Addr, s.Port),
		// Handler:      middleware.Recovery(router.Client),
		ReadTimeout:  app.ReadTimeout * 2,
		WriteTimeout: app.WriteTimeout * 2,
	}
}

func (s *server) StartServerHttp() error {
	if err := s.httpServer.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			fmt.Println("Finish server")
		} else {
			fmt.Println(err)
		}

		process, err := os.FindProcess(os.Getpid())
		if err != nil {
			fmt.Println("couldn't find process to exit")
			os.Exit(1)
		}

		if err := process.Signal(os.Interrupt); err != nil {
			fmt.Println("couldn't find process to exit")
			os.Exit(1)
		}

	}
	return nil
}

func (s *server) Run(app Application, log logger.ILoggerProvider) error {
	log.Info("server.main.Run", fmt.Sprintf("Server running on port :%d", s.Port))
	log.Info("server.main.Run", fmt.Sprintf("Environment: %s", app.Environment))
	log.Info("server.main.Run", fmt.Sprintf("Version: %s", app.Version))

	if connError := mongoConnetion(app); connError != nil {
		panic(fmt.Sprintf("Error connecting to mongodb: %v", connError.Error))
	}

	if connError := elasticsearchConnetion(app); connError != nil {
		panic(fmt.Sprintf("Error connecting to elasticsearch: %v", connError.Error))
	}

	s.HttpServerBuild(app)

	go s.StartServerHttp()

	return nil
}

func mongoConnetion(app Application) error {

	return nil
}

func elasticsearchConnetion(app Application) error {

	return nil
}
