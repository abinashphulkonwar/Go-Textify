package handlers_test

import (
	"os"
	"testing"

	"github.com/abinashphulkonwar/Textify/app/handlers"
	errorhandlers "github.com/abinashphulkonwar/Textify/error-handler"
)

func TestFileTypeCheck(t *testing.T) {

	file, err := os.Open("../../20230816_133449.jpg")
	errorhandlers.HandleError(err)

	defer file.Close()

	handlers.CheckFileTypeHandler(file)

}
