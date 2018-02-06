package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"github.com/mitchellh/go-homedir"
)

var (
	configPath string
)

type Config struct {
	ListenAddr string `json:"listen"`
	RemoteAddr string `json:"remote"`
	Password   string `json:"password"`
}

func init() {
	home, _ := homedir.Dir()
	configFilename := ".ss.json"
	if len(os.Args) == 2 {
		configFilename = os.Args[1]
	}
	configPath = path.Join(home, configFilename)
}

func (config *Config) SaveConfig() {
	configJson, _ := json.MarshalIndent(config, "", "	")
	err := ioutil.WriteFile(configPath, configJson, 0644)
	if err != nil {
		fmt.Errorf("failed to save configuration file %s ERR: %s", configPath, err)
	}
	log.Printf("successfully save to configuration file %s \n", configPath)
}

func (config *Config) ReadConfig() {
	if _, err := os.Stat(configPath); !os.IsNotExist(err) {
		log.Printf("reading configuration file %s \n", configPath)
		file, err := os.Open(configPath)
		if err != nil {
			log.Fatalf("failed to open configuration file %s ERR:%s", configPath, err)
		}
		defer file.Close()

		err = json.NewDecoder(file).Decode(config)
		if err != nil {
			log.Fatalf("Unformatting json file:\n%s", file)
		}
	}
}
