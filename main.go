package main

import (
	"os"

	"github.com/highxshell/assignm1BCPopov/cli"
)

func main() {
	defer os.Exit(0)
	cli := cli.CommandLine{}
	cli.Run()
}