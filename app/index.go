package app

import (
	command "github.com/abinashphulkonwar/go-text-extractor/Command"
	"github.com/abinashphulkonwar/go-text-extractor/app/handlers"
	"github.com/urfave/cli/v2"
)

func CMDHandler(ctx *cli.Context) error {
	if ctx.NArg() < 0 {
		return nil
	}
	input_file_path := ctx.String(command.INPUT)
	if input_file_path == "" {
		panic("file is emply")
	}
	url := handlers.Upload(input_file_path)
	println(url)
	return nil
}
