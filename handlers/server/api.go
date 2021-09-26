package server

import (
	"net/http"
	"time"

	"viniti.us/hashout/config/log"

	"viniti.us/hashout/handlers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Api struct {
	*http.Server
}

func NewHttpServer(router *gin.Engine, oh handlers.CheckoutHandler) *http.Server {
	router.Use(gin.Recovery())
	router.Use(handlers.ApiErrors())
	oh.Routes(router)

	// simple healthcheck
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

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
	log.Logger.Infof("Starting hahsout http server on port %s", n.Addr)
	return n.Server.ListenAndServe()
}
