package mware

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//Config is wrapper for config tasks
type Config struct {
	*viper.Viper
}

//New intializes config defaults
func (c *Config) New() {
	c.Viper = viper.New()
	c.AddConfigPath("/etc/conf/") // path to look for the config file in
	c.AddConfigPath("/conf")      // call multiple times to add many search paths
	c.AddConfigPath("../conf/")
	c.AddConfigPath("./mware/conf/")
	c.SetConfigName("config")
	c.SetConfigType("yaml")

	err := c.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}

	c.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	c.WatchConfig()

}
