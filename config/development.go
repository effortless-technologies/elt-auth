// +build dev

package config

func MakeConfig() *Config {
	c := new(Config)
	c.Db = "localhost:27017"

	return c
}
