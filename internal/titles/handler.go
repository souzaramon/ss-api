package titles

import (
	"errors"
	"net/http"

	"ss-api/pkg/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TitlesHandler struct {
	log              *zap.Logger
	titlesRepository *TitlesRepository
}

func NewTitlesHandler(log *zap.Logger, titlesRepository *TitlesRepository) *TitlesHandler {
	return &TitlesHandler{log: log, titlesRepository: titlesRepository}
}

// @Id           getAllTitles
// @Summary      List all titles
// @Description  Get a list of all titles
// @Tags         titles
// @Accept       json
// @Produce      json
// @Success      200  {array}   Title
// @Failure      500  {object}  util.ApiError
// @Router       /titles/ [get]
func (h *TitlesHandler) GetAll(c *gin.Context) {
	items, err := h.titlesRepository.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ApiError{Message: "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Id           getTitleById
// @Summary      Get title by ID
// @Description  Retrieve a single title by their ID
// @Tags         titles
// @Accept       json
// @Produce      json
// @Param        id   path      string        true  "Title ID"
// @Success      200  {object}  Title
// @Failure      404  {object}  util.ApiError
// @Failure      500  {object}  util.ApiError
// @Router       /titles/{id} [get]
func (h *TitlesHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	item, err := h.titlesRepository.FindById(id)

	if err != nil {
		if errors.Is(err, util.ErrNotFound) {
			c.JSON(http.StatusNotFound, util.ApiError{Message: "Item not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, util.ApiError{Message: "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Id           createTitle
// @Summary      Create title
// @Description  Create a new title
// @Tags         titles
// @Accept       json
// @Produce      json
// @Param        data  body      CreateTitleDto  true  "Title data"
// @Success      200   {object}  Title
// @Failure      400   {object}  util.ApiError
// @Failure      500   {object}  util.ApiError
// @Router       /titles/ [post]
func (h *TitlesHandler) Create(c *gin.Context) {
	var body CreateTitleDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, util.ApiError{Message: err.Error()})
		return
	}

	item, err := h.titlesRepository.Create(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ApiError{Message: "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// @Id           updateTitleById
// @Summary      Update title
// @Description  Update an existing title by ID
// @Tags         titles
// @Accept       json
// @Produce      json
// @Param        id    path      string           true  "Title ID"
// @Param        data  body      UpdateTitleDto  true  "Title data"
// @Success      200   {object}  Title
// @Failure      400   {object}  util.ApiError
// @Failure      500   {object}  util.ApiError
// @Router       /titles/{id} [put]
func (h *TitlesHandler) UpdateById(c *gin.Context) {
	id := c.Param("id")

	var body UpdateTitleDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, util.ApiError{Message: err.Error()})
		return
	}

	item, err := h.titlesRepository.UpdateById(id, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ApiError{Message: "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// @Id           deleteTitleById
// @Summary      Delete title
// @Description  Delete an existing title by ID
// @Tags         titles
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Title ID"
// @Success      204  {object}  nil
// @Failure      500  {object}  util.ApiError
// @Router       /titles/{id} [delete]
func (h *TitlesHandler) DeleteById(c *gin.Context) {
	id := c.Param("id")

	err := h.titlesRepository.DeleteById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ApiError{Message: "Internal server error"})
		return
	}

	c.Status(http.StatusNoContent)
}
