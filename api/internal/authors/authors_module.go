package authors

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewAuthorsHandler,
	),
	fx.Invoke(func(e *gin.Engine, h *AuthorsHandler) {
		g := e.Group("/authors")

		g.GET("/", h.GetAll)
	}),
)
