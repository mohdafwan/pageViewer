package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Get the port from the environment or set a default port
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Serve the PDF file
	http.HandleFunc("/ResumeAfwan", func(w http.ResponseWriter, r *http.Request) {
		// Replace with your PDF file path
		pdfPath := filepath.Join(".", "Afwan_Shaikh_-_Resume.pdf")
		http.ServeFile(w, r, pdfPath)
	})

	// Serve the HTML file to display the PDF
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Replace with your HTML file path
		htmlPath := filepath.Join(".", "index.html")
		http.ServeFile(w, r, htmlPath)
	})

	// Start the server
	log.Printf("Server is running on http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
