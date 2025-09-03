package authors

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type AuthorsHandler struct {
	log               *zap.Logger
	authorsRepository *AuthorsRepository
}

func NewAuthorsHandler(log *zap.Logger, authorsRepository *AuthorsRepository) *AuthorsHandler {
	return &AuthorsHandler{log: log, authorsRepository: authorsRepository}
}

func (h *AuthorsHandler) GetAll(c *gin.Context) {
	items, err := h.authorsRepository.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *AuthorsHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	item, err := h.authorsRepository.FindById(id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *AuthorsHandler) Create(c *gin.Context) {
	var body CreateAuthorDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.authorsRepository.Create(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create author"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (h *AuthorsHandler) UpdateById(c *gin.Context) {
	id := c.Param("id")

	var body CreateAuthorDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.authorsRepository.UpdateById(id, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update author"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (h *AuthorsHandler) DeleteById(c *gin.Context) {
	id := c.Param("id")
	h.authorsRepository.DeleteById(id)

	c.Status(http.StatusNoContent)
}
