package main

import (
	"gopkg.in/yaml.v2"
	"os"
	"os/exec"
)

type Port struct {
	Name string `yaml:"name"`
	Baud int    `yaml:"baud"`
}

type Command struct {
	Name   string `yaml:"name"`
	Key    string `yaml:"key"`
	Script string `yaml:"script"`
}

type Config struct {
	Commands []Command `yaml:"commands"`
	Port     Port      `yaml:"port"`
}

func GetConfig() (*Config, error) {
	file, err := os.Open("config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (cfg *Config) GetCommand(key string) *Command {
	for _, cmd := range cfg.Commands {
		if cmd.Key == key {
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
