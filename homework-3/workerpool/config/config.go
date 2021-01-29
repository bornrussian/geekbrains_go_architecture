package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Victim  Victim  `yaml:"victim"`
	Options Options `yaml:"options"`
}

type Victim struct {
	HTTPUrl	string `yaml:"http-url"`
}

type Options struct {
	Mode    string `yaml:"mode"`
	Until   int64  `yaml:"until"`
	Threads int64  `yaml:"threads"`
}

func ReadConfig(path string) (*Config, error) {
	f, _ := os.Open(path)
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	conf := Config{}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
