package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
	)

type Config struct {
	Server   		ServerConfig   				`yaml:"server"`
	Telegram 		TelegramConfig 				`yaml:"telegram"`
	Email    		EmailConfig    				`yaml:"email"`
	ItemsStorage    ItemsStorageConfig    		`yaml:"items-storage-grpc-server"`
}

type ServerConfig struct {
	DebugMode			bool 	`yaml:"debug"`
	Listen 				string 	`yaml:"listen"`
	Domain 				string 	`yaml:"domain"`
}

type TelegramConfig struct {
	BotToken 			string 	`yaml:"bot-token"`
	ChatID				int64 	`yaml:"chat-id"`
	MsgTmplFile			string	`yaml:"msg-template-file"`
}

type EmailConfig struct {
	SMTPServer 			string 	`yaml:"smtp-server"`
	Port				int64 	`yaml:"port"`
	From				string	`yaml:"from"`
	User				string	`yaml:"user"`
	Password			string	`yaml:"password"`
	MsgTmplFile			string	`yaml:"msg-template-file"`
}

type ItemsStorageConfig struct {
	Address 			string 	`yaml:"address"`
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
