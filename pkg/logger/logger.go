package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var Instance *logger = nil

type logger struct {
	client zerolog.Logger
}

var _ ILoggerProvider = &logger{}

func New() logger {
	if Instance == nil {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log := zerolog.New(os.Stderr).With().Timestamp().Logger()

		Instance = &logger{
			client: log,
		}
	}

	return *Instance
}

func GetInstance() logger {
	return *Instance
}

func (l logger) Error(namespace string, message string) {
	l.client.Error().Str("namespace", namespace).Msg(message)
}

func (l logger) Info(namespace string, message string) {
	l.client.Info().Str("namespace", namespace).Msg(message)
}
