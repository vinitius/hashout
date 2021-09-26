//+build wireinject

package app

import (
	"github.com/google/wire"
	"viniti.us/hashout/config/di"
	"viniti.us/hashout/handlers/server"
)

func SetupApplication() *server.Api {
	wire.Build(di.Container)
	return nil
}
