package main

import (
	db "example.com/myapp/database"
	logger "example.com/myapp/logging"
)

func main() {
	db.Init()
	logger.Log("This gets logged from logger to console")
}
