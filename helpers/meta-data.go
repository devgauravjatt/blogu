package helpers

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"

	"gopkg.in/yaml.v2"
)

// FrontMatter represents the metadata in the front matter of the Markdown file.
type FrontMatter struct {
	Title       string   `yaml:"title"`
	CoverImage  string   `yaml:"coverImage"`
	Slug        string   `yaml:"slug"`
	Tags        []string `yaml:"tags"`
	Description string   `yaml:"description"`
}

// parseFrontMatter extracts front matter from a Markdown document.
func ParseFrontMatter(md string) (*FrontMatter, error) {
	// Regular expression to match the front matter section
	re := regexp.MustCompile(`^---\s*([\s\S]*?)\s*---`)
	matches := re.FindStringSubmatch(md)

	if len(matches) < 2 {
		return nil, errors.New("post in meta not found")
	}

	// Extract the YAML content
	yamlContent := matches[1]

	// Unmarshal the YAML into the FrontMatter struct
	var frontMatter FrontMatter
	err := yaml.Unmarshal([]byte(yamlContent), &frontMatter)
	if err != nil {
		return nil, err
	}

	return &frontMatter, nil
}

func GetMetaData() ([]FrontMatter, error) {
	// Specify the directory containing the markdown files
	dir := "./data/posts"

	// Read all files in the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, errors.New("data/posts not found")
	}

	// Initialize an empty slice to store the front matter
	var frontMatters []FrontMatter

	// get config
	config, err := GetConfig()

	if err != nil {
		return nil, err
	}

	// Loop through the files and rename them
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".md" {

			content, err := os.ReadFile(dir + "/" + file.Name())

			if err != nil {
				return nil, errors.New("data/posts not found")
			}

			data, err := ParseFrontMatter(string(content))

			if err != nil {

				var rre = "post in meta not found -- " + file.Name()

				return nil, errors.New(rre)
			}
			var postImage = data.CoverImage

			if config.Image.ImagePathAuto {
				postImage = "/images/" + data.Slug + "." + config.Image.ImageType
			}

			data.CoverImage = postImage

			frontMatters = append(frontMatters, *data)

		}
	}

	return frontMatters, nil
}
