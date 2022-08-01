package cmd

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/kristiansantos/ms_first/initializer/env"
	"github.com/kristiansantos/ms_first/initializer/server"
	"github.com/kristiansantos/ms_first/pkg/logger"
	"github.com/spf13/cobra"
)

var (
	port        int
	addr        string
	environment string
	version     string
	apiCmd      = &cobra.Command{
		Use:   "api",
		Short: "Start HTTP server",
		Long: `
	Commands to start server:
	-p This flag option specified port HTTP server, default are 3000.
	-a This flag option binds specified IP, by default it is localhost.
	-e This flag option specified the environment.
	-v This flag option specified version to deploy
	`,
		Run: cmdRun,
	}
)

func init() {
	rootCmd.AddCommand(apiCmd)

	// Get start server options
	apiCmd.PersistentFlags().IntVarP(&port, "port", "p", 3000, "The -p option specified port HTTP server")
	apiCmd.PersistentFlags().StringVarP(&addr, "address", "a", "127.0.0.1", "The -b option binds specified IP, by default it is localhost")
	apiCmd.PersistentFlags().StringVarP(&environment, "environment", "e", "development", "The -e option specified the environment")
	apiCmd.PersistentFlags().StringVarP(&version, "version", "v", os.Getenv("VERSION"), "The -v option specified version to deploy")
}

func cmdRun(cmd *cobra.Command, args []string) {
	if err := godotenv.Load("./configs/.env." + environment); err != nil {
		panic(err)
	}

	os.Getenv("ENV")

	cfg, err := env.ReadEnvironments(environment, version)

	if err != nil {
		panic(err)
	}

	fmt.Println(cfg)

	svr := server.New(addr, port)
	log := logger.New()
	svr.Run(cfg, log)

	chanExit := make(chan os.Signal, 1)
	signal.Notify(chanExit, os.Interrupt)
	<-chanExit

}
