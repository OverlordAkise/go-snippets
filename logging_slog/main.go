package main

import (
	"log/slog"
	"os"
)

// NewTextHandler = logfmt format, NewJSONHandler = json format

func main() {
	// https://pkg.go.dev/log/slog#Level
	// -4=debug, 0=info, 4=warn, 8=error
	loglevel := -4
	logpath := "stdout"

	var logger *slog.Logger
	loglvl := new(slog.LevelVar)
	logconf := &slog.HandlerOptions{
		// AddSource: true, // this adds a source=/home/me/main.go:28 tag to loglines
		Level: loglvl,
	}
	if logpath == "stdout" {
		logger = slog.New(slog.NewTextHandler(os.Stdout, logconf))
	} else {
		f, err := os.OpenFile(logpath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		logger = slog.New(slog.NewTextHandler(f, logconf))
	}
	loglvl.Set(slog.Level(loglevel))
	logger.Info("message at info level", "ip", "127.0.0.1")
	logger.Error("mesage at error level", "mykey", 5, "err", nil)
}
