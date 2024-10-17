package helpers

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func removeMetadata(inputFilePath string) (string, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content []string
	inMetadata := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "---" {
			inMetadata = !inMetadata
			continue
		}
		if !inMetadata {
			content = append(content, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	// Join the remaining content into a single string
	return strings.Join(content, "\n"), nil
}

func GetBlogOne(postName string) (string, error) {
	content, err := removeMetadata("data/posts/" + postName + ".md")

	if err != nil {
		return "", err
	}

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	var buf bytes.Buffer
	if err := md.Convert([]byte(content), &buf); err != nil {
		return "", errors.New("failed to convert markdown to HTML")
	}

	return buf.String(), nil
}
