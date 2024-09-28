package database

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"os"
	"strings"
)

func (db *Database) Listen() {
	reader := bufio.NewReader(os.Stdin)

	for {
		db.logger.Debug("waiting for a query")

		fmt.Print("> ")
		req, err := reader.ReadString('\n')
		if err != nil {
			db.logger.Error("error while reading a query: "+err.Error(), zap.Error(err))
			continue
		}

		req = strings.TrimSpace(req)

		if req == "exit" {
			fmt.Println("bye")
			return
		}

		db.HandleRequest(req)
	}
}
