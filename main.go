package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/play-video", func(w http.ResponseWriter, r *http.Request) {
		videoUrl := r.URL.Query().Get("url")
		if videoUrl != "" {
			fmt.Printf("Playing video: %s\n", videoUrl)
			html := fmt.Sprintf("<video src=\"%s\" controls autoplay></video>", videoUrl)
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprint(w, html)
		} else {
			fmt.Fprintf(w, "Video URL not found in request")
		}
	})
	fmt.Println("Server listening on port 9999")
	http.ListenAndServe(":9999", nil)
}
