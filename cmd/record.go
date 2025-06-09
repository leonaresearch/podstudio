/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/leonaresearch/podstudio/internal/audio"
	"github.com/leonaresearch/podstudio/internal/screenreader"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		// Use flag if set, otherwise use viper config value
		card := cardFlag
		if !cmd.Flags().Changed("card") {
			card = viper.GetInt("inputdevice.card")
		}
		device := deviceFlag
		if !cmd.Flags().Changed("device") {
			device = viper.GetInt("inputdevice.device")
		}
		description := descriptionFlag
		if !cmd.Flags().Changed("description") {
			description = viper.GetString("inputdevice.description")
		}
		audioSource := audio.AudioSource{
			

		// print the selected device
		fmt.Printf("Selected audio device: %s (Card: %d, Device: %d)\n", recordingDevice.Description, recordingDevice.Card, recordingDevice.Device)
		recordingSampleRate := viper.GetInt("recording.samplerate")
		if len(args) > 0 && args[0] == "stop" {
			err := audio.StopRecording()
			if err != nil {
				fmt.Println("Fehler beim Stoppen der Aufnahme:", err)
				screenreader.SayLine("Fehler beim Stoppen der Aufnahme")
				return
			}
			screenreader.SayLine("Aufnahme gestoppt")
			fmt.Println("Aufnahme gestoppt")
			return
		}
		err := audio.StartRecording(recordingDevice, recordingSampleRate, "output.wav")
		if err != nil {
			fmt.Println("Fehler beim Starten der Aufnahme:", err)
			screenreader.SayLine("Fehler beim Starten der Aufnahme")
			return
		}
		screenreader.SayLine("Aufnahme gestartet")
		fmt.Println("Aufnahme gestartet")
	},
}

func init() {
	rootCmd.AddCommand(recordCmd)

	recordCmd.Flags().IntVar(&cardFlag, "card", 0, "Audio card number")
	recordCmd.Flags().IntVar(&deviceFlag, "device", 0, "Audio device number")
	recordCmd.Flags().StringVar(&descriptionFlag, "description", "", "Description of the audio device")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
