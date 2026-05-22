package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

const maxInputLength = 1000

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.Header().Set("Allow", http.MethodGet+", "+http.MethodHead)
		renderError(w, http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		renderError(w, http.StatusNotFound)
		return
	}
	renderTemplate(w, map[string]string{"Text": "", "Banner": "standard"})
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		renderError(w, http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		renderError(w, http.StatusBadRequest)
		return
	}

	text := normalizeNewlines(r.PostForm.Get("text"))
	bannerName := r.PostForm.Get("banner")

	if text == "" || bannerName == "" {
		renderError(w, http.StatusBadRequest)
		return
	}

	if len(text) > maxInputLength {
		renderError(w, http.StatusBadRequest)
		return
	}

	if !isPrintableASCIIText(text) {
		renderError(w, http.StatusBadRequest)
		return
	}

	if !isAllowedBanner(bannerName) {
		renderError(w, http.StatusBadRequest)
		return
	}

	banner, err := loadBannerByName(bannerName)
	if err != nil {
		renderError(w, http.StatusNotFound)
		return
	}

	result := renderASCII(text, banner)

	renderTemplate(w, map[string]string{"Result": result, "Text": text, "Banner": bannerName})
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		renderError(w, http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		renderError(w, http.StatusInternalServerError)
		return
	}

	text := normalizeNewlines(r.PostForm.Get("text"))
	bannerName := r.PostForm.Get("banner")

	if text == "" {
		renderError(w, http.StatusBadRequest)
		return
	}

	if !isAllowedBanner(bannerName) {
		renderError(w, http.StatusNotFound)
		return
	}

	banner, err := loadBannerByName(bannerName)
	if err != nil {
		renderError(w, http.StatusNotFound)
		return
	}

	result := renderASCII(text, banner)
	resultBytes := []byte(result)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Disposition", `attachment; filename="ascii-art.txt"`)
	w.Header().Set("Content-Length", fmt.Sprint(len(resultBytes)))
	w.WriteHeader(http.StatusOK)
	w.Write(resultBytes)
}

func renderTemplate(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func renderError(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles(
		fmt.Sprintf("templates/%d.html", status),
	)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)

	tmpl.Execute(w, nil)
}

func normalizeNewlines(s string) string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.ReplaceAll(s, "\r", "")
	return s
}

func loadBannerByName(bannerName string) (map[rune][]string, error) {
	bannerFile := "banners/" + bannerName + ".txt"
	return loadBanner(bannerFile)
}

func isAllowedBanner(bannerName string) bool {
	switch bannerName {
	case "standard", "shadow", "thinkertoy":
		return true
	default:
		return false
	}
}

func isPrintableASCIIText(s string) bool {
	for _, ch := range s {
		if ch == '\n' {
			continue
		}
		if ch < 32 || ch > 126 {
			return false
		}
	}
	return true
}
