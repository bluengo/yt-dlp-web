package command

import (
	"log"
	"os/exec"
)

const command = "yt-dlp"

func RunCommand(link string, outputDir string) error {
	log.Println("Running yt-dlp command...")
	cmd := exec.Command(
		command,
		"--no-playlist",
		"--restrict-filenames",
		"-x",
		"--audio-format", "mp3",
		"-o", outputDir+"/%(title)s.%(ext)s",
		link,
	)
	return cmd.Run()
}
