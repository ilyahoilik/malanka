package database

import (
	"go.uber.org/zap"
	"malanka/compute"
)

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
