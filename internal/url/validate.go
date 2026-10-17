package url

import (
	"log"
	"strings"
)

func ValidateURL(url string) bool {
	log.Println("Validating URL:", url)
	if url == "" {
		log.Println("URL is empty ❌")
		return false
	}
	if !strings.HasPrefix(url, "http://"+ytPath) && !strings.HasPrefix(url, "https://"+ytPath) {
		log.Println("URL is not a valid YouTube URL ❌")
		return false
	}

	log.Println("URL is valid ✔️")
	return true
}
