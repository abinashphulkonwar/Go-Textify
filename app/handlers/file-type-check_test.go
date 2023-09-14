package handlers_test

import (
	"testing"

	"github.com/abinashphulkonwar/Textify/app/handlers"
)

func TestFileTypeCheck(t *testing.T) {

	handlers.CheckFileTypeHandler("../../20230816_133449.jpg")

}
