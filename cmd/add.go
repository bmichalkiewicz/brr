package cmd

import (
	"brr/facts"
	"brr/template"
	"context"
	"fmt"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

func Add() *cli.Command {
	return &cli.Command{
		Name:  "add",
		Usage: "Add a repository to the repositories file.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Usage:    "The SSH URL of the repository to add.",
				Required: true,
				Validator: func(url string) error {
					// Validate the SSH URL
					err := template.Validate(url)
					if err != nil {
						return err
					}

					return nil
				},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {

			output := cmd.String("output")
			if output == "" {
				output = filepath.Join(facts.GetHomeDirectory(), ".brr")
			}

			update := cmd.Bool("update")
			url := cmd.String("url")

			repo, err := template.Decode(url)
			if err != nil {
				return fmt.Errorf("problem when decoding: %v", err)
			}

			// Initialize the template engine
			template.NewEngine()
			template.Engine.Path = output

			log.Info().Msgf("Using output path: %s", output)

			template.Engine.Create([]*template.GroupTemplate{repo}, update)
			return nil
		},
	}
}
