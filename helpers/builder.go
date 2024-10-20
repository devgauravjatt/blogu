package helpers

import (
	"html/template"
	"os"

	"github.com/gofiber/template/html/v2"
)

func Builder() error {
	config, err := GetConfig()

	if err != nil {
		return err
	}

	posts, err := GetMetaData()

	if err != nil {
		return err
	}

	if err := os.RemoveAll("build"); err != nil {
		return err
	}

	if err := os.MkdirAll("build", os.ModePerm); err != nil {
		return err
	}

	if err := os.MkdirAll("build/blog", os.ModePerm); err != nil {
		return err
	}

	if err := os.MkdirAll("build/tags", os.ModePerm); err != nil {
		return err
	}

	if err := os.MkdirAll("build/searching", os.ModePerm); err != nil {
		return err
	}

	engine := html.New("./theme", ".html")

	engine.AddFunc(
		"unescape", func(s string) template.HTML {
			return template.HTML(s)
		},
	)

	// home page build and render
	if err := HomeRender(config, posts, engine); err != nil {
		return err
	}

	// blog page build and render
	if err := PostsRender(config, posts, engine); err != nil {
		return err
	}

	// tags page build and render
	if err := TagsRender(config, posts, engine); err != nil {
		return err
	}

	// error page build and render
	if err := ErrorPageRender(config, engine); err != nil {
		return err
	}

	// search page build and render
	if err := SearchPageRender(config, posts, engine); err != nil {
		return err
	}

	// copy assets and images
	err = os.CopyFS("build/images", os.DirFS("data/images"))

	if err != nil {
		return err
	}

	err = os.CopyFS("build/assets", os.DirFS("theme/assets"))

	if err != nil {
		return err
	}

	return nil

}
