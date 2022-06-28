package handler

import (
	"AGZ/pkg/structures"
	"github.com/gin-gonic/gin"
	"net/http"
)
func (h *Handler) CreateNews(c *gin.Context) {
	var new structures.New
	new.Title = c.Query("title")
	new.Description = c.Query("description")
	new.Picture = c.Query("picture")
	new.TypeTags =  c.QueryArray("picture")
	new.Tags = c.Query("tags")

	err := h.services.Procedures.CreateNews(new)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": "1",
		"message": "OK",
		"data": "",
	})
}
func (h *Handler) GetEntity(c *gin.Context) {
	var entity[] structures.Entity
	entity, err := h.services.Procedures.GetEntity()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": "1",
		"message": "OK",
		"data": entity,
	})
}