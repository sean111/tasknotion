package backend

import (
	conf "github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	configDir string
	cacheDir string

	config struct{
		APIKey string `mapstructure:"api_key"`
		DatabaseID string `mapstructure:"database_id"`
	}
)

func LoadConfig() {
	systemOS := runtime.GOOS
	if systemOS == "windows" {
		configDir = filepath.Join(os.Getenv("APPDATA"), "tasknotion")
		cacheDir = filepath.Join(os.Getenv("LOCALAPPDATA"), "tasknotion")
	} else {
		configDir = filepath.Join(os.Getenv("HOME"), ".config", "tasknotion")
		cacheDir = filepath.Join(os.Getenv("HOME"), ".cache", "tasknotion")
	}

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.MkdirAll(configDir, 0755)
		os.MkdirAll(cacheDir, 0755)
		panic("first run blurb here, config & cache folders created")
	}
	conf.AddDriver(yaml.Driver)

	err := conf.LoadFiles(filepath.Join(configDir, "config.yaml"))
	if err != nil {
		log.Fatalln("Could not find the config file: ", err)
	}

	err = conf.BindStruct("", &config)
	if err != nil {
		log.Fatalln("Error loading the config file", err)
	}
}

