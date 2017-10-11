// +build !darwin

package main

func notify() {
	if cmd == "start" {
		fmt.Println("ntomata is complete")
	} else {
		fmt.Println("break is over")
	}
}
