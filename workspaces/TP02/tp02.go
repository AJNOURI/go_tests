package main

import (
	"fmt"
	"formation-go/workspaces/TP02/golaunch"
)

func main() {
	d := golaunch.GetDaysSinceGoLaunch()
	fmt.Printf("Hello it's been %d days since Go Launch\n", int(d))
}
