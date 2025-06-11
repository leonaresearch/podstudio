/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stianeikeland/go-rpio/v4"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
	},
}

func init() {
	recordCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func startRecording() (rpio.State, error) {
	err := rpio.Open()

	if err != nil {
		fmt.Printf("Error opening GPIO: %v\n", err)
		return rpio.Low, err
	}

	pin := rpio.Pin(17)

	pin.Input()
	res := pin.Read()

	if res == rpio.High {
		fmt.Println("Starting recording...")
		// Add the start Recording logic here, e.g., initializing audio recording
	}

	return res, nil
}

func stopRecording() (rpio.State, error) {
	err := rpio.Open()
	if err != nil {
		fmt.Printf("Error opening GPIO: %v\n", err)
		return rpio.Low, err
	}

	pin := rpio.Pin(17)

	pin.Input()
	res := pin.Read()

	if res == rpio.Low {
		fmt.Println("Stopping recording...")
		// Add the stop Recording logic here, e.g., finalizing audio recording
	}

	return res, nil
}