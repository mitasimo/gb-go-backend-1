package uploadhandler

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	partWriter, err := writer.CreateFormFile("file", "file.io")
	if err != nil {
		t.Fatalf("%v", err)
	}
	_, err = partWriter.Write([]byte("some data"))
	if err != nil {
		t.Fatalf("%v", err)
	}
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload", buf)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	h := &Handler{
		Saver: MockSaver{},
	}
	h.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	got := string(data)
	const must = "http://example.com/download/file.io\n"
	if got != must {
		t.Errorf("\ngot  %v\nmust %v", got, must)
	}

}

type MockSaver struct{}

func (s MockSaver) SaveFile(string, []byte) error {
	return nil
}
