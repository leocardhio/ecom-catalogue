package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseCreated(c *gin.Context, data interface{}) {
	// TODO: Remove data from response
	c.JSON(http.StatusCreated, gin.H{"message": "data successfully created", "data": data})
}

func ResponseBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func ResponseInternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}