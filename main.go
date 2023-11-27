package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Health Checker",
		Usage: "A CLI tool to check the health of your services",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "Domain to check",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port to check",
				Value:    "80",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			port := ctx.String("port")
			status := Check(ctx.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}
	error := app.Run(os.Args)
	if error != nil {
		log.Fatal(error)
	}
}
