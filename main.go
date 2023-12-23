package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

var validSuffixes = []string{"s", "m", "h"}

//go:embed resources
var embedFS embed.FS

var completedSoundPath = "resources/completed.wav"
var completedIconPath = "resources/stopwatch.png"

func main() {

	if len(os.Args) != 2 {
		printHelp()
		os.Exit(1)
	}

	timeString := os.Args[1]

	suffix := string(timeString[len(timeString)-1])
	if !slices.Contains(validSuffixes, suffix) {
		fmt.Printf(
			"time string suffix must be one of %s\n",
			strings.Join(validSuffixes, ", "),
		)
		os.Exit(1)
	}

	prefix := string(timeString[0 : len(timeString)-1])
	rawTimeFloat, err := strconv.ParseFloat(
		prefix,
		64)
	if err != nil {
		fmt.Printf("%s not parseable as a float", prefix)
		os.Exit(1)
	}

	var seconds = time.Duration(0)
	switch suffix {
	case "s":
		seconds = time.Duration(rawTimeFloat * float64(time.Second))
	case "m":
		seconds = time.Duration(rawTimeFloat * float64(time.Minute))
	case "h":
		seconds = time.Duration(rawTimeFloat * float64(time.Hour))
	default:
	}

	fmt.Printf("Sleeping %s\n", seconds)

	time.Sleep(seconds)

	showNotification(timeString)
	playSound()
}

func printHelp() {
	fmt.Println(
		`
wt is a Workout Timer. Pass it a short timestring like 30s, 1m or 1h to
run a timer for that length and then trigger an OS notification and play
a sound.

Example:

    wt 0.5m
		`,
	)
}

// showNotification shows an OS notification
func showNotification(timeString string) {
	messageBody := fmt.Sprintf("%s has elapsed!", timeString)

	embeddedIcon, err := embedFS.Open(completedIconPath)
	if err != nil {
		panic(err)
	}
	defer embeddedIcon.Close()

	tempIcon, err := os.CreateTemp("", "wt*.png")
	if err != nil {
		panic(err)
	}
	defer tempIcon.Close()

	_, err = io.Copy(tempIcon, embeddedIcon)
	if err != nil {
		panic(err)
	}

	tempIconPath := tempIcon.Name()

	err = beeep.Alert("wt timer complete!", messageBody, tempIconPath)
	if err != nil {
		panic(err)
	}
}

// playSound plays the completion sound and blocks until it's finished.
func playSound() {
	f, err := embedFS.Open(completedSoundPath)
	if err != nil {
		panic(err)
	}
	streamer, format, err := wav.Decode(f)
	if err != nil {
		panic(err)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
