package main

import (
	"github.com/deckarep/gosx-notifier"
)

func notify() {
	if cmd == "start" {
		note := gosxnotifier.NewNotification("ntomata is complete")
		note.Push()
	} else {
		note := gosxnotifier.NewNotification("break is over")
		note.Push()
	}
}
