/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show or generate configuration for PodStudio",
	Long: `Manage PodStudio configuration.

Usage:
  podstudio config            Show the current configuration settings
  podstudio config generate   Start an interactive form to generate a config file

Examples:
  podstudio config            # Print the current config settings
  podstudio config generate   # Start guided config file creation
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && args[0] == "generate" {
			cmd.Help()
			return
		}
		// Print all current config settings
		fmt.Println("Current configuration:")
		settings := viper.AllSettings()
		if len(settings) == 0 {
			fmt.Println("  No configuration loaded.")
			return
		}
		for k, v := range settings {
			fmt.Printf("  %s = %v\n", k, v)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
