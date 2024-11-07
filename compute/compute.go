package compute

import (
	"errors"
	"go.uber.org/zap"
	"strings"
)

type Compute struct {
	logger *zap.Logger
}

func NewCompute(logger *zap.Logger) *Compute {
	logger.Debug("initializing compute layer")

	return &Compute{
		logger: logger,
	}
}

func (c *Compute) Parse(q string) (Query, error) {
	c.logger.Debug("parsing a new query", zap.String("q", q))

	data := strings.Fields(q)

	if len(data) == 0 {
		return Query{}, errors.New("the query is empty")
	}

	command, err := NewCommand(data[0], c.logger)
	if err != nil {
		return Query{}, err
	}

	query, err := NewQuery(command, data[1:], c.logger)
	if err != nil {
		return Query{}, err
	}

	return *query, nil
}
