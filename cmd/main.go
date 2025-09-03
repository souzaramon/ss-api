package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"ss-api/internal/authors"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// @title       SS API
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func main() {
	fx.New(
		fx.Provide(
			NewGinApp,
			NewPgxConn,
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

func NewPgxConn(lc fx.Lifecycle) (*pgx.Conn, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:123@db:5432/test")

	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			conn.Close(ctx)
			return nil
		},
	})

	return conn, nil
}
