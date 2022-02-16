package main

import (
	"fmt"
	"log"
	"os"

	"github.com/peakier/gen-gql-model/internal/generator"
	"github.com/urfave/cli/v2"
)

func main() {

	cfgFileFlag := &cli.StringFlag{
		Name:     "config-file",
		Usage:    "Path to .yaml config file. If config file is used, the program will ignore other parameters.",
		Aliases:  []string{"c"},
		Required: false,
	}

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "urlmode",
				Usage: "Generate gql model",
				Flags: []cli.Flag{
					cfgFileFlag,
				},
				Action: func(c *cli.Context) error {
					cfgFile := c.String("config-file")
					fmt.Println(cfgFile)
					return generator.Generate(cfgFile, false)
				},
			},
			{
				Name:  "jsonmode",
				Usage: "Generate gql model",
				Flags: []cli.Flag{
					cfgFileFlag,
				},
				Action: func(c *cli.Context) error {
					cfgFile := c.String("config-file")
					fmt.Println(cfgFile)
					return generator.Generate(cfgFile, false)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
