package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"nama":    "Anton Wijaya",
		"kelas":   "XII",
		"jurusan": "IPA",
	})
}

func ShowParams(c *gin.Context) {
	nama := c.Param("nama")
	c.JSON(http.StatusOK, gin.H{
		"nama": nama,
	})
}

func QueryApi(c *gin.Context) {
	nama := c.Query("nama")
	c.JSON(http.StatusOK, gin.H{
		"nama": nama,
	})
}
