//+build wireinject

package app

import (
	"github.com/google/wire"
	config "viniti.us/hashout/config/di"
	"viniti.us/hashout/handlers/server"
)

func SetupApplication() *server.Api {
	wire.Build(config.Container)
	return nil
}
