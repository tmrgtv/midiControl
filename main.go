package main

import (
	"fmt"

	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv" // autoregisters driver
)

func main() {
	defer midi.CloseDriver()

	_, err := midi.FindInPort("V-1HD")
	if err != nil {
		fmt.Println("can't find in V-1HD")
	}

	_, err = midi.FindOutPort("V-1HD")
	if err != nil {
		fmt.Println("can't find out V-1HD")
	}
}
