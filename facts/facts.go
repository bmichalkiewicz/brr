package facts

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/mitchellh/go-homedir"
)

const applicationName = "brr"

type Analysis struct {
	homeDirectory string
	configFile    string
	configPath    string
}

func (a *Analysis) GetApplicationName() string {
	return applicationName
}
func (a *Analysis) GetHomeDirectory() string {
	return a.homeDirectory
}

func (a *Analysis) GetTemplateFile() string {
	return a.configFile
}

func (a *Analysis) GetTemplatePath() string {
	return a.configPath
}

func Analyse() *Analysis {
	return &Analysis{
		homeDirectory: GetHomeDirectory(),
		configFile:    GetTemplateFile(),
		configPath:    GetTemplatePath(),
	}
}

func GetDistribution() string {
	dist, err := exec.Command("lsb_release", "-sd").Output()
	if err != nil {
		return ""
	}
	return string(dist)
}

func GetHomeDirectory() string {
	homeDir, err := homedir.Dir()
	if err != nil {
		return ""
	}

	return homeDir
}

func GetTemplateFile() string {
	return "repositories.yaml"
}

func GetTemplatePath() string {
	return fmt.Sprintf(
		"%s/.%s", GetHomeDirectory(), strings.ToLower(applicationName))
}
