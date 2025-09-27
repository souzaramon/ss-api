package titles

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewTitlesHandler,
		NewTitlesRepository,
	),
	fx.Invoke(func(e *gin.Engine, h *TitlesHandler) {
		g := e.Group("/v1/titles")

		g.GET("/", h.GetAll)
		g.GET("/:id", h.GetById)
		g.POST("/", h.Create)
		g.PUT("/:id", h.UpdateById)
		g.DELETE("/:id", h.DeleteById)
	}),
)
