package config

import (
	"viniti.us/hashout/clients"
	"viniti.us/hashout/config/app"
	"viniti.us/hashout/config/cb"
	"viniti.us/hashout/config/db"
	"viniti.us/hashout/config/grpc"
	"viniti.us/hashout/handlers"
	"viniti.us/hashout/handlers/server"
	"viniti.us/hashout/storage"
	"viniti.us/hashout/usecase"

	"github.com/google/wire"
)

var Container = wire.NewSet(
	app.Provider,
	cb.Provider,
	grpc.Provider,
	db.Provider,
	storage.Provider,
	usecase.Provider,
	clients.Provider,
	handlers.Provider,
	server.NewRouter,
	server.NewHttpServer,
	wire.Struct(new(server.Api), "*"),
)
