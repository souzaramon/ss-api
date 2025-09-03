package authors

import (
	"api/pkg/util"
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

// GetAll godoc
// @Summary      List all authors
// @Description  Get a list of all authors
// @Tags         authors
// @Accept       json
// @Produce      json
// @Success      200  {array}   Author
// @Failure      500  {object}  util.ApiError
// @Router       /authors/ [get]
func (h *AuthorsHandler) GetAll(c *gin.Context) {
	items, err := h.authorsRepository.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ApiError{Message: "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, items)
}

// GetById godoc
// @Summary      Get author by ID
// @Description  Retrieve a single author by their ID
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param        id   path      string        true  "Author ID"
// @Success      200  {object}  Author
// @Failure      404  {object}  util.ApiError
// @Failure      500  {object}  util.ApiError
// @Router       /authors/{id} [get]
func (h *AuthorsHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	item, err := h.authorsRepository.FindById(id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, util.ApiError{Message: "Item not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, util.ApiError{Message: "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// Create godoc
// @Summary      Create author
// @Description  Create a new author
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param        data  body      CreateAuthorDto  true  "Author data"
// @Success      200   {object}  Author
// @Failure      400   {object}  util.ApiError
// @Failure      500   {object}  util.ApiError
// @Router       /authors/ [post]
func (h *AuthorsHandler) Create(c *gin.Context) {
	var body CreateAuthorDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, util.ApiError{Message: err.Error()})
		return
	}

	item, err := h.authorsRepository.Create(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ApiError{Message: "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// UpdateById godoc
// @Summary      Update author
// @Description  Update an existing author by ID
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param        id    path      string           true  "Author ID"
// @Param        data  body      UpdateAuthorDto  true  "Author data"
// @Success      200   {object}  Author
// @Failure      400   {object}  util.ApiError
// @Failure      500   {object}  util.ApiError
// @Router       /authors/{id} [put]
func (h *AuthorsHandler) UpdateById(c *gin.Context) {
	id := c.Param("id")

	var body UpdateAuthorDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, util.ApiError{Message: err.Error()})
		return
	}

	item, err := h.authorsRepository.UpdateById(id, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ApiError{Message: "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// DeleteById godoc
// @Summary      Delete author
// @Description  Delete an existing author by ID
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Author ID"
// @Success      204  {object}  nil
// @Failure      500  {object}  util.ApiError
// @Router       /authors/{id} [delete]
func (h *AuthorsHandler) DeleteById(c *gin.Context) {
	id := c.Param("id")

	err := h.authorsRepository.DeleteById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ApiError{Message: "Internal server error"})
		return
	}

	c.Status(http.StatusNoContent)
}
