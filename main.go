package main

import (
	"log"
//	"os"

	"github.com/bluengo/yt-dlp-web/internal/server"
//	"github.com/bluengo/yt-dlp-web/internal/ytdlp"
)

const (
	exampleURL = "https://www.youtube.com/watch?v=2wVTBOTsqmo"
	port = 8080
)

func main() {
//	var downloadsDir string = os.Getenv("HOME") + "/Downloads"

	log.Println("Starting yt-dlp-web...")
//	downloader := ytdlp.NewDownloader()
//	downloader.Download(exampleURL, downloadsDir)

	s := server.NewServer(port)
	s.Start()
}
