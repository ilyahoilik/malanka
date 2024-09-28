package storage

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"malanka/storage/engine"
)

type Storage struct {
	engine engine.Engine
	logger *zap.Logger
}

func NewStorage(logger *zap.Logger) *Storage {
	logger.Info("initializing storage")

	return &Storage{
		engine: engine.NewHashTable(),
		logger: logger,
	}
}

func (s *Storage) Set(key string, value string) error {
	s.logger.Info(fmt.Sprintf("setting %s to storage", key))
	s.engine.Set(key, value)
	return nil
}

func (s *Storage) Get(key string) (string, error) {
	s.logger.Info(fmt.Sprintf("getting %s from storage", key))
	if value, found := s.engine.Get(key); !found {
		return "", errors.New("key not found")
	} else {
		return value, nil
	}
}

func (s *Storage) Del(key string) error {
	s.logger.Info(fmt.Sprintf("deleting %s from storage", key))
	if deleted := s.engine.Del(key); !deleted {
		s.logger.Info(fmt.Sprintf("trying to delete non-existence key %s", key))
		return errors.New("key not found")
	} else {
		return nil
	}
}
