/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/leonaresearch/podstudio/internal/screenreader"
	"github.com/spf13/cobra"
)

// recordCmd represents the record command
var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "Start a local podcast recording",
	Long: `Start a local podcast recording with audio feedback and accessible controls.

Examples:
  podstudio record


This command is designed for accessible podcast production, providing clear feedback and simple operation.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("record called")
		screenreader.SayLine("Aufnahme gestartet")
		
	},
}

func init() {
	rootCmd.AddCommand(recordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
