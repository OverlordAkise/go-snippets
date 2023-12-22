package main

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func main() {

	cfg := zap.NewProductionConfig()
	cfg.DisableStacktrace = false
	cfg.OutputPaths = []string{
		"stdout", // or  ./access.log
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	//Log panics with gin:
	app.Use(ginzap.RecoveryWithZap(logger, true))
	//Log normal web accesses with gin:
	app.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		SkipPaths:  []string{"/metrics"},
	}))

	app.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})

	app.GET("/extra", func(c *gin.Context) {
		logger.Info("Extra log event")
		c.String(200, "Extra OK")
	})
	fmt.Println("Now listening on http://localhost:8765")
	fmt.Println(app.Run(":8765"))
}
