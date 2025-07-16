package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func Init() {
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.DateTime,
	}

	log.Logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		log.Info().
			Str("METHOD", c.Request.Method).
			Str("PATH", c.Request.URL.Path).
			Int("STATUS", c.Writer.Status()).
			Dur("LATENCY", duration).
			Msg("HTTP Request")
	}
}
