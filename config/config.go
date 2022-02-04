package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	token  string
	prefix string

	config *ConfigStruct
)

type ConfigStruct struct {
	token  string `json:"token"`
	prefix string `json:"prefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println((err.Error()))
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	token = config.token
	prefix = config.prefix

	return nil
}
