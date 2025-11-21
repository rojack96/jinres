package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/rojack96/jinres/gen"
)

const defaultOut = "status_structs.go"

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, "  jinres init [-o output][-p package]")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Commands:")
	fmt.Fprintln(os.Stderr, "  init    generate status structs file (default output: status_structs.go)")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	cmd := os.Args[1]
	switch cmd {
	case "init":
		initCmd := flag.NewFlagSet("init", flag.ExitOnError)
		out := initCmd.String("o", defaultOut, "output file path")
		pkg := initCmd.String("p", "statusstructs", "package name for the generated file")
		// parse flags for init command (skip program name and command)
		if err := initCmd.Parse(os.Args[2:]); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("generating status structs file at %s...\n", *out)
		if err := gen.GenerateStructsFile(*out, pkg); err != nil {
			log.Fatalf("failed to generate file: %v", err)
		}
		fmt.Printf("generated %s\n", *out)

		// run gofmt -w on the generated file
		cmd := exec.Command("gofmt", "-w", *out)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Printf("gofmt failed: %v", err)
		}

	case "help", "-h", "--help":
		usage()

	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n\n", cmd)
		usage()
		os.Exit(2)
	}
}
