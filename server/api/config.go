package api

import (
	"fmt"
	"github.com/deeprave/go-auth/lib"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type Env string

//goland:noinspection GoUnusedConst
const Production, Staging, NonProd, Development string = "prod", "stg", "np", "dev"

type Config struct {
	Version string   `yaml:"version"`
	Env     string   `yaml:"env"`
	Host    string   `yaml:"host"`
	Port    int      `yaml:"port"`
	Cors    []string `yaml:"cors"`
	Auth    Auth     `yaml:"auth"`
}

func NewConfig(version string) *Config {
	return &Config{
		Version: version,
		Env:     Production,
		Port:    9000,
		Host:    "localhost",
		Cors: []string{
			"http://localhost:9000",
		},
		Auth: Auth{
			Issuer:        "acme.co",
			Audience:      "acme.co",
			Secret:        RandString(32),
			TokenExpiry:   ParseDuration(DefaultTokenExpiry),
			RefreshExpiry: ParseDuration(DefaultTokenRefresh),
			Cookie: Cookie{
				Path: "/",
				Name: "__Host-refresh_token",
			},
		},
	}
}

func NewConfigFromFile(filename string, version string, v ...any) (*Config, error) {
	cfg := NewConfig(version)
	return cfg, cfg.Read(filename, version, v...)
}

func (cfg *Config) VersionOk(version string) bool {
	return version == cfg.Version
}

func (cfg *Config) Address() string {
	return fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
}

func (cfg *Config) Read(filename string, version string, v ...any) error {
	filepath, err := lib.FindFile(filename, v...)
	if err == nil {
		var data []byte
		if data, err = os.ReadFile(filepath); err == nil {
			err = yaml.Unmarshal(data, cfg)
		}
		if err == nil && !cfg.VersionOk(version) {
			err = fmt.Errorf("config version '%s' is not compatible", version)
		}
	}
	return err
}

func (cfg *Config) Write(filename string, v ...any) (string, error) {
	filepath, err := lib.FindFile(filename, v...)
	if err != nil {
		filepath = filename
	}
	var data []byte
	data, err = yaml.Marshal(cfg)
	if err == nil {
		err = os.WriteFile(filepath, data, 0644)
	}
	return filepath, err
}

func (cfg *Config) AllowOrigin(origin string) string {
	result := origin // by default, allow from any origin
	for index, orig := range cfg.Cors {
		if index == 0 { // if configured, use first by default if no match
			result = orig
		}
		if strings.HasPrefix(origin, orig) {
			return orig // otherwise return an exact (prefix) match
		}
	}
	return result
}
