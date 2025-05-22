package main

import (
	"fmt"
	"os"
)

var version = "dev"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: example-cli <command> [args]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "echo":
		if len(os.Args) < 3 {
			fmt.Println("Usage: example-cli echo <message>")
			os.Exit(1)
		}
		fmt.Println(os.Args[2])
	case "version":
		fmt.Print(version)
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}
