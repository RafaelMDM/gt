package gt

import "github.com/hellflame/argparse"

type Options struct {
	Args       []string
	ConfigPath string
}

func GetOptions() (*Options, error) {
	parser := argparse.NewParser("gt", "Save and GoTo your favorite locations", &argparse.ParserConfig{
		DisableDefaultShowHelp: true,
	})

	args := parser.Strings("a", "args", &argparse.Option{
		Positional: true,
		Required:   false,
		Default:    "",
	})

	config_path := parser.String("c", "config", &argparse.Option{
		Required: false,
		Default:  "",
	})

	err := parser.Parse(nil)
	if err != nil {
		return nil, err
	}

	return &Options{
		Args:       *args,
		ConfigPath: *config_path,
	}, nil
}
