package authors

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewAuthorsHandler,
		NewAuthorsRepository,
	),
	fx.Invoke(func(e *gin.Engine, h *AuthorsHandler) {
		g := e.Group("/v1/authors")

		g.GET("/", h.GetAll)
		g.GET("/:id", h.GetById)
		g.POST("/", h.Create)
		g.PUT("/:id", h.UpdateById)
		g.DELETE("/:id", h.DeleteById)
	}),
)
