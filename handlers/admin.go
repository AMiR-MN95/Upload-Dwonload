package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func AdminInit(router *gin.Engine) {
	router.GET("/admin/upload", adminUploadPage)
	router.POST("/admin/upload", adminUploadHandler)
}

func adminUploadPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func adminUploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the file to storage
	storagePath := "./storage/files/"
	dst := filepath.Join(storagePath, file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
