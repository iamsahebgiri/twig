/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/iamsahebgiri/twig/internal/cli"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Generate static html files for hosting",
	Long: `This commands take path to git repository and computes relevant
files in build time and allows you to serve it in any CDN providers.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		cli.Build(path)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.Flags().String("path", "/Users/sahebg/oss/stagit", "Path to github repo")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
