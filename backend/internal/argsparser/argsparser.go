package argsparser

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

type CommandLineArguments struct {
	ConfigFilePath string
	StoragePath    string
	PresetsPath    string
}

func Parse() *CommandLineArguments {
	result := &CommandLineArguments{}

	parser := argparse.NewParser("llm-workbench", "Web UI for instructing Large Language Models")

	configPath := parser.String("c", "config",
		&argparse.Options{
			Required: false,
			Help:     "Path to the configuration file",
			Default:  "backend.yaml"})

	storagePath := parser.String("s", "storage",
		&argparse.Options{
			Required: false,
			Help:     "Path to the session data storage directory",
			Default:  "data"})

	presetsPath := parser.String("p", "presets",
		&argparse.Options{
			Required: false,
			Help:     "Path to the file containing generation parameter presets",
			Default:  "presets.yaml"})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		return nil
	}

	result.ConfigFilePath = *configPath
	result.StoragePath = *storagePath
	result.PresetsPath = *presetsPath

	return result
}
