package audio

import (
	"fmt"
	"os/exec"
	"sync"
)

var (
	recorderInstance *Recorder
	recorderOnce     sync.Once
)

type Recorder struct {
	cmd    *exec.Cmd
	active bool
}

func getRecorder() *Recorder {
	recorderOnce.Do(func() {
		recorderInstance = &Recorder{}
	})
	return recorderInstance
}

func StartRecording(device InputDevice, sampleRate int, outputFile string) error {
	r := getRecorder()
	if r.active {
		return fmt.Errorf("recording already in progress")
	}
	if device.Description == "" {
		return fmt.Errorf("invalid audio device: %v", device)
	}
	cmd := exec.Command("arecord",
		"-D", fmt.Sprintf("hw:%d,%d", device.Card, device.Device),
		"-f", "S16_LE",
		"-r", fmt.Sprintf("%d", sampleRate),
		outputFile,
	)
	r.cmd = cmd
	err := cmd.Start()
	if err != nil {
		return err
	}
	r.active = true
	return nil
}

func StopRecording() error {
	r := getRecorder()
	if !r.active || r.cmd == nil || r.cmd.Process == nil {
		return fmt.Errorf("no active recording to stop")
	}
	err := r.cmd.Process.Kill()
	if err != nil {
		return err
	}
	r.active = false
	r.cmd = nil
	return nil
}
