package handlers

import (
	"fmt"
	"io"
	"net/http"
)

// userVideoListHandler handles the logic for listing available videos for users
func userVideoListHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch the list of available videos from the database
	// You'll need to replace this with your database logic

	videos := []string{"video1.mp4", "video2.mp4", "video3.mp4"}

	// Render the list of videos (this could be an HTML template or another method)
	fmt.Fprintf(w, "Available Videos:\n")
	for _, video := range videos {
		fmt.Fprintf(w, "- %s\n", video)
	}
}

// userDownloadHandler handles the logic for allowing users to download a specific video
func userDownloadHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the videoID from the URL path
	videoID := r.URL.Path[len("/user/download/"):]

	// Validate videoID (you may want to check against a list of valid video IDs or use a database)

	// For demonstration purposes, let's assume the video file is in the storage directory
	videoFilePath := fmt.Sprintf("./storage/videos/%s", videoID)

	// Open the video file
	videoFile, err := http.ServeFile(w, r, videoFilePath)
	if err != nil {
		// Handle the error (e.g., log it, return an error response)
		http.Error(w, "Error serving video file", http.StatusInternalServerError)
		return
	}

	defer videoFile.Close()

	// Set the appropriate Content-Type header for video files
	w.Header().Set("Content-Type", "video/mp4")

	// Stream the video file to the response writer
	_, err = io.Copy(w, videoFile)
	if err != nil {
		// Handle the error (e.g., log it, return an error response)
		http.Error(w, "Error streaming video file", http.StatusInternalServerError)
		return
	}
}
