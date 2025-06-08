/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/leonaresearch/podstudio/internal/screenreader"
	"github.com/spf13/cobra"
)

// listEspeakVoicesCmd represents the listEspeakVoices command
var listEspeakVoicesCmd = &cobra.Command{
	Use:   "listEspeakVoices",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// List available eSpeak voices
		voices, err := screenreader.ListGermanVoices()
		if err != nil {
			fmt.Println("Error listing eSpeak voices:", err)
			return
		}

		if len(voices) == 0 {
			fmt.Println("No eSpeak voices available.")
			return
		}

		fmt.Println("Available eSpeak voices:")
		for _, voice := range voices {
			fmt.Printf("- %s (%s)\n", voice.Name, voice.Languages)
		}

	},
}

func init() {
	configCmd.AddCommand(listEspeakVoicesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listEspeakVoicesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listEspeakVoicesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
