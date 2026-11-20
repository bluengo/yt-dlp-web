package server

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//go:embed assets
var assets embed.FS

type Server struct {
	port int
}

func NewServer(port int) *Server {
	return &Server{
		port: port,
	}
}

//func StartServer(port int) {
func (s *Server) Start() {
	portStr := ":" + strconv.Itoa(s.port)

	assetsFS, err := fs.Sub(assets, "assets") // remove root folder
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.FS(assetsFS))
	http.Handle("/", fs)
	http.HandleFunc("/download", downloadHandler)

	log.Printf("Starting server on %s...\n", portStr)
	
	err = http.ListenAndServe(portStr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	urls := splitLines(r.FormValue("youtube-urls"))

	// call downloader placeholder

	log.Printf("URLs: %v\n", urls)
//	w.WriteHeader(http.StatusOK)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// splitLines returns the non-empty lines of text, without surrounding spaces
func splitLines(text string) []string {
	var lines []string

	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)

		if line != "" {
			lines = append(lines, line)
		}
	}

	return lines
}

