package url

import "testing"

func TestValidateURL(t *testing.T) {
	tests := []struct {
		url  string
		want bool
	}{
		{"https://www.youtube.com/watch?v=dQw4w9WgXcQ", true},
		{"http://www.youtube.com/watch?v=dQw4w9WgXcQ", true},
		{"", false},
		{"https://non-youtube-url.com", false},
	}

	for _, tt := range tests {
		if got := ValidateURL(tt.url); got != tt.want {
			t.Errorf("ValidateURL(%q) = %v, want %v", tt.url, got, tt.want)
		}
	}
}
