package upload

import (
	"encoding/base64"
	"errors"
	"io"
	"os"
	"strings"
)

// SaveBase64 保存base64编码过的内容
func SaveBase64(data string, path string) error {
	idx := strings.Index(data, ";base64,")
	if idx < 0 {
		return errors.New("base64 error")
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[idx+8:]))
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, reader)
	f.Close()
	return err
}
