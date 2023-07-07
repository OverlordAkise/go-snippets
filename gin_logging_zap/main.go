package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	//logging
	ginzap "github.com/gin-contrib/zap"
	"go.uber.org/zap"
)

func AddLogging(r *gin.Engine, logger *zap.Logger) {
	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		t := time.Now()
		logger.Info("webrequest",
			zap.String("url", c.Request.URL.String()),
			zap.Int("ret", c.Writer.Status()),
			zap.String("ip", c.ClientIP()),
			zap.Duration("duration", t.Sub(start)),
			zap.Int("rsize", c.Writer.Size()),

			//If you use "go.opentelemetry.io/otel":
			//zap.String("trace_id",tr.SpanFromContext(c.Request.Context()).SpanContext().TraceID().String()),
		)
	})
}

func main() {

	cfg := zap.NewProductionConfig()
	//cfg.DisableStacktrace = true
	cfg.OutputPaths = []string{
		"./myapp.log", //Set this to "stdout" for logging to stdout
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	//Uncomment the above line and use the following function
	//if you want to use a custom gin logging function with more control:
	//AddLogging(r,logger)

	r.GET("/", func(c *gin.Context) {
		c.String(302, "Request to: /")
	})

	r.GET("/panic", func(c *gin.Context) {
		panic("PANIC!!")
	})

	fmt.Println("Now listening on :8080")
	fmt.Println(r.Run("0.0.0.0:8080"))
}
