package main

import (
	"log"
	"reflect"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	StaticPath       string `default:"/etc/backend/static" envconfig:"STATIC_PATH"`
	DatabaseHost     string `default:"db" envconfig:"DATABASE_HOST"`
	DatabasePort     int    `default:"3306" envconfig:"DATABASE_PORT"`
	DatabaseUser     string `default:"user" envconfig:"DATABASE_USER"`
	DatabasePassword string `default:"pass" envconfig:"DATABASE_PASSWORD"`
	DatabaseName     string `default:"supinfo" envconfig:"DATABASE_NAME"`
}

func newConfig() *config {
	return &config{}
}

func (c *config) prettyLog() {
	log.Println("Here is the configuration used: ")
	s := reflect.ValueOf(c).Elem()
	typeOfConfig := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		log.Printf("%s %s = %v\n", typeOfConfig.Field(i).Name, f.Type(), f.Interface())
	}
}

func (c *config) getFromEnv() error {
	err := envconfig.Process("", c)
	if err != nil {
		return err
	}
	return nil
}
