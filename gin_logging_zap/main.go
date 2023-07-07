package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
    //logging
    "go.uber.org/zap"
    ginzap "github.com/gin-contrib/zap"
)

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
    r.Use(ginzap.RecoveryWithZap(logger,true))
    r.Use(ginzap.Ginzap(logger,time.RFC3339,true))

	r.GET("/", func(c *gin.Context) {
		c.String(302, "Request to: /")
	})
    
    r.GET("/panic", func(c *gin.Context) {
        panic("PANIC!!")
    })

	fmt.Println("Now listening on :8080")
	fmt.Println(r.Run("0.0.0.0:8080"))
}
