package main

import (
	"2021_1_Noskool_team/configs"
	profiles "2021_1_Noskool_team/internal/app/profiles/delivery/http"
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"time"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
}

func main() {
	time.Sleep(10 * time.Second)
	flag.Parse()

	config := configs.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := profiles.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
