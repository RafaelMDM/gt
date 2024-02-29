package gt

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
)

type Data struct {
	Locations map[string]string `json:"locations"`
}

type GotoCLI struct {
	config *Config
	data   *Data
}

func (g *GotoCLI) Goto(target string) (out string, err error) {
	location, ok := g.data.Locations[target]
	if !ok {
		return "", fmt.Errorf("Cannot find target %q", target)
	}

	return location, nil
}

func (g *GotoCLI) List() string {
	var out strings.Builder
	out.WriteString("Usage: gt [location]\nSaved locations:\n\n")
	for key, value := range g.data.Locations {
		str := fmt.Sprintf("Location %q: Path %q\n", key, value)
		out.WriteString(str)
	}

	return out.String()
}

func (g *GotoCLI) Add(name, location string) error {
	if !path.IsAbs(location) {
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}

		location = path.Join(pwd, location)
	}
	g.data.Locations[name] = location
	return nil
}

func (g *GotoCLI) Remove(name string) error {
	_, ok := g.data.Locations[name]
	if !ok {
		return fmt.Errorf("Cannot find location %q", name)
	}

	delete(g.data.Locations, name)
	return nil
}

func (g *GotoCLI) Execute() (out string, err error) {
	switch g.config.Command {
	case Goto:
		target := g.config.Args[0]
		out, err = g.Goto(target)
	case List:
		out = g.List()
	case Add:
		name := g.config.Args[0]
		location := g.config.Args[1]
		err = g.Add(name, location)
	case Remove:
		name := g.config.Args[0]
		err = g.Remove(name)
	default:
		err = fmt.Errorf("Unknown command %q", g.config.Command)
	}

	return
}

func (g *GotoCLI) Save() error {
	dir := path.Dir(g.config.ConfigPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	bytes, err := json.Marshal(g.data)
	if err != nil {
		return err
	}

	err = os.WriteFile(g.config.ConfigPath, bytes, 0755)
	if err != nil {
		return err
	}

	return nil
}

func defaultGotoCLI(config *Config) *GotoCLI {
	return &GotoCLI{
		config: config,
		data: &Data{
			Locations: map[string]string{},
		},
	}
}

func NewGotoCLI(config *Config) (*GotoCLI, error) {
	_, err := os.Stat(config.ConfigPath)
	if err != nil {
		return defaultGotoCLI(config), err
	}

	contents, err := os.ReadFile(config.ConfigPath)
	if err != nil {
		return defaultGotoCLI(config), err
	}

	var data Data
	err = json.Unmarshal(contents, &data)
	if err != nil {
		return defaultGotoCLI(config), err
	}

	return &GotoCLI{
		config: config,
		data:   &data,
	}, nil
}
