package compute

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"strings"
)

// Commands contains how many arguments each command has
var Commands = map[string]int{
	"SET": 2,
	"GET": 1,
	"DEL": 1,
}

type Command string

func NewCommand(cmd string, logger *zap.Logger) (Command, error) {
	logger.Debug(fmt.Sprintf("detecting a command: %s", cmd))

	command := strings.ToUpper(cmd)

	if command != cmd {
		return "", errors.New(fmt.Sprintf("command must be uppercase: %s", cmd))
	}

	if _, exists := Commands[command]; !exists {
		return "", errors.New(fmt.Sprintf("command must be one of SET, GET or DEL; %s given", command))
	}

	logger.Debug(fmt.Sprintf("command validated: %s", command))

	return Command(command), nil
}
