package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"

	core_logger "github.com/shitaiv1ck/todolist/internal/core/logger"
	core_server "github.com/shitaiv1ck/todolist/internal/core/server"
	core_postgres "github.com/shitaiv1ck/todolist/internal/core/store/postgres"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		panic(err)
	}

	store := core_postgres.NewStore(logger)
	if err := store.Open(); err != nil {
		panic(err)
	}

	defer store.Close()

	router := http.NewServeMux()

	server := core_server.NewServer(router, logger)

	if err := server.Run(ctx); err != nil {
		panic(err)
	}
}
