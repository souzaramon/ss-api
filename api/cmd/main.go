package main

import (
	"api/internal/authors"
	"context"
	"net"
	"net/http"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func main() {
	fx.New(
		fx.Provide(
			NewGinApp,
			zap.NewProduction,
		),
		authors.Module,
		fx.Invoke(func(*gin.Engine) {}),
	).Run()
}

func NewGinApp(lc fx.Lifecycle, log *zap.Logger) *gin.Engine {
	g := gin.New()

	g.RedirectTrailingSlash = false

	g.Use(ginzap.Ginzap(log, time.RFC3339, true))
	g.Use(ginzap.RecoveryWithZap(log, true))

	server := &http.Server{Addr: ":8080", Handler: g}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", server.Addr)
			if err != nil {
				return err
			}

			go server.Serve(ln)

			log.Info("Starting HTTP server", zap.String("addr", server.Addr))

			return nil

		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})

	return g
}
