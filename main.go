package main

import (
	"github.com/koushiTakayama423/go-unit-test/cmd"
)

func main() {
	csv, err := cmd.Input_csv("input/input.csv")
	if err != nil {
		panic("")
	}

	cmd.Output_csv(csv)
}
