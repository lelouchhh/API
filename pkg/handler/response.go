package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Success 	string `json:"Status"`
	Message 	string `json:"message"`
	Data		string `json:"Data"`
}

type statusResponse struct {
	Status		string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{Success: "0", Message: "ERROR", Data: message})
}