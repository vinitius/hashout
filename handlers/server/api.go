package server

import (
	"net/http"
	"time"

	"viniti.us/hashout/config/log"

	"viniti.us/hashout/handlers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Api struct {
	*http.Server
}

// @title Hashout - Cart API
// @version 1.0
// @description Rest API.
// @termsOfService https://viniti.us/terms

// @contact.name API Support
// @contact.url https://viniti.us/contact
// @contact.email salomao.tcn@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8181
// @BasePath /
// @query.collection.format multi
func NewHttpServer(router *gin.Engine, c handlers.CheckoutHandler) *http.Server {
	router.Use(gin.Recovery())
	router.Use(handlers.ApiErrors())
	c.Routes(router)

	// simple healthcheck
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// swagger files
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := viper.GetString("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	return &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func NewRouter() *gin.Engine {
	return gin.New()
}

func (n Api) ListenAndServe() error {
	log.Logger.Infof("Starting hashout http server on port %s", n.Addr)
	return n.Server.ListenAndServe()
}
