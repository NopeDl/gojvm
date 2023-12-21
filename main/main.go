package main

import "fmt"
import "gojvm/core"

func main() {
	cmd := core.NewCmd()
	if cmd.GetHelpFlag() {
		core.PrintUsage()
	} else if cmd.GetHelpFlag() {
		fmt.Println("ez-jvm version 0.0.1")
	} else {
		fmt.Println("start jvm")
	}
}
