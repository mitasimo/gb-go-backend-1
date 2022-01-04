package uploadhandler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type FileSaver interface {
	SaveFile(string, []byte) error
}

type Handler struct {
	Saver FileSaver
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}

	err = h.Saver.SaveFile(header.Filename, fileData)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "http://"+r.Host+"/download/"+header.Filename)
}
