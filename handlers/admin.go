package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// adminUploadHandler handles video uploading for admin
func adminUploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse the form data, including the uploaded file
		err := r.ParseMultipartForm(10 << 20) // 10 MB limit
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		// Get the file from the form data
		file, handler, err := r.FormFile("videoFile")
		if err != nil {
			http.Error(w, "Unable to get file from form", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Create a unique filename for the uploaded file
		uploadDir := "./storage/videos/"
		filename := filepath.Join(uploadDir, handler.Filename)

		// Create the file on the server
		dst, err := os.Create(filename)
		if err != nil {
			http.Error(w, "Unable to create the file on the server", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		// Copy the file content from the uploaded file to the destination file
		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, "Unable to copy the file content", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "File %s uploaded successfully", handler.Filename)
	} else {
		// Render the admin interface or provide a form for video uploading
		// This could be an HTML template or another method based on your frontend technology
		http.ServeFile(w, r, "admin_interface.html")
	}
}

// adminInterfaceHandler renders the admin interface
func adminInterfaceHandler(w http.ResponseWriter, r *http.Request) {
	// Render your admin interface here
	// This could be an HTML template or another method based on your frontend technology
	http.ServeFile(w, r, "admin_interface.html")
}
