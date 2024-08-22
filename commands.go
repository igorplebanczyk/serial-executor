package main

import (
	"gopkg.in/yaml.v2"
	"os"
	"os/exec"
)

type Command struct {
	Name   string
	Script string
}

type Config struct {
	Commands []Command `yaml:"commands"`
}

func LoadCommands() (*Config, error) {
	file, err := os.Open("config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (cfg *Config) GetCommand(name string) *Command {
	for _, cmd := range cfg.Commands {
		if cmd.Name == name {
			return &cmd
		}
	}
	return nil
}

func (cmd *Command) Run() error {
	winCommand := exec.Command("cmd", "/C", cmd.Script)

	_, err := winCommand.CombinedOutput() // Capture potential errors
	if err != nil {
		return err
	}

	return nil
}
