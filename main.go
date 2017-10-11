package main

import (
	"fmt"
	"os"
	"time"
)

var (
	ntomata    = 25
	shortBreak = 5
	longBreak  = 10
	unit       = time.Minute
)

var cmd string

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	done := make(chan struct{})
	cmd = os.Args[1]
	switch cmd {
	case "start":
		d := time.Duration(ntomata) * unit
		fmt.Println("starting a ntomata -- don't get distracted!")
		updateDisplay(d)
		go startTicker(d, done)
		time.Sleep(d)
	case "sbreak":
		d := time.Duration(shortBreak) * unit
		fmt.Println("taking a short break -- don't wander too far!")
		updateDisplay(d)
		go startTicker(d, done)
		time.Sleep(d)
	case "lbreak":
		d := time.Duration(longBreak) * unit
		fmt.Println("taking a long break -- enjoy it!")
		updateDisplay(d)
		go startTicker(d, done)
		time.Sleep(d)
	default:
		fmt.Printf("unrecognized command\n\n")
		usage()
		return
	}
	done <- struct{}{}
	notify()
}

func startTicker(d time.Duration, done <-chan struct{}) {
	timeLeft := d
	t := time.Tick(1 * time.Second)
	for {
		select {
		case <-t:
			timeLeft = timeLeft - 1*time.Second
			updateDisplay(timeLeft)
		case <-done:
			return
		}
	}
}

func updateDisplay(d time.Duration) {
	fmt.Printf("\rtime left: %s", d)
}

func usage() {
	fmt.Println("Usage: ntomata cmd [flags]")
	fmt.Println("\tcommands:")
	fmt.Println("\t\tstart: starts a ntomata")
	fmt.Println("\t\tsbreak: starts a short break")
	fmt.Println("\t\tlbreak: starts a long break")
}
