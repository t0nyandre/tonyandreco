package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/t0nyandre/tonyandreco/internal/config"
)

func New(cfg *config.Config) (*zerolog.Logger, error) {
	dateString := time.Now().Format("20060102")

	file, err := os.OpenFile(
		fmt.Sprintf("logs/%s.log", dateString),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		return nil, err
	}

	multi := zerolog.MultiLevelWriter(os.Stdout, file)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logger := zerolog.New(multi).With().Timestamp().Str("app_name", cfg.Name).Logger()

	logger.Debug().Msg("logger created")
	return &logger, nil
}
