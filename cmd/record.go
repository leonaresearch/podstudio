/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/leonaresearch/podstudio/internal/audio"
	"github.com/leonaresearch/podstudio/internal/screenreader"
	"github.com/spf13/cobra"
)

var cardFlag int
var deviceFlag int
var descriptionFlag string

// recordCmd represents the record command
var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "Start a local podcast recording",
	Long: `Start a local podcast recording with audio feedback and accessible controls.

Examples:
  podstudio record

This command is designed for accessible podcast production, providing clear feedback and simple operation.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the configured microphone
		mic, err := audio.GetMicFromConfig()
		if err != nil {
			fmt.Printf("Error getting microphone: %v\n", err)
			return
		}

		// Create the recordings directory structure
		recordingsPath := "recordings/" + time.Now().Format("2006/01/02")
		if err := os.MkdirAll(recordingsPath, 0755); err != nil {
			fmt.Printf("Error creating recordings directory: %v\n", err)
			return
		}

		// Generate filename with proper numbering
		filename := recordingsPath + "/01.wav"
		if _, err := os.Stat(filename); err == nil {
			// file exists, increment the number
			number := 1
			for {
				filename = recordingsPath + "/" + fmt.Sprintf("%02d", number) + ".wav"
				if _, err := os.Stat(filename); err != nil {
					break
				}
				number++
			}
		}

		err = audio.StartRecording(mic, 44100, filename)
		if err != nil {
			fmt.Printf("Error starting recording: %v\n", err)
			return
		}

		// Announce recording start
		screenreader.SayLine("Aufnahme gestartet")

		// Wait for user input to stop
		fmt.Println("Press Enter to stop recording...")
		fmt.Scanln()

		// Stop recording
		err = audio.StopRecording()
		if err != nil {
			fmt.Printf("Error stopping recording: %v\n", err)
			return
		}
		// say just the number of the filename
		number := filename[len(filename)-6 : len(filename)-4]
		screenreader.SayLine("Aufnahme " + number + " beendet")
	},
}

func init() {
	rootCmd.AddCommand(recordCmd)
}
