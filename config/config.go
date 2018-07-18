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

func (conf *Config) SetListenAddr(listenAddr string) {
	conf.ListenAddr = listenAddr
}

func (conf *Config) SetRemoteAddr(remoteAddr string) {
	conf.RemoteAddr = remoteAddr
}

func (conf *Config) SetPassword(password string) {
	conf.Password = password
}

func init() {
	home, _ := homedir.Dir()
	configFilename := ".ss.json"
	configPath = path.Join(home, configFilename)
}

func (config *Config) SaveConfig() {
	configJson, _ := json.MarshalIndent(config, "", "	")
	err := ioutil.WriteFile(configPath, configJson, 0644)
	if err != nil {
		fmt.Errorf("failed to save conf file %s ERR: %s", configPath, err)
	}
	log.Printf("successfully save to conf file %s \n", configPath)
}

func (config *Config) ReadConfig() {
	if _, err := os.Stat(configPath); !os.IsNotExist(err) {
		log.Printf("reading conf file %s \n", configPath)
		file, err := os.Open(configPath)
		if err != nil {
			log.Fatalf("failed to open conf file %s ERR:%s", configPath, err)
		}
		defer file.Close()
		err = json.NewDecoder(file).Decode(config)
		if err != nil {
			log.Fatalf("Unformatting json file:\n%s", file)
		}
	}
}