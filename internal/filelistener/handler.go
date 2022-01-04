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

	ext := r.FormValue("ext")
	if ext != "" && !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}

	entries, err := os.ReadDir(h.UploadDir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cnt := 0
	for _, entry := range entries {
		fn := entry.Name()
		if ext != "" && filepath.Ext(fn) != ext {
			continue // расширение задано и не соответвтвует расшширению файла
		}
		info, err := entry.Info()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			continue
		}
		fmt.Fprintf(w, "%s\tsize: %d bytes\n", fn, info.Size())
		cnt++
	}
	if cnt == 0 {
		fmt.Fprint(w, "No files")
	}
}
