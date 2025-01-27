package main

import (
	"brr/cmd"
	"brr/config"
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

func main() {
	// Configure logging
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	config.Init()

	// Create a new CLI application
	app := &cli.Command{
		Name:    "brr",
		Usage:   "'brr' is a tool to fetch and process repositories, producing a ready-to-go 'repositories.yaml' file for use with the aww CLI.",
		Version: "1.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "Path where `brr` outputs the repositories.yaml",
				DefaultText: "~/.brr/",
			},
			&cli.BoolFlag{
				Name:    "update",
				Aliases: []string{"p"},
				Usage:   "Updates groups already present in the repository file.",
			},
		},
		Commands: []*cli.Command{
			cmd.Gitlab(),
			cmd.Add(),
		},
	}

	// Run the application
	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal().Err(err).Msg("Application encountered an error")
	}
}
