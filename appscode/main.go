package main

import (
	"fmt"

	"github.com/heheh13/cobra-practise/appscode/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
