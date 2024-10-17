package helpers

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

type SettingsBlog struct {
	Title       string   `yaml:"title"`
	Url         string   `yaml:"url"`
	Description string   `yaml:"description"`
	Keywords    []string `yaml:"keywords"`
	Author      string   `yaml:"author"`
	Language    string   `yaml:"language"`
}

type SettingsImage struct {
	ImagePathAuto bool   `yaml:"image-path-auto"`
	ImageType     string `yaml:"image-type"`
	ImageName     string `yaml:"image-name"`
}

type Config struct {
	Blog  SettingsBlog  `yaml:"settings-blog"`
	Image SettingsImage `yaml:"settings-image"`
}

func GetConfig() (Config, error) {

	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		return Config{}, errors.New("config.yaml not found - check config.yaml")
	}
	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return Config{}, errors.New("config.yaml bat format error")
	}
	return config, nil

}
