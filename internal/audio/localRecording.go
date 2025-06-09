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

func StartRecording(device AudioSource, sampleRate int, outputFile string) error {
	r := getRecorder()
	if r.active {
		return fmt.Errorf("recording already in progress")
	}
	if device.Name == "" {
		return fmt.Errorf("invalid audio device: %v", device)
	}
	// Use parecord instead of arecord, with PulseAudio options
	cmd := exec.Command("parecord",
		"--device", device.Name, // Use the symbolic name or description as PulseAudio source
		"--rate", fmt.Sprintf("%d", sampleRate),
		"--channels", "1",
		"--format", "s16le",
		outputFile,
	)
	r.cmd = cmd

	// Capture stdout and stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout: %v", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to get stderr: %v", err)
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	fmt.Printf("Started recording process with PID: %d\n", cmd.Process.Pid)

	// Print logs from stdout and stderr
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := stdout.Read(buf)
			if n > 0 {
				fmt.Printf("[parecord stdout] %s", string(buf[:n]))
			}
			if err != nil {
				break
			}
		}
	}()
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := stderr.Read(buf)
			if n > 0 {
				fmt.Printf("[parecord stderr] %s", string(buf[:n]))
			}
			if err != nil {
				break
			}
		}
	}()

	r.active = true
	// play sounds/recording-start-sound.wav
	cmd = exec.Command("paplay", "sounds/recording-start-sound.wav")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Warning: Failed to play start sound: %v\n", err)
	}

	return nil
}

func StopRecording() error {
	r := getRecorder()
	if !r.active || r.cmd == nil || r.cmd.Process == nil {
		return fmt.Errorf("no active recording to stop")
	}
	// play sounds/recording-end-sound.wav
	cmd := exec.Command("paplay", "sounds/recording-end-sound.wav")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Warning: Failed to play end sound: %v\n", err)
	}
	err := r.cmd.Process.Kill()
	if err != nil {
		return err
	}
	r.active = false
	r.cmd = nil
	return nil
}
