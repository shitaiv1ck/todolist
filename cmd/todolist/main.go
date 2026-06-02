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
	sessions_repository "github.com/shitaiv1ck/todolist/internal/features/sessions/repository"
	sessions_service "github.com/shitaiv1ck/todolist/internal/features/sessions/service"
	sessions_transport "github.com/shitaiv1ck/todolist/internal/features/sessions/transport"
	tasks_repository "github.com/shitaiv1ck/todolist/internal/features/tasks/repository"
	tasks_service "github.com/shitaiv1ck/todolist/internal/features/tasks/service"
	tasks_transport "github.com/shitaiv1ck/todolist/internal/features/tasks/transport"
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

	sessionsRepository := sessions_repository.NewRepository(store)
	sessionsService := sessions_service.NewService(usersRepository, sessionsRepository)
	sessionsTransport := sessions_transport.NewTransport(sessionsService)

	tasksRepisotory := tasks_repository.NewRepository(store)
	tasksService := tasks_service.NewService(tasksRepisotory)
	tasksTransport := tasks_transport.NewTransport(tasksService)

	private := http.NewServeMux()
	private.Handle("GET /users/me", usersTransport.GetMeHandler())
	privateRouter := core_middleware.ChainAuthenticated(private, sessionsService)

	protected := http.NewServeMux()
	protected.Handle("POST /tasks", tasksTransport.CreateTaskHandler())
	protected.Handle("DELETE /sessions", sessionsTransport.DeleteSessionHandler())
	protectedRouter := core_middleware.ChainProtected(protected, sessionsService)

	common := http.NewServeMux()
	common.Handle("POST /api/users", usersTransport.CreateUserHandler())
	common.Handle("POST /api/sessions", sessionsTransport.CreateSessionHandler())
	common.Handle("/api/", http.StripPrefix("/api", privateRouter))
	common.Handle("/api/protected/", http.StripPrefix("/api/protected", protectedRouter))
	routerAPI := core_middleware.ChainCommon(common, logger)

	server := core_server.NewServer(routerAPI, logger)

	if err := server.Run(ctx); err != nil {
		panic(err)
	}
}
