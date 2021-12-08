package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	// listenAddr := os.Getenv("LISTEN_ADDRESS")
	// addr := listenAddr + `:` + os.Getenv("PORT")

	listenAddr := `localhost`
	addr := listenAddr + `:8080`

	http.HandleFunc("/watch", stream)
	log.Printf("Server started at: %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func stream(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("v")

	if v == `` {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("\nWrong url format: youtube/watch?v=%s", v)
		return
	}

	err := videoFactory(v, w)
	if err != nil {
		log.Printf("stream error :%v", err)
		fmt.Fprintf(w, "stream error: %v", err)
		return
	}

}

func videoFactory(id string, out io.Writer) error {
	url := fmt.Sprintf("https://youtube.com/watch?v=" + id)

	video_pipe_reader, video_pipe_writer := io.Pipe()
	defer video_pipe_reader.Close()

	ytdl := exec.Command("youtube-dl", url, "-o-")

	ytdl.Stdout = video_pipe_writer // PIPE INPUT
	ytdl.Stderr = os.Stderr         // PIPE2

	ffmpeg := exec.Command("ffmpeg", "-i", "/", "-f", "mp3", "-ab", "96000", "-vn", "-")

	ffmpeg.Stdin = video_pipe_reader // PIPE OUTPUT
	ffmpeg.Stdout = out              // PIPE2 INPUT
	ffmpeg.Stderr = os.Stderr

	// starting to fetching video
	go func() {
		if err := ytdl.Run(); err != nil {
			log.Printf("ERROR ON YOUTUBE FETCH %v", err)
		}
	}()

	err := ffmpeg.Run()

	return err // PIPE2 OUTPUT
}
