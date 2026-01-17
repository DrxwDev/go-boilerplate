package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/DrxwDev/boilerplate/internal/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewHTTPServer(lc fx.Lifecycle, router *gin.Engine, app *config.AppConfig, srv *config.ServerConfig) *http.Server {
	server := &http.Server{
		Addr:              ":" + app.PORT,
		Handler:           router,
		ReadHeaderTimeout: srv.ReadHeaderTimeout,
		ReadTimeout:       srv.ReadTimeout,
		WriteTimeout:      srv.WriteTimout,
		IdleTimeout:       srv.IdleTimeot,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Server starting at PORT", app.PORT)
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("Error starting the server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Shutting down the server...")
			return server.Shutdown(ctx)
		},
	})

	return server
}
