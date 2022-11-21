package main

import (
	"fmt"
	"io/ioutil"
	"sigs.k8s.io/yaml"
)

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Debug    bool   `json:"debug"`
}

func main() {
	//read config
	configfile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	var config Config
	err = yaml.Unmarshal(configfile, &config)
	if err != nil {
		panic(err)
	}

	fmt.Println("Username:", config.Username)
	fmt.Println("Password:", config.Password)
	fmt.Println("Debug:", config.Debug)
}
