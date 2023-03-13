package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/play-video", func(w http.ResponseWriter, r *http.Request) {
		videoUrl := r.URL.Query().Get("url")
		if videoUrl != "" {
			fmt.Printf("Playing video: %s\n", videoUrl)
			if strings.HasSuffix(videoUrl, ".m3u8") {
				w.Header().Set("Content-Type", "application/vnd.apple.mpegurl")
			} else {
				fmt.Fprintf(w, "Unsupported video format")
				return
			}
			fmt.Fprintf(w, "<video src=\"%s\" controls autoplay></video>", videoUrl)
		} else {
			fmt.Fprintf(w, "Video URL not found in request")
		}
	})
	fmt.Println("Server listening on port 9999")
	http.ListenAndServe(":9999", nil)
}
