package main

import (
	"go.uber.org/zap"
)

func main() {

	cfg := zap.NewProductionConfig()
	cfg.DisableStacktrace = true
	//To log to console you can use "stdout" here:
	cfg.OutputPaths = []string{
		"./access.log",
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Info("I am the msg part", zap.Int("newkey", 3))
}
