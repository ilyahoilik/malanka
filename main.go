package main

import (
	"fmt"
	"malanka/database"
	"malanka/logger"
)

func main() {
	log, err := logger.NewLogger()
	if err != nil {
		fmt.Println("failed to initialize logger")
		return
	}
	defer logger.SyncLogger(log)

	db := database.NewDatabase(log)
	db.Listen()
}
