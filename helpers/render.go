package helpers

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type MetaData struct {
	Title       string
	Description string
	Keywords    string
	Author      string
	Canonical   string
	OpenGraph   struct {
		Title       string
		Description string
		Image       string
		URL         string
		Type        string
	}
	TwitterCard struct {
		Title       string
		Description string
		Image       string
	}
}

func Render() error {
	config, err := GetConfig()

	if err != nil {
		return nil
	}

	if err := os.RemoveAll("build"); err != nil {
		return err
	}

	if err := os.MkdirAll("build/dev", os.ModePerm); err != nil {
		return err
	}

	if err := os.MkdirAll("build/dev/blog", os.ModePerm); err != nil {
		return err
	}

	engine := html.New("./theme", ".html")

	engine.AddFunc(
		"unescape", func(s string) template.HTML {
			return template.HTML(s)
		},
	)

	// home page render for post list
	posts, err := GetMetaData()

	if err != nil {
		return err
	}

	var buf bytes.Buffer

	homeMeta := MetaData{
		Title:       config.Blog.Title,
		Description: config.Blog.Title,
		Keywords:    strings.Join(config.Blog.Keywords, ", "),
		Author:      "John Doe",
		Canonical:   config.Blog.Url,
	}

	homeMeta.OpenGraph.Title = config.Blog.Title
	homeMeta.OpenGraph.Description = config.Blog.Title
	homeMeta.OpenGraph.Image = config.Blog.Url + "/images/" + config.Image.ImageName + "." + config.Image.ImageType
	homeMeta.OpenGraph.URL = config.Blog.Url
	homeMeta.OpenGraph.Type = "website"

	homeMeta.TwitterCard.Title = config.Blog.Title
	homeMeta.TwitterCard.Description = config.Blog.Description
	homeMeta.TwitterCard.Image = config.Blog.Url + "/images/" + config.Image.ImageName + "." + config.Image.ImageType

	if err := engine.Render(&buf, "post", fiber.Map{
		"Posts": posts, "meta": homeMeta,
	}, "layouts/main"); err != nil {
		panic(err)
	}

	if err := os.WriteFile("build/dev/index.html", buf.Bytes(), 0644); err != nil {
		panic(err)
	}

	// searching page for search posts

	type searchingMeta struct {
		Title string
		Slug  string
	}

	var buffaaa bytes.Buffer

	var searchingMetas []searchingMeta

	for _, post := range posts {
		searchingMetas = append(searchingMetas, searchingMeta{
			Title: post.Title,
			Slug:  post.Slug,
		})
	}

	if err := engine.Render(&buffaaa, "searching", fiber.Map{
		"Posts": searchingMetas, "meta": homeMeta,
	}, "layouts/main"); err != nil {
		panic(err)
	}

	if err := os.WriteFile("build/dev/searching.html", buffaaa.Bytes(), 0644); err != nil {
		panic(err)
	}

	// blog page render for single post
	posts, errr := GetMetaData()

	if errr != nil {
		return err
	}

	for _, post := range posts {
		data, err := GetBlogOne(post.Slug)

		if err != nil {
			return err
		}

		var postImage = post.CoverImage

		if config.Image.ImagePathAuto {
			postImage = "/images/" + post.Slug + "." + config.Image.ImageType
		}

		postMeta := MetaData{
			Title:       post.Title,
			Description: post.Description,
			Keywords:    strings.Join(post.Tags, ", "),
			Author:      "John Doe",
			Canonical:   config.Blog.Url + "/blog/" + post.Slug,
		}

		postMeta.OpenGraph.Title = post.Title
		postMeta.OpenGraph.Description = post.Description
		postMeta.OpenGraph.Image = config.Blog.Url + postImage
		postMeta.OpenGraph.URL = config.Blog.Url + "/blog/" + post.Slug
		postMeta.OpenGraph.Type = "article"

		postMeta.TwitterCard.Title = post.Title
		postMeta.TwitterCard.Description = post.Description
		postMeta.TwitterCard.Image = config.Blog.Url + postImage

		var buf bytes.Buffer
		if err := engine.Render(&buf, "blog", fiber.Map{"html": data, "meta": postMeta}, "layouts/main"); err != nil {
			panic(err)
		}
		if err := os.WriteFile("build/dev/blog/"+post.Slug+".html", buf.Bytes(), 0644); err != nil {
			panic(err)
		}
	}

	err = os.CopyFS("build/dev/images", os.DirFS("data/images"))

	if err != nil {
		panic(fmt.Sprintf("Failed to copy images: %v", err))
	}

	err = os.CopyFS("build/dev/assets", os.DirFS("theme/assets"))

	if err != nil {
		panic(fmt.Sprintf("Failed to copy assets: %v", err))
	}

	return nil

}
