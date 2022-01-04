package filelistener

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Handler struct {
	UploadDir string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	filterExt := r.FormValue("ext")
	if filterExt != "" && !strings.HasPrefix(filterExt, ".") {
		filterExt = "." + filterExt
	}

	entries, err := os.ReadDir(h.UploadDir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	cnt := 0
	for _, entry := range entries {
		fileName := entry.Name()
		fileExt := filepath.Ext(fileName)
		if filterExt != "" && fileExt != filterExt {
			continue // расширение задано и не соответвтвует расшширению файла
		}
		info, err := entry.Info()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			continue
		}
		fmt.Fprintf(w, "Name: %s\tExt: %s\tsize: %d bytes\n", fileName, fileExt, info.Size())
		cnt++
	}
	if cnt == 0 {
		fmt.Fprint(w, "No files")
	}
}
