package database

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"malanka/compute"
	"malanka/storage"
	"os"
	"strings"
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

func (db *Database) RouteQuery(q *compute.Query) (string, error) {
	db.logger.Debug("routing query", zap.Any("query", q))

	var value string
	var err error

	switch q.Command {
	case "SET":
		err = db.runSetQuery(q)
	case "GET":
		value, err = db.runGetQuery(q)
	case "DEL":
		err = db.runDelQuery(q)
	default:
		err = fmt.Errorf("unknown command: %s", q.Command)
	}

	if err != nil {
		return "", err
	} else {
		return value, nil
	}
}

func (db *Database) runSetQuery(query *compute.Query) error {
	return db.storage.Set(query.Args[0], query.Args[1])
}

func (db *Database) runGetQuery(query *compute.Query) (string, error) {
	return db.storage.Get(query.Args[0])
}

func (db *Database) runDelQuery(query *compute.Query) error {
	return db.storage.Del(query.Args[0])
}
