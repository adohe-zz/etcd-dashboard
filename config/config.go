package config

import (
	"os"
	"flag"
)

type Config struct {
	
	Host string `toml:"host"`
	Port string `toml:"port"`
}

func New() *Config {
	c := new(Config)
	c.Host = "127.0.0.1"
	c.Port = "8080"
	
	return c
}

// Loads configuration from command line flags.
func (c *Config) LoadFlags(arguments []string) error {
	
	f := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	f.StringVar(&c.Host, "host", c.Host, "")
	f.StringVar(&c.Port, "port", c.Port, "")
	
	if err := f.Parse(arguments); err != nil {
		return err
	}	
	
	return nil
}