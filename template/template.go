package template

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"brr/facts"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

var Engine *templateEngine

type templateEngine struct {
	Path     string
	Filename string

	fullPath string
}

func NewEngine() {
	Engine = &templateEngine{
		Path:     facts.GetTemplatePath(),
		Filename: facts.GetTemplateFile(),
	}
}

// LoadTemplate loads the template from the file.
func (e *templateEngine) load() ([]*GroupTemplate, error) {
	template := []*GroupTemplate{}

	templateFile, err := os.Open(e.fullPath)
	if errors.Is(err, os.ErrNotExist) {
		// Return empty template if the file doesn't exist
		return template, nil
	} else if err != nil {
		return nil, fmt.Errorf("error opening template file: %w", err)
	}
	defer templateFile.Close()

	byteValue, err := io.ReadAll(templateFile)
	if err != nil {
		return nil, fmt.Errorf("error reading template file: %w", err)
	}

	if err := yaml.Unmarshal(byteValue, &template); err != nil {
		return nil, fmt.Errorf("error parsing template file: %w", err)
	}

	return template, nil
}

// CreateTemplate creates or updates the template file.
func (e *templateEngine) Create(content []*GroupTemplate, update bool) error {
	e.setFullPath()

	// Ensure ~/.brr/ directory exists
	if _, err := os.Stat(e.Path); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(e.Path, 0755); err != nil {
			return fmt.Errorf("error creating config directory: %w", err)
		}
	}

	// Load existing template
	existingTemplate, err := e.load()
	if err != nil {
		return fmt.Errorf("warning: could not load existing template: %v", err)
	}

	// Create a map for quick lookup of existing groups
	existingGroupMap := make(map[string]int)
	for i, group := range existingTemplate {
		existingGroupMap[group.Name] = i
	}

	// Add or update groups
	for _, group := range content {
		if index, found := existingGroupMap[group.Name]; found {
			// Update existing group with new projects
			if update {
				projectSet := make(map[string]struct{})
				for _, project := range existingTemplate[index].Projects {
					projectSet[project.Url] = struct{}{}
				}
				for _, project := range group.Projects {
					if _, exists := projectSet[project.Url]; !exists {
						existingTemplate[index].Projects = append(existingTemplate[index].Projects, project)
					}
				}
			}
		} else {
			// Add new group if not found
			existingTemplate = append(existingTemplate, group)
		}
	}

	// Open the file for writing
	file, err := os.OpenFile(e.fullPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening template file for writing: %w", err)
	}
	defer file.Close()

	// Sort the existing template alphabetically by group name
	slices.SortFunc(existingTemplate, func(a, b *GroupTemplate) int {
		return strings.Compare(a.Name, b.Name)
	})

	// Write the template to the file
	encoder := yaml.NewEncoder(file)
	encoder.SetIndent(2)

	if err := encoder.Encode(existingTemplate); err != nil {
		return fmt.Errorf("error encoding template content: %w", err)
	}

	log.Info().Msg("Repositories saved")

	return nil
}

// CreateTemplate creates or updates the template file.
func (e *templateEngine) setFullPath() {
	e.fullPath = filepath.Join(e.Path, e.Filename)
}
