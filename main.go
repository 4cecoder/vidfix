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
			fmt.Printf("Playing video: %s\n", videoUrl)
			cmd := exec.Command("/usr/bin/mpv", videoUrl)
			if err := cmd.Run(); err != nil {
				fmt.Fprintf(w, "Error playing video: %s", err)
			} else {
				fmt.Fprintf(w, "Video played successfully")
			}
		} else {
			fmt.Fprintf(w, "Video URL not found in request")
		}
	})
	fmt.Println("Server listening on port 9999")
	http.ListenAndServe(":9999", nil)
}
