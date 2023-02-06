package api

import (
	"github.com/deeprave/go-testutils/test"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const (
	configFile    = "config.yml"
	configFileOut = "config-copy.yml"
)

func readConfig(t *testing.T) (string, *Config) {
	version := "1.0"
	cfg, err := NewConfigFromFile(configFile, version, ".", "..", "../..")
	if err != nil {
		t.Fatalf("%v", err)
	}
	return version, cfg
}

func TestNewConfigFromFile(t *testing.T) {
	version, cfg := readConfig(t)
	test.ShouldBeTrue(t, cfg.VersionOk(version), "incorrect version")
	test.ShouldBeEqual(t, cfg.Env, Development)
	test.ShouldBeEqual(t, cfg.Host, "localhost")
	test.ShouldBeEqual(t, cfg.Port, 3000)
}

func TestWriteConfigToFile(t *testing.T) {
	version, cfg := readConfig(t)
	test.ShouldBeTrue(t, cfg.VersionOk(version), "incorrect version")

	filepath, err := cfg.Write(configFileOut)
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Logf("os.Remove(%s): %v", name, err)
		}
	}(filepath)
	test.ShouldBeNoError(t, err, "unexpected error: %v", err)
	test.ShouldBeEqual(t, filepath, configFileOut)
}

func TestCorsOrigin(t *testing.T) {
	version, cfg := readConfig(t)
	test.ShouldBeTrue(t, cfg.VersionOk(version), "incorrect version")
	app := NewApp(cfg, "test", "test", version)

	cors := []string{
		"http://localhost:9000",
		"https://localhost:9000",
		"http://localhost:3000",
		"https://localhost:3000",
		"http://localhost",
	}

	var (
		origin string
		r      *http.Request
	)

	// empty cors list, allow any remote
	cfg.Cors = []string{}
	r = httptest.NewRequest(http.MethodGet, "http://localhost:9000/testing-123", nil)
	origin = app.CorsOrigin(r)
	test.ShouldBeEqual(t, origin, "http://192.0.2.1:1234")

	// strict cors list, remoteAddr not in it, returns first
	cfg.Cors = cors
	r = httptest.NewRequest(http.MethodGet, "http://localhost:9000/testing-123", nil)
	origin = app.CorsOrigin(r)
	test.ShouldBeEqual(t, origin, "http://localhost:9000")

	// strict cors list remoteAddr added, returns matched address
	cfg.Cors = append(cors, "http://192.0.2.1:1234")
	r = httptest.NewRequest(http.MethodGet, "http://localhost:9000/testing-123", nil)
	origin = app.CorsOrigin(r)
	test.ShouldBeEqual(t, origin, "http://192.0.2.1:1234")
}
