package enumeratorhandler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	files := make(map[string][]struct {
		name string
		size int
	})
	files["png"] = []struct {
		name string
		size int
	}{
		{name: "image1.png", size: 34156},
		{name: "image2.png", size: 26156},
	}
	files["jpg"] = []struct {
		name string
		size int
	}{
		{name: "image1.jpg", size: 34156},
		{name: "image2.jpg", size: 26156},
	}

	h := Handler{
		Enumerator: MockEnumerator{files: files},
	}

	const pngOut = "Name: image1.png\tExt: png\tsize: 34156 bytes\nName: image2.png\tExt: png\tsize: 26156 bytes\n"

	req := httptest.NewRequest(http.MethodGet, "/list?ext=png", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != pngOut {
		t.Errorf("got %v, must %v", string(data), pngOut)
	}
}

type MockEnumerator struct {
	files map[string][]struct {
		name string
		size int
	}
}

func (e MockEnumerator) EnumerateFiles(ext string) ([]string, error) {
	result := make([]string, 0)
	if ext == "" || ext == ".png" {
		pngFiles := e.files["png"]
		for _, file := range pngFiles {
			result = append(result, fmt.Sprintf("Name: %s\tExt: %s\tsize: %d bytes", file.name, "png", file.size))
		}
	}
	if ext == "" || ext == ".jpg" {
		pngFiles := e.files["jpg"]
		for _, file := range pngFiles {
			result = append(result, fmt.Sprintf("Name: %s\tExt: %s\tsize: %d bytes", file.name, "jpg", file.size))
		}
	}
	return result, nil
}
