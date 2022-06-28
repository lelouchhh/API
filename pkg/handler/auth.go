package handler

import (
	"AGZ/pkg/structures"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
func (h *Handler) signIn(c *gin.Context) {
	var input structures.User
	input.Password = c.Query("pass")
	input.Name = c.Query("login")
	input.DeviceId = c.Query("device_id")
	tokens, err := h.services.Authorization.AuthenticationUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(tokens)
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": "1",
		"message": "OK",
		"data": tokens,
	})
}



func (h *Handler) signUp(c *gin.Context) {
	var input structures.User
	input.Name = c.Query("pass")
	input.Password = c.Query("login")

	err := h.services.Authorization.SetReg(input)
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