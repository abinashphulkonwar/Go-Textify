package main

import (
	"log"
	"os"

	command "github.com/abinashphulkonwar/go-text-extractor/Command"
	"github.com/abinashphulkonwar/go-text-extractor/app"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "text extractor",
		Usage:  "make an explosive entrance",
		Action: app.CMDHandler,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    command.INPUT,
				Value:   "input file path",
				Aliases: []string{"c"},

				Usage:    "input file path (image JPG, PNG, JPGE)",
				Required: true,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
