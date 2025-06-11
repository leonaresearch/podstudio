/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"github.com/charmbracelet/huh"
	"github.com/leonaresearch/podstudio/internal/audio"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	//cobra.OnInitialize(printAudioInputDevices)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printAudioInputDevices() {
	fmt.Println("Available audio input devices:")

	// Use pactl to get sources in JSON
	cmd := exec.Command("bash", "-c", "pactl -f json list sources short")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal JSON into []audio.AudioSource (update AudioSource struct as needed)
	var availableDevices []audio.AudioSource
	err = json.Unmarshal(output, &availableDevices)
	if err != nil {
		log.Fatal(err)
	}

	if len(availableDevices) == 0 {
		fmt.Println("No audio input devices found.")
		return
	}

	// Use huh to prompt the user to select a device
	var selectedName string
	var selectedDevice audio.AudioSource
	options := make([]huh.Option[string], len(availableDevices))
	deviceMap := make(map[string]audio.AudioSource)
	for i, device := range availableDevices {
		label := fmt.Sprintf("%s (index %d)", device.Name, device.Index)
		options[i] = huh.NewOption(label, device.Name)
		deviceMap[device.Name] = device
	}
	var selectedSampleRate int
	var selectedFormat string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select an audio input device").
				Options(options...).
				Value(&selectedName),
			// add sample rate selection
			huh.NewSelect[int]().
				Title("Select sample rate").
				Options(
					huh.NewOption("48000 Hz", 48000),
					huh.NewOption("96000 Hz", 96000),
					huh.NewOption("192000 Hz", 192000),
					huh.NewOption("384000 Hz", 384000),
				).
				Value(&selectedSampleRate),
			// add format selection
			huh.NewSelect[string]().
				Title("Select audio format").
				Options(
					huh.NewOption("WAV", "wav"),
					huh.NewOption("FLAC", "flac"),
					huh.NewOption("MP3", "mp3"),
					huh.NewOption("OGG", "ogg"),
				).
				Value(&selectedFormat),
		),
	)

	err = form.Run()
	if err != nil {
		log.Fatal(err)
	}

	selectedDevice, ok := deviceMap[selectedName]
	if !ok {
		log.Fatalf("Selected device not found: %s", selectedName)
	}

	fmt.Printf("Selected device: %s\n", selectedName)

	// Store device details in config with prefix inputDevice.
	viper.Set("inputDevice.name", selectedDevice.Name)
	viper.Set("inputDevice.index", selectedDevice.Index)
	// If you want to store more info, add more fields from AudioSource as needed
	viper.Set("recording.sampleRate", selectedSampleRate)
	viper.Set("recording.format", selectedFormat)

	err = viper.WriteConfigAs(".podstudio.toml")
	if err != nil {
		log.Fatalf("Failed to write config: %v", err)
	}

	fmt.Println("Device selection saved to .podstudio.toml")
}

