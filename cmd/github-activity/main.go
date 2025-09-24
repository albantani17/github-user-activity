package main

import (
	"fmt"
	"os"

	"github.com/albantani17/github-user-activity/internal/github"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: github-activity <username>")
		os.Exit(1)
	}

	username := os.Args[1]

	github.GetActivity(username)
}