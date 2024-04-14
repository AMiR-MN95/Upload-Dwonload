package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func UserInit(router *gin.Engine) {
	router.GET("/user/files", userFilesPage)
	router.GET("/user/download/:filename", userDownloadHandler)
}

func userFilesPage(c *gin.Context) {
	// Fetch the list of available files from the storage directory
	files, err := getFilesList("./storage/files/")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{"files": files})
}

func userDownloadHandler(c *gin.Context) {
	filename := c.Param("filename")
	filePath := filepath.Join("./storage/files/", filename)

	// Check if the file exists
	_, err := os.Stat(filePath)
	if err != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"error": "File not found"})
		return
	}

	c.File(filePath)
}

func getFilesList(directory string) ([]string, error) {
	var files []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, info.Name())
		}
		return nil
	})
	return files, err
}
