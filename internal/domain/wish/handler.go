package wish

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopkg.in/validator.v2"
)

type Handler interface {
	CreateWish(c *gin.Context) error
}

type handler struct {
	s Service
	g *gin.Engine
	l *zap.Logger
}

func NewHandler(s Service, g *gin.Engine, l *zap.Logger) {
	h := &handler{s: s, g: g, l: l}
	group := g.Group("/wish")
	group.POST("", h.CreateWish)
}

func (h handler) CreateWish(c *gin.Context) {
	var createWishRequestDTO CreateWishRequest

	if err := c.ShouldBindJSON(&createWishRequestDTO); err != nil {
		errorResponse := Error{
			Message: "Invalid request body",
		}
		h.l.Error(err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)

		return
	}

	if err := validator.Validate(createWishRequestDTO); err != nil {
		errorResponse := Error{
			Message: "Request body isn't valid",
		}
		h.l.Error(err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)

		return
	}

	err := h.s.CreateWish(&createWishRequestDTO)

	if err != nil {
		errorResponse := Error{
			Message: "Couldn't create your wish",
		}
		h.l.Error(err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)

		return
	}

	c.JSON(http.StatusOK, CreateWishResponse{Message: "OK"})
}
