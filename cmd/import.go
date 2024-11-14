/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"github.com/inchestnov/bookmarks-cli/internal/core"
	"github.com/inchestnov/bookmarks-cli/internal/domain"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: importFunc,
}

func init() {
	importCmd.Flags().StringP("source", "s", "", "Relative path to source file")
	if err := importCmd.MarkFlagRequired("source"); err != nil {
		panic(err)
	}

	importCmd.Flags().StringP("template", "t", "", "Template for import")
	importCmd.Flags().StringP("engine", "e", "", "Engine for bookmarks")
	importCmd.Flags().StringP("folder", "", "", "Name of folder for bookmarks")

	importCmd.Flags().BoolP("append", "", false, "Append to existing bookmarks")
	importCmd.Flags().BoolP("override", "o", false, "Override existing bookmarks")

	rootCmd.AddCommand(importCmd)
}

func importFunc(cmd *cobra.Command, args []string) {
	sourceContent, err := readSourceContent(cmd, args)
	if err != nil {
		onError(cmd, err)
		return
	}

	opts, err := buildProcessorOptions(cmd, args)
	if err != nil {
		onError(cmd, err)
		return
	}

	if err = core.Process(sourceContent, opts); err != nil {
		onError(cmd, err)
		return
	}
}

func readSourceContent(cmd *cobra.Command, _ []string) ([]byte, error) {
	sourceFileLocation, err := cmd.Flags().GetString("source")
	if err != nil {
		return nil, err
	}

	sourceFileLocationAbsolute, err := filepath.Abs(sourceFileLocation)
	if err != nil {
		return nil, err
	}

	return os.ReadFile(sourceFileLocationAbsolute)
}

func buildProcessorOptions(cmd *cobra.Command, _ []string) (core.ProcessOptions, error) {
	var opts core.ProcessOptions

	template, err := cmd.Flags().GetString("template")
	if err != nil {
		return opts, err
	}
	opts.Template = template

	engine, err := cmd.Flags().GetString("engine")
	if err != nil {
		return opts, err
	}
	opts.Engine = engine

	folder, err := cmd.Flags().GetString("folder")
	if err != nil {
		return opts, err
	}
	opts.Folder = folder

	override, err := cmd.Flags().GetBool("override")
	if err != nil {
		return opts, err
	}

	apnd, err := cmd.Flags().GetBool("append")
	if err != nil {
		return opts, err
	}

	if override && apnd {
		return opts, errors.New("--override and --append are mutually exclusive")
	}

	mode := domain.ImportModeAppend // prioritize
	if override {
		mode = domain.ImportModeOverride
	}
	opts.Mode = mode

	return opts, nil
}
