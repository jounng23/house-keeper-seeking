package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"

	"pricing-svc/internal/di/configfx"
	"pricing-svc/internal/di/dbfx"
	"pricing-svc/internal/di/handlerfx"
	"pricing-svc/internal/di/pricefx"
	"pricing-svc/internal/di/routerfx"
	"pricing-svc/internal/server/router"
	"pricing-svc/pkg/config"
)

func main() {
	app := fx.New(
		configfx.Initialize(".env"),
		dbfx.Module,
		handlerfx.Module,
		pricefx.Module,
		routerfx.Module,
		fx.Provide(provideGinEngine),
		fx.Invoke(registerHTTPService,
			startHTTPServer),
	)
	app.Run()
}

func provideGinEngine() *gin.Engine {
	r := gin.New()
	return r
}

func registerHTTPService(g *gin.Engine,
	router router.Router) {
	api := g.Group("/api/v1")
	router.Register(api)
}

func startHTTPServer(lifecycle fx.Lifecycle, g *gin.Engine) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				port := fmt.Sprintf("%d", config.ServerConfig().Port)
				log.Info().Msgf("Pricing service is listening on port: %s", port)
				go func() {
					server := http.Server{
						Addr:    ":" + port,
						Handler: g,
					}
					if err := server.ListenAndServe(); err != nil {
						log.Error().Msgf("failed to listen and serve from server: %v", err.Error())
					}
				}()
				return nil
			},
			OnStop: func(context.Context) error {
				log.Info().Msg("service stopped")
				return nil
			},
		},
	)
}
