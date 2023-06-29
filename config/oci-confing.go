package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type OCIConfig struct {
	OciVersion    string         `json:"ociVersion"`
	ProcessConfig ProcessConfig  `json:"process"`
	Hostname      string         `json:"hostname"`
	MountsConfig  []MountsConfig `json:"mounts"`
}

type ProcessConfig struct {
	Terminal bool           `json:"terminal"`
	User     map[string]int `json:"user"`
	Args     []string       `json:"args"`
	Env      []string       `json:"env"`
	Cwd      string         `json:"cwd"`
}

type MountsConfig struct {
	Destination string   `json:"destination"`
	Source      string   `json:"source"`
	Type        string   `json:"type"`
	Options     []string `json:"options,omitempty"`
}

func GetOCIConfig(containerDirPath string) (*OCIConfig, error) {
	configFile, err := os.Open(fmt.Sprintf("%s/config.json", containerDirPath))
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	bytes, _ := io.ReadAll(configFile)
	if err != nil {
		return nil, err
	}

	var config OCIConfig
	json.Unmarshal(bytes, &config)
	if err := json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
