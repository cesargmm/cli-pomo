package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/mbndr/figlet4go"
	"os"
	"time"
)

func main() {
	// Declare command line parameters
	// TODO: Add check, if parameters hasn't been specified ask on command start up
	focusDuration := flag.Int("focus", 25, "Duration of focus time in minutes")
	breakDuration := flag.Int("break", 5, "Duration of break time in minutes")
	flag.Parse()
	input := bufio.NewScanner(os.Stdin)
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()

	// Display program banner
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorYellow,
	}

	renderStr, _ := ascii.RenderOpts("cli-pomodoro", options)
	fmt.Print(renderStr)

	// Specify cycles
	var cycles int
	fmt.Print("\nHow many cycles do you want to repeat? ")
	_, err := fmt.Scan(&cycles)
	if err != nil {
		return
	}

	// Start new cycle
	// TODO: Add loading bar
	for i := 0; i < cycles; i++ {
		// Start focus phase
		fmt.Printf("Cycle %d: Focus for %d minutes\n", i+1, *focusDuration)
		time.Sleep(time.Duration(*focusDuration) * time.Minute)

		// Beeep at phase's end
		err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
		if err != nil {
			panic(err)
		}

		// Ask the user for an input to start break phase
		fmt.Print("Focus has ended! Press any key to start your break\n")
		input.Scan()

		//Start break phase
		fmt.Printf("Cycle %d: Break for %d minutes\n", i+1, *breakDuration)
		time.Sleep(time.Duration(*breakDuration) * time.Minute)

		// Beeep at phase's end
		err = beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
		if err != nil {
			panic(err)
		}

		// Ask the user for an input to start break phase
		fmt.Print("Break has ended! Press any key to start focus\n")
		input.Scan()
	}

	// Finish pomodoro session
	fmt.Println("Pomodoro session complete!")
}
