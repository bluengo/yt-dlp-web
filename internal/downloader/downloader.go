package downloader

import (
	"errors"
	"log"
	"os"

	"github.com/bluengo/yt-dlp-web/internal/command"
	"github.com/bluengo/yt-dlp-web/internal/url"
)

type Downloader struct{}

func NewDownloader() *Downloader {
	return &Downloader{}
}

func (d *Downloader) Download(link string, outputDir string) error {
	if !url.ValidateURL(link) {
		return errors.New("Invalid URL: " + link)
	}

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err != nil {
			return err
		}
	}

	log.Printf("Downloading %s to %s... ▰▰▰▱\n", link, outputDir)
	err := command.RunCommand(link, outputDir)
	if err != nil {
		return err
	} else {
		log.Println("Download completed successfully! ✅")
	}

	return nil
}
