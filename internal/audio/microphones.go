package audio

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/spf13/viper"
)

func GetAvailableMicrophones() ([]AudioSource, error) {
	// Use pactl to get sources in JSON
	cmd := exec.Command("bash", "-c", "pactl -f json list sources short")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON into []AudioSource
	var availableDevices []AudioSource
	err = json.Unmarshal(output, &availableDevices)
	if err != nil {
		return nil, err
	}

	return availableDevices, nil
}

func GetMicFromConfig() (AudioSource, error) {
	mic := AudioSource{
		Driver:              viper.GetString("inputDevice.driver"),
		Index:               viper.GetInt("inputDevice.index"),
		Name:                viper.GetString("inputDevice.name"),
		SampleSpecification: viper.GetString("inputDevice.sampleSpecification"),
		State:               viper.GetString("inputDevice.state"),
	}
	// check  if the mic is available, otherwise return an error that the mic is not available
	if mic.Name == "" || mic.Index < 0 {
		return AudioSource{}, fmt.Errorf("invalid microphone configuration: %v", mic)
	}
	return mic, nil
}

func SetMicInConfig(mic AudioSource) error {
	viper.Set("inputDevice.driver", mic.Driver)
	viper.Set("inputDevice.index", mic.Index)
	viper.Set("inputDevice.name", mic.Name)
	viper.Set("inputDevice.sampleSpecification", mic.SampleSpecification)
	viper.Set("inputDevice.state", mic.State)

	// Write the updated configuration to the file
	return viper.WriteConfig()
}
