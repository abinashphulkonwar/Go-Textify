package handlers

import (
	"errors"
	"net/http"
	"os"
	"strings"

	errorhandlers "github.com/abinashphulkonwar/Textify/error-handler"
)

func CheckFileTypeHandler(path string) {
	file, err := os.Open(path)
	errorhandlers.HandleError(err)
	defer file.Close()

	stat, err := file.Stat()
	errorhandlers.HandleError(err)
	if stat.IsDir() {
		errorhandlers.HandleError(errors.New("not a file"))
	}
	buf := make([]byte, 1024)

	_, err = file.Read(buf)

	errorhandlers.HandleError(err)

	contentType := http.DetectContentType(buf)

	println("file type: ", contentType)
	if !strings.Contains(contentType, "image/") {
		errorhandlers.HandleError(errors.New("not a valid file"))
	}

}
