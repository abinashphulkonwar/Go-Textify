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

	if input_file_path == "" {
		panic("file is emply")
	}

	//	url := handlers.Upload(input_file_path)
	data, err := os.Open(input_file_path)
	errorhandlers.HandleError(err)
	defer data.Close()
	handlers.CheckFileTypeHandler(data)
	text := handlers.ExtractText(data)
	//handlers.OutputJson(text, input_file_path)

	handlers.OutputTextFile(text, output_file_path)
	return nil
}
