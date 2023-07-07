package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
    //Use Metrics before Recovery to also get metrics about http500 panics
    RegisterMetrics(r)
    r.Use(gin.Recovery())
    // If you use "github.com/gin-contrib/zap"
    // Skip the logging of /metrics request with:
    /*
        r.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
            TimeFormat: time.RFC3339,
            UTC:        true,
            SkipPaths:  []string{"/metrics"},
        }))
    */

	r.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})
	r.GET("/panic", func(c *gin.Context) {
		panic("PANIC!!")
	})
	r.GET("/invalid", func(c *gin.Context) {
		c.String(400, "invalid data")
	})

	fmt.Println("Now listening on *:8080")
	r.Run("0.0.0.0:8080")
}
