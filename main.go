package main

import (
	"blogu/helpers"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Print the version number of blogu",
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()

		fmt.Println("building website...")

		err := helpers.Render()
		if err != nil {
			helpers.ShowErr(err.Error())
			os.Exit(1)
		}

		elapsed := time.Since(start)

		fmt.Println("website build complete in ", elapsed)

	},
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Print the version number of blogu",
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Render()
		if err != nil {
			helpers.ShowErr(err.Error())
			os.Exit(1)
		}
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
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(serverCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
