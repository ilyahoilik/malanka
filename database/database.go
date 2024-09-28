package database

import (
	"fmt"
	"go.uber.org/zap"
	"malanka/compute"
	"malanka/storage"
)

type Database struct {
	storage storage.Storage
	logger  *zap.Logger
}

func NewDatabase(logger *zap.Logger) *Database {
	logger.Debug("initializing database")

	return &Database{
		storage: *storage.NewStorage(logger),
		logger:  logger,
	}
}

func (db *Database) HandleRequest(req string) {
	parser := compute.NewCompute(db.logger)

	query, err := parser.Parse(req)
	if err != nil {
		db.logger.Error("error while parsing query: ", zap.Error(err))
		fmt.Printf("[err]: %s\n", err)
	}

	val, err := db.RouteQuery(&query)

	if err != nil {
		db.logger.Error("error while running query", zap.Error(err))
		fmt.Printf("[err]: %s\n", err)
	} else {
		db.logger.Debug("query ran successful")
		if len(val) > 0 {
			fmt.Println(val)
		}
	}
}
