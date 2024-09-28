package compute

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"strings"
)

type Query struct {
	Command Command
	Args    []string
}

func NewQuery(cmd Command, args []string, logger *zap.Logger) (*Query, error) {
	logger.Debug(fmt.Sprintf("detecting a query: %s", string(cmd)+" "+strings.Join(args, " ")))

	if err := validateQuery(cmd, args); err != nil {
		return nil, err
	}

	logger.Debug(fmt.Sprintf("query validated: %s", string(cmd)+" "+strings.Join(args, " ")))

	return &Query{
		Command: cmd,
		Args:    args,
	}, nil
}

func validateQuery(cmd Command, args []string) error {
	actual, expected := len(args), Commands[string(cmd)]
	if actual != expected {
		return errors.New(fmt.Sprintf("expected %d arguments, got %d", expected, actual))
	}
	return nil
}
