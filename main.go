package main

import (
	"log"
	"os"

	command "github.com/abinashphulkonwar/Textify/Command"
	"github.com/abinashphulkonwar/Textify/app"
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
				Aliases: []string{"i"},
				Usage:   "input file path (image JPG, PNG, JPGE)",
			},
			&cli.StringFlag{
				Name:        command.OUTPUT,
				Value:       "output",
				Aliases:     []string{"o"},
				Usage:       "input file will be a .txt file with ***new page*** separated the pages",
				DefaultText: "output",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
