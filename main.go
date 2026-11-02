package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
)

// Embed the app assets so they are available in Vercel's deployed Go backend.
//
//go:embed static templates banners
var appFiles embed.FS

func main() {
	staticFiles, err := fs.Sub(appFiles, "static")
	if err != nil {
		panic(err)
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFiles))))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	http.HandleFunc("/download", downloadHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server running at http://localhost:%s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
