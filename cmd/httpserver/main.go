package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/timickb/url-shortener-client/internal/server"
	"gopkg.in/yaml.v3"
)

const (
	defaultConfigPath = "config.yml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config.yml", "path to config in filesystem")
}

type config struct {
	api          string `yaml:"api"`
	frontendPath string `yaml:"frontendPath"`
	port         int    `yaml:"port"`
}

func main() {
	logger := logrus.New()
	flag.Parse()

	if configPath == "" {
		configPath = strings.Clone(defaultConfigPath)
	}

	configContent, ioErr := ioutil.ReadFile(configPath)
	conf := &config{}

	if ioErr != nil {
		log.Fatal(ioErr)
	}

	if err := yaml.Unmarshal(configContent, &conf); err != nil {
		log.Fatal(err)
	}

	srv := server.New(logger, conf.port, conf.api, conf.frontendPath)
	srv.Start()
}
