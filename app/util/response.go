package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseCreated(c *gin.Context, data interface{}) {
	// TODO: Refine data returned from response
	c.JSON(http.StatusCreated, gin.H{"message": "data successfully created", "data": data})
}

func ResponseUpdated(c *gin.Context, data interface{}) {
	// TODO: Refine data returned from response
	c.JSON(http.StatusOK, gin.H{"message": "data successfully updated", "data": data})
}

func ResponseDeleted(c *gin.Context, data interface{}) {
	// TODO: Refine data returned from response
	c.JSON(http.StatusOK, gin.H{"message": "data successfully deleted", "data": data})
}

func ResponseBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func ResponseInternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
