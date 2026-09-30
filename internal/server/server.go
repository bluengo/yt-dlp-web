package server

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"strconv"
)

//go:embed assets
var assets embed.FS

func StartServer(port int) {
	portStr = ":" + strconv.Itoa(port)

	assetsFS, err := fs.Sub(assets, "assets") // remove root folder
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.FS(assetsFS))
	http.Handle("/", fs)

	log.Printf("Starting server on %s...\n", portStr)
	
	err := http.ListenAndServe(portStr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
