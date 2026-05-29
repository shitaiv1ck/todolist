package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"

	core_logger "github.com/shitaiv1ck/todolist/internal/core/logger"
	core_server "github.com/shitaiv1ck/todolist/internal/core/server"
	core_postgres "github.com/shitaiv1ck/todolist/internal/core/store/postgres"
	core_middleware "github.com/shitaiv1ck/todolist/internal/core/transport/middleware"
	users_repository "github.com/shitaiv1ck/todolist/internal/features/users/repository"
	users_service "github.com/shitaiv1ck/todolist/internal/features/users/service"
	users_transport "github.com/shitaiv1ck/todolist/internal/features/users/transport"
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

	usersRepository := users_repository.NewRepository(store)
	usersService := users_service.NewService(usersRepository)
	usersTransport := users_transport.NewTransport(usersService)

	router := http.NewServeMux()
	router.Handle("POST /users", usersTransport.CreateUserHandler())

	public := core_middleware.PublicMiddleware(router, logger)

	server := core_server.NewServer(public, logger)

	if err := server.Run(ctx); err != nil {
		panic(err)
	}
}
