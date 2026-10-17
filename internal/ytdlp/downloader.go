package ytdlp

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

type Downloader struct{}

func NewDownloader() *Downloader {
	return &Downloader{}
}

func (d *Downloader) Download(url string, outputDir string) error {
	if !validateURL(url) {
		return errors.New("Invalid URL: " + url)
	}

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err != nil {
			return err
		}
	}

	log.Printf("Downloading %s to %s... ▰▰▰▱\n", url, outputDir)
	err := download(url, outputDir)
	if err != nil {
		return err
	} else {
		log.Println("Download completed successfully! ✅")
	}

	return nil
}

func download(url string, outputDir string) error {
	log.Println("Running yt-dlp command...")
	cmd := exec.Command(
		"yt-dlp",
		"--no-playlist",
		"--restrict-filenames",
		"-x",
		"--audio-format", "mp3",
		"-o", outputDir+"/%(title)s.%(ext)s",
		url,
	)
	return cmd.Run()
}
