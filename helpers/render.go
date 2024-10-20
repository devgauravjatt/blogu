package helpers

import (
	"bytes"
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

func HomeRender(config Config, posts []FrontMatter, engine *html.Engine) error {

	var postOneFeed = config.Posts.PostsOnFeed

	if postOneFeed > len(posts) {
		postOneFeed = len(posts)
	} else {
		postOneFeed = config.Posts.PostsOnFeed
	}

	var getPosts = posts[:postOneFeed]

	homeMeta := MetaData{
		Title:       config.Blog.Title,
		Description: config.Blog.Title,
		Keywords:    strings.Join(config.Blog.Keywords, ", "),
		Author:      config.Blog.Author,
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

	var homeBuff bytes.Buffer

	if err := engine.Render(&homeBuff, "post", fiber.Map{"blogData": config.Blog,
		"Posts": getPosts, "SocialLinks": config.SocialLinks, "meta": homeMeta,
	}, "layouts/main"); err != nil {
		return err
	}

	if err := os.WriteFile("build/index.html", homeBuff.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}

func PostsRender(config Config, posts []FrontMatter, engine *html.Engine) error {

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
			Author:      config.Blog.Author,
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

		var blogBuff bytes.Buffer
		if err := engine.Render(&blogBuff, "blog", fiber.Map{"blogData": config.Blog, "html": data, "meta": postMeta}, "layouts/main"); err != nil {
			return err
		}
		if err := os.MkdirAll("build/blog/"+post.Slug, os.ModePerm); err != nil {
			return err
		}
		if err := os.WriteFile("build/blog/"+post.Slug+"/index.html", blogBuff.Bytes(), 0644); err != nil {
			return err
		}
	}

	return nil
}

func TagsRender(config Config, posts []FrontMatter, engine *html.Engine) error {

	homeMeta := MetaData{
		Title:       config.Blog.Title,
		Description: config.Blog.Title,
		Keywords:    strings.Join(config.Blog.Keywords, ", "),
		Author:      config.Blog.Author,
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

	tagSet := make(map[string]struct{})

	for _, post := range posts {
		for _, tag := range post.Tags {
			tagSet[tag] = struct{}{}
		}
	}
	var uniqueTags []string
	for tag := range tagSet {
		uniqueTags = append(uniqueTags, tag)
	}

	// build all tags for folders
	for _, tag := range uniqueTags {
		if err := os.MkdirAll("build/tags/"+tag, os.ModePerm); err != nil {
			return err
		}
	}

	// posts list one line by tag
	for _, mainTag := range uniqueTags {
		var bufTag bytes.Buffer

		var tagPosts []FrontMatter

		for _, post := range posts {
			for _, tag := range post.Tags {
				if tag == mainTag {
					tagPosts = append(tagPosts, post)
				}
			}
		}

		homeMeta.Keywords = strings.Join(tagPosts[0].Tags, ", ")
		homeMeta.Title = mainTag + " - " + config.Blog.Title
		homeMeta.OpenGraph.Title = mainTag + " - " + config.Blog.Title
		homeMeta.TwitterCard.Title = mainTag + " - " + config.Blog.Title
		homeMeta.OpenGraph.URL = config.Blog.Url + "/tags/" + mainTag
		homeMeta.Canonical = config.Blog.Url + "/tags/" + mainTag

		if err := engine.Render(&bufTag, "tag", fiber.Map{"blogData": config.Blog,
			"Posts": tagPosts, "meta": homeMeta,
		}, "layouts/main"); err != nil {
			return err
		}

		if err := os.WriteFile("build/tags/"+mainTag+"/index.html", bufTag.Bytes(), 0644); err != nil {
			return err
		}

	}

	return nil
}

func ErrorPageRender(config Config, engine *html.Engine) error {
	var errorBuff bytes.Buffer

	errorMeta := MetaData{
		Title:       "404 Page Not Found | " + config.Blog.Title,
		Description: "404 Page Not Found | " + config.Blog.Title,
		Keywords:    strings.Join(config.Blog.Keywords, ", "),
		Author:      config.Blog.Author,
		Canonical:   config.Blog.Url,
	}

	errorMeta.OpenGraph.Title = "404 Page Not Found | " + config.Blog.Title
	errorMeta.OpenGraph.Description = config.Blog.Description
	errorMeta.OpenGraph.Image = config.Blog.Url + "/assets/" + config.Image.ImageName + "." + config.Image.ImageType
	errorMeta.OpenGraph.URL = config.Blog.Url
	errorMeta.OpenGraph.Type = "website"

	errorMeta.TwitterCard.Title = "404 Page Not Found | " + config.Blog.Title
	errorMeta.TwitterCard.Description = config.Blog.Description
	errorMeta.TwitterCard.Image = config.Blog.Url + "/assets/" + config.Image.ImageName + "." + config.Image.ImageType

	if err := engine.Render(&errorBuff, "404", fiber.Map{"blogData": config.Blog,
		"meta": errorMeta, "error": "404 Page Not Found", "img": "/assets/" + config.Image.ImageName + "." + config.Image.ImageType,
	}, "layouts/main"); err != nil {
		panic(err)
	}

	if err := os.WriteFile("build/404.html", errorBuff.Bytes(), 0644); err != nil {
		panic(err)
	}
	return nil
}

func SearchPageRender(config Config, posts []FrontMatter, engine *html.Engine) error {
	searchMeta := MetaData{
		Title:       "searching page | " + config.Blog.Title,
		Description: "searching page | " + config.Blog.Title,
		Keywords:    strings.Join(config.Blog.Keywords, ", "),
		Author:      config.Blog.Author,
		Canonical:   config.Blog.Url,
	}

	searchMeta.OpenGraph.Title = "searching page | " + config.Blog.Title
	searchMeta.OpenGraph.Description = config.Blog.Description
	searchMeta.OpenGraph.Image = config.Blog.Url + "/assets/" + config.Image.ImageName + "." + config.Image.ImageType
	searchMeta.OpenGraph.URL = config.Blog.Url
	searchMeta.OpenGraph.Type = "website"

	searchMeta.TwitterCard.Title = "searching page | " + config.Blog.Title
	searchMeta.TwitterCard.Description = config.Blog.Description
	searchMeta.TwitterCard.Image = config.Blog.Url + "/assets/" + config.Image.ImageName + "." + config.Image.ImageType

	type searchingMeta struct {
		Title string
		Slug  string
	}

	var searchBuff bytes.Buffer

	var searchingMetas []searchingMeta

	for _, post := range posts {
		searchingMetas = append(searchingMetas, searchingMeta{
			Title: post.Title,
			Slug:  post.Slug,
		})
	}

	if err := engine.Render(&searchBuff, "searching", fiber.Map{"blogData": config.Blog,
		"Posts": searchingMetas, "meta": searchMeta,
	}, "layouts/main"); err != nil {
		panic(err)
	}

	if err := os.WriteFile("build/searching/index.html", searchBuff.Bytes(), 0644); err != nil {
		panic(err)
	}

	return nil
}
