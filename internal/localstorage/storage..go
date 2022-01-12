package localstorage

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Storage struct {
	Dir string
}

func (s Storage) SaveFile(fileName string, fileData []byte) error {
	filePath := s.Dir + "/" + fileName
	err := ioutil.WriteFile(filePath, fileData, 0777)
	if err != nil {
		return err
	}
	return nil
}

func (s Storage) EnumerateFiles(filterExt string) ([]string, error) {
	files := make([]string, 0)

	entries, err := os.ReadDir(s.Dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		fileName := entry.Name()
		fileExt := filepath.Ext(fileName)
		if filterExt != "" && fileExt != filterExt {
			continue // расширение задано и не соответвтвует расшширению файла
		}
		info, err := entry.Info()
		if err != nil {
			continue
		}
		files = append(files, fmt.Sprintf("Name: %s\tExt: %s\tsize: %d bytes", fileName, fileExt, info.Size()))
	}
	return files, nil
}
