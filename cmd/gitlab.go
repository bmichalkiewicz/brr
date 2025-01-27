package cmd

import (
	"brr/config"
	"brr/facts"
	"brr/gitlab"
	"brr/template"
	"context"
	"fmt"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

func Gitlab() *cli.Command {
	return &cli.Command{
		Name:  "gitlab",
		Usage: "Download GitLab repositories",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "token",
				Aliases: []string{"t"},
				Usage:   "GitLab API token for authentication.",
			},
			&cli.StringSliceFlag{
				Name:     "groups",
				Aliases:  []string{"g"},
				Usage:    "A comma-separated list of GitLab group names to fetch repositories.",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "url",
				Aliases: []string{"u"},
				Usage:   "URL of the GitLab instance.",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			output := cmd.String("output")
			if output == "" {
				output = filepath.Join(facts.GetHomeDirectory(), ".brr")
			}

			update := cmd.Bool("update")
			groups := cmd.StringSlice("groups")

			if config.Settings.Gitlab.Token == "" {
				if cmd.String("token") == "" {
					return fmt.Errorf("required flag \"token\" not set")
				}
			} else {
				log.Info().Msgf("Used token from config")
			}

			// Initialize the template engine
			template.NewEngine()
			template.Engine.Path = output

			log.Info().Msgf("Using output path: %s", output)

			// Initialize GitLab client
			gm, err := gitlab.Init(config.Settings.Gitlab.Token, config.Settings.Gitlab.Url)
			if err != nil {
				log.Error().Err(err).Msgf("Failed to initialize GitLab client with URL: %s", config.Settings.Gitlab.Url)
				return err
			}

			log.Info().Msgf("Fetching repositories for groups: %v", groups)

			projects, err := gm.GetGroupProjects(groups)
			if err != nil {
				log.Error().Err(err).Msg("Failed get groups projects")
				return err
			}

			err = template.Engine.Create(projects, update)
			if err != nil {
				log.Error().Err(err).Msg("Failed to save projects to file")
				return err
			}

			log.Info().Msg("GitLab repository fetch completed successfully.")
			return nil
		},
	}
}
