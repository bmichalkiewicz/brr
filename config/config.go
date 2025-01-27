package config

import (
	"brr/facts"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

var Settings *Config

type Config struct {
	Gitlab struct {
		Token string `yaml:"token"`
		Url   string `yaml:"url"`
	} `yaml:"gitlab"`
}

// Create config if they dont exist and read it into Settings
func Init() {

	filepath, err := CreateConfig()

	if err != nil {
		log.Fatal().Msgf("Error creating config: %v", err)
	}

	configFile, err := os.Open(filepath)
	if err != nil {
		log.Fatal().Msgf("Error opening config file: %v", err)
	}
	defer configFile.Close()

	byteValue, _ := io.ReadAll(configFile)
	if err := yaml.Unmarshal(byteValue, &Settings); err != nil {
		log.Fatal().Msgf("Error parsing config file: %v", err)
	}
}

// Ensures ~/.brr/config.yaml is created with the template and returns the filepath
func CreateConfig() (string, error) {

	homeDir := facts.GetHomeDirectory()
	configDir := filepath.Join(homeDir, fmt.Sprintf(".%s", facts.Analyse().GetApplicationName()))

	// Create ~/.brr if it does not exist
	err := os.MkdirAll(configDir, 0755)

	if err != nil {
		return "", err
	}

	// ~/.brr/config.yaml
	configFilePath := filepath.Join(configDir, "config.yaml")

	// check if ~/.brr/config.yaml exists
	if _, err = os.Stat(configFilePath); errors.Is(err, os.ErrNotExist) {

		// if not, create it with the config.yaml
		file, err := os.Create(configFilePath)
		if err != nil {
			return "", err
		}
		defer file.Close()

		encoder := yaml.NewEncoder(file)
		encoder.SetIndent(2)

		if err := encoder.Encode(configTemplate); err != nil {
			return "", fmt.Errorf("while encoding the config.yaml : %w", err)
		}

	}

	return configFilePath, nil

}
