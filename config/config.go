package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

var cfg *config

func Load(path string) error {

	if err := FileExist(path); err != nil {
		return err
	}
	if f, err := os.ReadFile(path); err == nil {
		return yaml.Unmarshal(f, &cfg)
	} else {
		return err
	}
}

func FileExist(path string) error {

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return errors.New("error file \"" + path + "\" not exist")
	} else if err != nil {
		return err
	} else {
		return nil
	}
}

type config struct {
	Cache   cache
	Limiter limiter
	Server  server
	Tokens  []token
}

type cache struct {
	Host string
	Port string
	DB   int
	Pwd  string
}
type server struct {
	Port string
}

type limiter struct {
	LimitRequestPerIp int
	Expiration        int
	BlockingTimeIP    int
	BlockingTimeToken int
}

type token struct {
	Token      string
	Limiter    int
	Expiration int
}

func Cache() cache {
	return cfg.Cache
}
func Limiter() limiter {
	return cfg.Limiter
}
func Server() server {
	return cfg.Server
}
func Token() []token {
	return cfg.Tokens
}
