package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	Limit = 50
)

type Filters map[string]any

type Pagination[T any] struct {
	Data        []T  `json:"data"`
	Limit       uint `json:"limit"`
	TotalCount  uint `json:"total_count"`
	Pages       uint `json:"pages"`
	CurrentPage uint `json:"current_page"`
	NextPage    uint `json:"next_page"`
}

type Entity interface {
	ID() uuid.UUID

	SetID(uuid.UUID)
}

type CrudRepository[T Entity] interface {
	Create(*T) (*T, error)

	Update(*T) (*T, error)

	Delete(uuid.UUID) error

	All() (*T, error)

	FindById(uuid.UUID) (*T, error)

	Where(Filters) (*T, error)

	WherePaginated(Filters) (*Pagination[T], error)
}

type CrudHandler[T Entity] struct {
	Repository CrudRepository[T]
}

func (h *CrudHandler[T]) Find(c *gin.Context) {
	id, err := uuid.FromBytes([]byte(c.Param("id")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	entity, err := h.Repository.FindById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, entity)
}

func (h *CrudHandler[T]) All(c *gin.Context) {
	entities, err := h.Repository.WherePaginated(nil)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, entities)
}

func (h *CrudHandler[T]) Create(c *gin.Context) {
	var body T

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	newEntry, err := h.Repository.Create(&body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, newEntry)
}

func (h *CrudHandler[T]) Update(c *gin.Context) {
	id, err := uuid.FromBytes([]byte(c.Param("id")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var body T

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	body.SetID(id)

	updatedEntry, err := h.Repository.Update(&body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, updatedEntry)
}

func (h *CrudHandler[T]) Delete(c *gin.Context) {
	id, err := uuid.FromBytes([]byte(c.Param("id")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	if err := h.Repository.Delete(id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.Status(http.StatusNoContent)
}
