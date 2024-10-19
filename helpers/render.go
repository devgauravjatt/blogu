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

	var admin = config.Blog.Author

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

	engine := html.New("./theme", ".html")

	engine.AddFunc(
		"unescape", func(s string) template.HTML {
			return template.HTML(s)
		},
	)

	// home page render for post list
	posts, err := GetMetaData()

	var socialLinks = config.SocialLinks
	fmt.Println(socialLinks)

	var postOneFeed = config.Posts.PostsOnFeed

	if postOneFeed > len(posts) {
		postOneFeed = len(posts)
	} else {
		postOneFeed = config.Posts.PostsOnFeed
	}

	var getPosts = posts[:postOneFeed]

	if err != nil {
		return err
	}

	var buf bytes.Buffer

	homeMeta := MetaData{
		Title:       config.Blog.Title,
		Description: config.Blog.Title,
		Keywords:    strings.Join(config.Blog.Keywords, ", "),
		Author:      admin,
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
		"Posts": getPosts, "SocialLinks": socialLinks, "meta": homeMeta,
	}, "layouts/main"); err != nil {
		panic(err)
	}

	if err := os.WriteFile("build/index.html", buf.Bytes(), 0644); err != nil {
		panic(err)
	}

	// folder for all tags

	tagSet := make(map[string]struct{}) // Use a map to store unique tags

	for _, post := range posts {
		for _, tag := range post.Tags {
			tagSet[tag] = struct{}{} // Add each tag to the map (duplicate tags will be ignored)
		}
	}

	// Convert the map keys to a slice
	var uniqueTags []string
	for tag := range tagSet {
		uniqueTags = append(uniqueTags, tag)
	}

	for _, tag := range uniqueTags {
		if err := os.MkdirAll("build/tags/"+tag, os.ModePerm); err != nil {
			return err
		}
	}

	//  post list by tag
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

		if err := engine.Render(&bufTag, "tag", fiber.Map{
			"Posts": tagPosts, "meta": homeMeta,
		}, "layouts/main"); err != nil {
			panic(err)
		}

		if err := os.WriteFile("build/tags/"+mainTag+"/index.html", bufTag.Bytes(), 0644); err != nil {
			panic(err)
		}

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

	if err := os.MkdirAll("build/searching", os.ModePerm); err != nil {
		return err
	}
	if err := os.WriteFile("build/searching/index.html", buffaaa.Bytes(), 0644); err != nil {
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
			Author:      admin,
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
		if err := os.MkdirAll("build/blog/"+post.Slug, os.ModePerm); err != nil {
			return err
		}
		if err := os.WriteFile("build/blog/"+post.Slug+"/index.html", buf.Bytes(), 0644); err != nil {
			panic(err)
		}
	}

	// build 404 page

	var bufy bytes.Buffer

	errorMeta := MetaData{
		Title:       config.Blog.Title + " | 404 Page Not Found",
		Description: config.Blog.Title + " | 404 Page Not Found",
		Keywords:    strings.Join(config.Blog.Keywords, ", "),
		Author:      admin,
		Canonical:   config.Blog.Url,
	}

	errorMeta.OpenGraph.Title = config.Blog.Title
	errorMeta.OpenGraph.Description = config.Blog.Description
	errorMeta.OpenGraph.Image = config.Blog.Url + "/assets/" + config.Image.ImageName + "." + config.Image.ImageType
	errorMeta.OpenGraph.URL = config.Blog.Url
	errorMeta.OpenGraph.Type = "website"

	errorMeta.TwitterCard.Title = config.Blog.Title
	errorMeta.TwitterCard.Description = config.Blog.Description
	errorMeta.TwitterCard.Image = config.Blog.Url + "/assets/" + config.Image.ImageName + "." + config.Image.ImageType

	if err := engine.Render(&bufy, "404", fiber.Map{
		"meta": errorMeta, "error": "404 Page Not Found", "img": "/assets/" + config.Image.ImageName + "." + config.Image.ImageType,
	}, "layouts/main"); err != nil {
		panic(err)
	}

	if err := os.WriteFile("build/404.html", bufy.Bytes(), 0644); err != nil {
		panic(err)
	}

	err = os.CopyFS("build/images", os.DirFS("data/images"))

	if err != nil {
		panic(fmt.Sprintf("Failed to copy images: %v", err))
	}

	err = os.CopyFS("build/assets", os.DirFS("theme/assets"))

	if err != nil {
		panic(fmt.Sprintf("Failed to copy assets: %v", err))
	}

	return nil

}
