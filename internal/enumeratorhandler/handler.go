package enumeratorhandler

import (
	"fmt"
	"net/http"
	"strings"
)

type FileEnumerator interface {
	EnumerateFiles(ext string) ([]string, error)
}

type Handler struct {
	Enumerator FileEnumerator
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	filterExt := r.FormValue("ext")
	if filterExt != "" && !strings.HasPrefix(filterExt, ".") {
		filterExt = "." + filterExt
	}

	filesList, err := h.Enumerator.EnumerateFiles(filterExt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	if len(filesList) == 0 {
		fmt.Fprint(w, "No files")
		return
	}
	for _, file := range filesList {
		fmt.Fprintln(w, file)
	}
}
