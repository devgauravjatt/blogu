package helpers

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

type NameLogo struct {
	First string `yaml:"first"`
	Last  string `yaml:"last"`
}

type SettingsBlog struct {
	Title       string   `yaml:"title"`
	NameLogo    NameLogo `yaml:"name-logo"`
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

type SettingsPosts struct {
	PostsOnFeed int `yaml:"posts-on-feed"`
}

type SocialLinks struct {
	Twitter   string `yaml:"twitter"`
	Github    string `yaml:"github"`
	Linkedin  string `yaml:"linkedin"`
	Instagram string `yaml:"instagram"`
}

type Config struct {
	Blog        SettingsBlog  `yaml:"settings-blog"`
	Image       SettingsImage `yaml:"settings-image"`
	Posts       SettingsPosts `yaml:"settings-posts"`
	SocialLinks SocialLinks   `yaml:"social-links"`
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
