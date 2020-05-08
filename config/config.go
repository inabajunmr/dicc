package config

import (
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

// Config for dicc
type config struct {
	MerriamWebsterApiKey string `yaml:"merriam-webster-api-key"`
}

// Read config
func read() (*config, error) {
	buf, err := ioutil.ReadFile(getConfigPath())
	if err != nil {
		fmt.Println("Config is something wrong.")
		return &config{}, err
	}

	var d config

	if err = yaml.Unmarshal(buf, &d); err != nil {
		fmt.Println("Config is something wrong.")
		return nil, err
	}

	return &d, nil
}

// Write config
func write(conf *config) {
	buf, err := yaml.Marshal(*conf)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if err = ioutil.WriteFile(getConfigPath(), buf, 0666); err != nil {
		os.Exit(1)
	}
}

func SetMerriamWebsterApiKey(key string) {
	conf, _ := read()
	if conf == nil {
		conf = &config{}
	}
	conf.MerriamWebsterApiKey = key
	write(conf)
}

func GetMerriamWebsterApiKey() string {
	c, _ := read()
	return c.MerriamWebsterApiKey
}

func getConfigPath() string {
	usr, err := user.Current()
	if err != nil {
		os.Exit(1)
	}
	return usr.HomeDir + "/.dicc/config.yaml"
}
