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

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printAudioInputDevices() {
	fmt.Println("Available audio input devices:")

	var availableDevices []audio.InputDevice
	cmd := exec.Command("bash", "-c",
		`arecord -l | grep "^card" | sed -E 's/card ([0-9]+): ([^,]+), device ([0-9]+): (.+)/{"card":\1,"name":"\2","device":\3,"description":"\4"}/' | jq -s '.'`)

	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

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
	var selectedDevice audio.InputDevice
	options := make([]huh.Option[string], len(availableDevices))
	deviceMap := make(map[string]audio.InputDevice)
	for i, device := range availableDevices {
		label := fmt.Sprintf("%s (card %d, device %d) - %s", device.Name, device.Card, device.Device, device.Description)
		options[i] = huh.NewOption(label, device.Name)
		deviceMap[device.Name] = device
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select an audio input device").
				Options(options...).
				Value(&selectedName),
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
	viper.Set("inputDevice.card", selectedDevice.Card)
	viper.Set("inputDevice.device", selectedDevice.Device)
	viper.Set("inputDevice.description", selectedDevice.Description)

	err = viper.WriteConfigAs(".podstudio.toml")
	if err != nil {
		log.Fatalf("Failed to write config: %v", err)
	}

	fmt.Println("Device selection saved to .podstudio.toml")
}
