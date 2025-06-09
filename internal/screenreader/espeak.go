package screenreader

import (
	"bytes"
	"fmt"
	"os/exec"
	"sync"

	"gopkg.in/BenLubar/espeak.v2"
)

var (
	speechReady bool
	speechMutex sync.Mutex
)

func InitSpeech() {
	ctx := &espeak.Context{}
	ctx.SetRate(175)
	ctx.SetVolume(100)
	ctx.SetVoice("de") // Use German voice
	speechReady = true
}

// List all German Vocies
func ListGermanVoices() ([]*espeak.Voice, error) {
	voices := espeak.ListVoices()

	var germanVoices []*espeak.Voice
	for _, v := range voices {
		// Some voices may have Language as "de", "deu", "de-DE", or similar
		for _, lang := range v.Languages {
			if lang.Name == "de" || lang.Name == "deu" || lang.Name == "de-DE" {
				germanVoices = append(germanVoices, v)
				break
			}
		}
	}
	return germanVoices, nil
}

// PrintGermanVoices prints detailed info about available German voices
func PrintGermanVoices() error {
	voices, err := ListGermanVoices()
	if err != nil {
		return err
	}
	fmt.Println("Available German eSpeak voices:")
	fmt.Println((voices))
	return nil
}

// SayLine speaks text and prints it to stdout
func SayLine(text string) {
	fmt.Println(text)

	speechMutex.Lock()
	defer speechMutex.Unlock()

	if !speechReady {
		fmt.Println("Speech system not initialized. Please call InitSpeech() first.")
		return
	}

	ctx := &espeak.Context{}
	ctx.SetRate(175)
	ctx.SetVolume(100)
	ctx.SetVoice("de") // Use German voice or change as needed
	err := ctx.SynthesizeText(text)
	if err != nil {
		fmt.Printf("espeak error: %v\n", err)
		return
	}

	var buf bytes.Buffer
	_, err = ctx.WriteTo(&buf)
	if err != nil {
		fmt.Printf("espeak WriteTo error: %v\n", err)
		return
	}

	cmd := exec.Command("aplay", "-q")
	cmd.Stdin = &buf
	if err := cmd.Run(); err != nil {
		fmt.Printf("aplay error: %v\n", err)
	}
}


