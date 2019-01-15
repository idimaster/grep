package main

import (
	"os"
	"fmt"
)

func main() {
	operations := [2]Operations{CountDOS{}, Match{}}
	for _, op := range operations {
		if op.Validate(os.Args) {
			code, message, err := op.Execute(os.Args)
			if message != "" {
				fmt.Println(message)
			}
			if err != nil {
				fmt.Println(err)
			}
			os.Exit(code)
		}
	}
	fmt.Println("Supported operations:")
	for _, op := range operations {
		fmt.Println(op.Help())
	}
}
