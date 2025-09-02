package authors

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorsHandler struct {
	log *zap.Logger
}

func NewAuthorsHandler(log *zap.Logger) *AuthorsHandler {
	return &AuthorsHandler{log: log}
}

func (h *AuthorsHandler) GetAll(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{"message": "hello world"},
	)
}
