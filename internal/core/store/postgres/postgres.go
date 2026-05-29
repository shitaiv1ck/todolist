package core_postgres

import (
	"database/sql"
	"fmt"

	core_logger "github.com/shitaiv1ck/todolist/internal/core/logger"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

type Store struct {
	config Config
	db     *sql.DB
	logger *core_logger.Logger
}

func NewStore(logger *core_logger.Logger) *Store {
	return &Store{
		config: NewConfigMust(),
		logger: logger,
	}
}

func (s *Store) GetDB() *sql.DB {
	return s.db
}

func (s *Store) Open() error {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		s.config.User,
		s.config.Password,
		s.config.Host,
		s.config.Port,
		s.config.DB,
	)

	s.logger.Debug("opening the database")

	db, err := sql.Open("postgres", url)
	if err != nil {
		s.logger.Error("open the database", zap.Error(err))

		return err
	}

	if err := db.Ping(); err != nil {
		s.logger.Error("ping the database", zap.Error(err))

		return err
	}

	s.db = db

	s.logger.Info("the database was opened")

	return nil
}

func (s *Store) Close() {
	s.logger.Debug("closing the database")

	if err := s.db.Close(); err != nil {
		s.logger.Panic("failed to close the database", zap.Error(err))

		panic(err)
	}

	s.logger.Info("the database was closed")
}
