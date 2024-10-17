package main

import (
	"blogu/helpers"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Print the version number of blogu",
	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("target")

		// Check if target is empty or not from python, node, or html
		if target == "" {
			helpers.ShowErr("error: target is empty")
			os.Exit(1)
		}

		// Valid targets: python, node, html
		allowedTargets := map[string]bool{
			"html":   true,
			"python": true,
			"node":   true,
		}

		if !allowedTargets[target] {
			helpers.ShowErr(fmt.Sprintf("error: invalid target '%s'. Allowed targets: python, node, html", target))
			os.Exit(1)
		}

		// Output message for the successful build
		switch target {
		case "html":
			err := helpers.BuildForHtml()
			if err != nil {
				helpers.ShowErr(err.Error())
				os.Exit(1)
			}
			fmt.Println("website build complete for html")
		case "python":
			err := helpers.BuildForPython()
			if err != nil {
				helpers.ShowErr(err.Error())
				os.Exit(1)
			}
			fmt.Println("website build complete for python")
		case "node":
			err := helpers.BuildForNode()
			if err != nil {
				helpers.ShowErr(err.Error())
				os.Exit(1)
			}
			fmt.Println("website build complete for node js")
		}

	},
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Print the version number of blogu",
	Run: func(cmd *cobra.Command, args []string) {
		helpers.Render()
		fmt.Println("website running on http://localhost:3000")
		helpers.HtmlServer()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of blogu",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("blogu Static Site Generator v0.1.0")
	},
}

var rootCmd = &cobra.Command{
	Use:   "blogu",
	Short: "blogu is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://blogu.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello blogu")
	},
}

func Execute() {
	buildCmd.Flags().StringP("target", "t", "html", "The target to build like = html, python, nodejs")

	// Add the version command to the root command
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(serverCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
