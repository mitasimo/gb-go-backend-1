package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"mitasimo/gb-go-backend-1/internal/filelistener"
	uploader "mitasimo/gb-go-backend-1/internal/uploadhandler"
)

var (
	uploadDir  string
	serverPort string
)

func main() {

	flag.StringVar(&uploadDir, "ud", "", "path to upload dir")
	flag.StringVar(&serverPort, "sp", "7319", `server's port`)
	flag.Parse()

	if uploadDir == "" {
		log.Fatalln("path to upload dir is not set")
	}

	mux := http.NewServeMux()
	mux.Handle("/upload", &uploader.Handler{
		UploadDir: uploadDir,
	})
	mux.Handle("/download/", http.StripPrefix("/download/", http.FileServer(http.Dir(uploadDir))))
	mux.Handle("/list", &filelistener.Handler{
		UploadDir: uploadDir,
	})

	srv := &http.Server{
		Addr:         ":" + serverPort,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Panic(err)
	}
}
