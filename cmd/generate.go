/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"
	""

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Interactively generate a new configuration file",
	Long: `Start an interactive form to generate a PodStudio config file.

This will guide you through the required settings and write them to the config file.
`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement interactive form for config generation
		printAudioInputDevices()
	},
}

func init() {
	configCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}



func printAudioInputDevices() {
	// This function will list available audio input devices
	// and allow the user to select one for recording.
	fmt.Println("Available audio input devices:")



	cmd := exec.Command("bash", "-c", 
		`arecord -l | grep "^card" | sed -E 's/card ([0-9]+): ([^,]+), device ([0-9]+): (.+)/{"card":\1,"name":"\2","device":\3,"description":"\4"}/' | jq -s '.'`)
	
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	
	
	// TODO: Implement device listing and selection
}
