package env

import (
	"fmt"
	"sync"

	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	SingletonInstance *Application = nil
	once              sync.Once
)

type (
	Mongo struct {
		User     string `env:"mongo_user"`
		Pass     string `env:"mongo_pass"`
		Host     string `env:"mongo_host"`
		Args     string `env:"mongo_args"`
		Database string `env:"mongo_database"`
	}
	Log struct {
		Level string `env:"LOG_LEVEL"`
	}

	Application struct {
		Environment  string
		Version      string
		Mongo        Mongo
		Log          Log
		ReadTimeout  time.Duration `env:"app_readTimeout"`
		WriteTimeout time.Duration `env:"app_writeTimeout"`
	}
)

func ReadEnvironments(environment, version string) (Application, error) {
	var app Application

	if err := cleanenv.ReadEnv(&app); err != nil {
		return Application{}, fmt.Errorf(`error reading env: %w`, err)
	}

	if SingletonInstance == nil {
		app.Environment = environment
		app.Version = version

		once.Do(
			func() {
				SingletonInstance = &app
			})
		return app, nil
	} else {
		return *SingletonInstance, nil
	}
}
