package api

import (
	"github.com/deeprave/go-testutils/test"
	"testing"
)

var testConfig = Config{
	Version: "1.0",
	Env:     "dev",
	Host:    "localhost",
	Port:    9000,
	Auth: Auth{
		Issuer:        "Acme, Inc",
		Audience:      "Acme and Clients",
		Secret:        RandString(32),
		TokenExpiry:   ParseDuration("1d"),
		RefreshExpiry: ParseDuration("30m"),
		Cookie: Cookie{
			//Domain: "acme.co",
			Path: "/",
			Name: "__Host_acme_jwt",
		},
	},
}

func MakeTestApp() *App {
	cfg := &testConfig
	app := NewApp(cfg, "test ", "test", "")
	return app
}

func TestApp(t *testing.T) {
	app := MakeTestApp()
	test.ShouldBeEqual(t, app.Config(), &testConfig)
}

func TestAppAddress(t *testing.T) {
	app := MakeTestApp()
	test.ShouldBeEqual(t, app.Address(), "localhost:9000")
}
