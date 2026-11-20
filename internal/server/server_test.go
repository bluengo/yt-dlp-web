package server

import (
	"slices"
	"testing"
)

func TestSplitLines(t *testing.T) {
	tests := []struct {
		text string
		want []string
	}{
		{"https://first-url.com", []string{"https://first-url.com"}},
		{"https://first-url.com\nhttps://second-url.com", []string{"https://first-url.com", "https://second-url.com"}},
		{"  https://first-url.com  \r\n\r\nhttps://second-url.com\r\n", []string{"https://first-url.com", "https://second-url.com"}},
		{"", nil},
	}

	for _, tt := range tests {
		if got := splitLines(tt.text); !slices.Equal(got, tt.want) {
			t.Errorf("splitLines(%q) = %q, want %q", tt.text, got, tt.want)
		}
	}
}
