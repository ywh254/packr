package cmd

import (
	"os"

	"github.com/gobuffalo/packr/plog"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var globalOptions = struct {
	Verbose       bool
	IgnoreImports bool
}{}

var rootCmd = &cobra.Command{
	Use:   "packr",
	Short: "A brief description of your application",
	RunE: func(cmd *cobra.Command, args []string) error {
		if globalOptions.Verbose {
			plog.Default.SetLevel(logrus.DebugLevel)
		}
		return pack(args...)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&globalOptions.Verbose, "verbose", "v", false, "enables verbose logging")
	rootCmd.PersistentFlags().BoolVar(&globalOptions.IgnoreImports, "ignore-imports", false, "when set to true packr won't resolve imports for boxes")
}