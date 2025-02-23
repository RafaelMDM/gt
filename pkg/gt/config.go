package gt

import (
	"fmt"
	"os"
	"path"
)

type Command = int

const (
	Goto Command = iota
	Add
	Remove
	List
)

type Config struct {
	Args       []string
	Command    Command
	ConfigPath string
}

func getCommand(options *Options) Command {
	if len(options.Args) == 0 {
		return List
	}

	switch options.Args[0] {
	case "add":
		return Add
	case "rm":
		return Remove
	case "list":
		return List
	default:
		if len(options.Args) == 1 {
			return Goto
		}
		return Add
	}
}

func getConfigPath(opts *Options) (string, error) {
	if opts.ConfigPath != "" {
		return opts.ConfigPath, nil
	}

	config, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return path.Join(config, "gt", "locations.json"), nil
}

func getArgs(options *Options) ([]string, error) {
	numArgs := len(options.Args)
	if numArgs == 0 {
		return []string{}, nil
	}

	command := getCommand(options)
	switch command {
	case Goto:
		// gt [location_name]
		if numArgs != 1 {
			return nil, fmt.Errorf("GoTo command expects 1 argument, but got %d", numArgs)
		}
		return options.Args, nil
	case List:
		// gt list
		if numArgs != 1 {
			return nil, fmt.Errorf("List command expects no arguments, but got %d", numArgs-1)
		}
		return options.Args, nil
	case Add:
		// gt add [location_name] [location]
		if options.Args[0] == "add" {
			if numArgs != 3 {
				return nil, fmt.Errorf("Add command expects 2 arguments, but got %d", numArgs-1)
			}
			return options.Args[1:], nil
		}
		// gt [location_name] [location]
		if numArgs != 2 {
			return nil, fmt.Errorf("Add command expects 2 arguments, but got %d", numArgs)
		}
		return options.Args, nil
	default:
		// gt rm [location_name]
		if numArgs != 2 {
			return nil, fmt.Errorf("Remove command expects 1 arguments, but got %d", numArgs-1)
		}
		return options.Args[1:], nil
	}
}

func NewConfig(options *Options) (*Config, error) {
	command := getCommand(options)

	args, err := getArgs(options)
	if err != nil {
		return nil, err
	}

	config_path, err := getConfigPath(options)
	if err != nil {
		return nil, err
	}

	return &Config{
		Args:       args,
		Command:    command,
		ConfigPath: config_path,
	}, nil
}
