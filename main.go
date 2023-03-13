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
				w.Header().Set("Content-Type", "text/html")
				fmt.Fprintf(w, `<!DOCTYPE html>
                    <html>
                        <head>
                            <meta charset="UTF-8">
                            <title>Video Player</title>
                            <link href="https://vjs.zencdn.net/7.14.3/video-js.css" rel="stylesheet" />
                            <script src="https://vjs.zencdn.net/7.14.3/video.js"></script>
                        </head>
                        <body>
                            <video id="my-video" class="video-js vjs-default-skin" controls preload="auto" width="640" height="264"
                                data-setup='{}'>
                                <source src="%s" type="application/x-mpegURL">
                            </video>
                        </body>
                    </html>`, videoUrl)
			} else {
				fmt.Fprintf(w, "Unsupported video format")
			}
		} else {
			fmt.Fprintf(w, "Video URL not found in request")
		}
	})
	fmt.Println("Server listening on port 9999")
	http.ListenAndServe(":9999", nil)
}
