package Flight

import (
	"os"
)

type struct Config {
	filename string
}

func (c *Config) Filename() string {
	return c.filename
}

func NewConfig() *Config {
	return &Config{
		Filename: os.Getenv("FLIGHT_DATA"),
	}
}