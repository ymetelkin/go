package main

import (
	"fmt"
)

func main() {
	alerts := newAlertsFix()
	/*
		e := qa()
		alerts := alerts{
			Config:  e.SSO,
			FeedURL: e.FeedURL,
		}
	*/
	alerts.Fix()
	fmt.Println("\nDone! Press ENTER to exit...")
	var s string
	fmt.Scanln(&s)
}
