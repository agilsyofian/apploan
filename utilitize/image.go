package utilitize

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

type Base64 struct {
	Text string
}

func NewBase64(text string) *Base64 {
	return &Base64{
		Text: text,
	}
}

func (b *Base64) CheckMimeType() (valid bool, mimeType string) {
	suffix := ";base64"
	str := strings.Split(b.Text, suffix)
	if len(str) < 2 {
		return false, ""
	}
	mime := str[0][5:]
	switch mime {
	case "image/png":
		return true, mime
	case "image/jpeg":
		return true, mime
	case "image/jpg":
		return true, mime
	}

	return false, ""
}

func (b *Base64) StoreBase64ToImage(mime string, path string, filename string) (string, error) {

	trimmed := strings.TrimPrefix(b.Text, "data:"+mime+";base64,")

	data, err := base64.StdEncoding.DecodeString(trimmed)
	if err != nil {
		err := fmt.Errorf("error decoding base64 string: %s", err)
		return "", err
	}

	// currentWorkDirectory, _ := os.Getwd()

	info, err := os.Stat(path)
	if err != nil || !info.IsDir() {
		os.MkdirAll(path, 0755)
	}

	ext := strings.Split(mime, "/")
	truePath := path + "/" + filename + "." + ext[1]

	file, err := os.Create(truePath)
	if err != nil {
		err := fmt.Errorf("error creating file: %s", err)
		return "", err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		err := fmt.Errorf("error writing file: %s", err)
		return "", err
	}

	return truePath, nil
}
