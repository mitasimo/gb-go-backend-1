package filelistener

import (
	"fmt"
	"net/http"
	"os"
)

type Handler struct {
	UploadDir string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	entries, err := os.ReadDir(h.UploadDir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, entry := range entries {
		fn := entry.Name()
		info, err := entry.Info()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			continue
		}
		fmt.Fprintf(w, "%s\tsize: %d bytes\n", fn, info.Size())
	}
}
