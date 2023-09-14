package app

import (
	"os"

	command "github.com/abinashphulkonwar/Textify/Command"
	"github.com/abinashphulkonwar/Textify/app/handlers"
	errorhandlers "github.com/abinashphulkonwar/Textify/error-handler"
	"github.com/urfave/cli/v2"
)

func CMDHandler(ctx *cli.Context) error {
	if ctx.NArg() < 0 {
		return nil
	}
	input_file_path := ctx.String(command.INPUT)
	output_file_path := ctx.String(command.OUTPUT)
	isUpload := ctx.String(command.IS_UPLOAD)
	if input_file_path == "" {
		panic("file is emply")
	}
	handlers.CheckFileTypeHandler(input_file_path)

	data, err := os.Open(input_file_path)
	errorhandlers.HandleError(err)
	defer data.Close()

	url := ""
	isUploadState := false
	if isUpload == "true" {
		isUploadState = true
		url = handlers.Upload(input_file_path)
	}

	text := handlers.ExtractText(data, url, isUploadState)
	//handlers.OutputJson(text, input_file_path)

	handlers.OutputTextFile(text, output_file_path)
	return nil
}
