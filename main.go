package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/play-video", func(w http.ResponseWriter, r *http.Request) {
		videoUrl := r.URL.Query().Get("url")
		if videoUrl != "" {
			cmd := exec.Command("mpv", videoUrl)
			if err := cmd.Run(); err != nil {
				fmt.Fprintf(w, "Error playing video: %s", err)
			}
			fmt.Fprintf(w, "Video played successfully")
		} else {
			fmt.Fprintf(w, "Video URL not found in request")
		}
	})
	http.ListenAndServe(":9999", nil)
}
