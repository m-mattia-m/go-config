package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Client struct {
	ConfigData *Config
}

func NewClient(file string) (*Client, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var ConfigData Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&ConfigData)
	if err != nil {
		return nil, err
	}

	return &Client{
		ConfigData: &ConfigData,
	}, nil
}

func (c *Client) GetDatabasePassword() string {
	return c.ConfigData.Database.Password
}

func (c *Client) GetDatabaseUsername() string {
	return c.ConfigData.Database.Username
}
