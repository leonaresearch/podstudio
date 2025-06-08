/*
Copyright © 2025 René Kuhn renekuhn@posteo.de

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "podstudio",
	Short: "Accessible podcast recording and management CLI",
	PodStudio is a command-line tool designed for accessible podcast production.
Built specifically for blind and non-tech-savvy users, PodStudio provides:

- Simple recording control with audio feedback
- Admin dashboard for managing podcast sessions
- Automated config file generation
- Integration with Mumble servers for remote guests
- Macro keyboard support for one-button recording
- Syncthing integration for seamless file sharing

Examples:
//   podstudio record start     # Start recording with audio confirmation
  podstudio config generate  # Create configuration files
  podstudio status           # Check system status with audio feedback

PodStudio makes professional podcast recording accessible to everyone,
with clear audio feedback and simple commands that work great with
screen readers and assistive technology.`,
}
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.podstudio.toml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".thomas" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("toml") // Set the config file type to toml
		viper.SetConfigName(".podstudio")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
