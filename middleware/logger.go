package middleware

import (
	"io"
	"time"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

var Log = logger.SetLogger(
	logger.WithLogger(func(c *gin.Context, out io.Writer, latency time.Duration) zerolog.Logger {
		return zerolog.New(out).With().
			Str("path", c.Request.URL.Path).
			Str("Method", c.Request.Method).
			Dur("latency", latency).
			Logger()
	}),
)
