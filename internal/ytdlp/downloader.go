package ytdlp

import (
	"errors"
	"log"
	"os"
	"os/exec"

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
	err := download(link, outputDir)
	if err != nil {
		return err
	} else {
		log.Println("Download completed successfully! ✅")
	}

	return nil
}

func download(link string, outputDir string) error {
	log.Println("Running yt-dlp command...")
	cmd := exec.Command(
		"yt-dlp",
		"--no-playlist",
		"--restrict-filenames",
		"-x",
		"--audio-format", "mp3",
		"-o", outputDir+"/%(title)s.%(ext)s",
		link,
	)
	return cmd.Run()
}
