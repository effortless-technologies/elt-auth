// +build prod

package config

func MakeConfig() *Config {
	c := new(Config)
	c.Db = "104.198.34.190:27017"

	return c
}