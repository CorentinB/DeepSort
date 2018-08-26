package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/akamensky/argparse"
)

var arguments = struct {
	Input        string
	Output       string
	OutputChoice bool
	URL          string
	DryRun       bool
	Recursive    bool
	Jobs         int
	Network      string
}{
	// Default arguments
	OutputChoice: false,
	Jobs: 1,
}

func argumentParsing(args []string) {
	// Create new parser object
	parser := argparse.NewParser("deepsort", "AI powered image tagger backed by DeepDetect")

	// Create flags
	URL := parser.String("u", "url", &argparse.Options{
		Required: true,
		Help: "URL of your DeepDetect instance (i.e: http://localhost:8080)"})

	input := parser.String("i", "input", &argparse.Options{
		Required: true,
		Help: "Your input folder."})

	//output := parser.String("o", "output", &argparse.Options{
	//	Required: false,
	//	Help: "Your output folder, if output is set, " +
	//		"original files will not be renamed, " +
	//		"but the renamed version will be copied in the output folder."})

	network := parser.Selector("n", "network", []string{"resnet-50", "googlenet"}, &argparse.Options{
		Required: false,
		Help: "The pre-trained deep neural network you want to use, " +
			"can be resnet-50 or googlenet",
		Default: "resnet-50"})

	recursive := parser.Flag("R", "recursive", &argparse.Options{
		Required: false,
		Help: "Process files recursively."})

	jobs := parser.Int("j", "jobs", &argparse.Options{
		Required: false,
		Help: "Number of parallel jobs",
		Default: 1})

	dryRun := parser.Flag("d", "dry-run", &argparse.Options{
		Required: false,
		Help: "Just classify images and return results, do not apply."})

	// Parse input
	err := parser.Parse(args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		os.Exit(0)
	}

	// Handle the input flag
	inputFolder, _ := filepath.Abs(*input)
	//arguments.Output = outputFolder
	//arguments.OutputChoice = true
	// Finally save the collected flags
	arguments.Network = *network
	arguments.Jobs = *jobs
	arguments.DryRun = *dryRun
	arguments.Recursive = *recursive
	arguments.Input = inputFolder
	arguments.URL = *URL
}
