package api

import (
	"fmt"
	"github.com/deeprave/go-auth/repository"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type App struct {
	cfg *Config
	appname,
	version string
	DB repository.DB
}

func NewApp(cfg *Config, logprefix string, appName string, version string) *App {
	log.SetOutput(os.Stdout)
	log.SetPrefix(logprefix)
	log.SetFlags(log.Lmsgprefix | log.Ldate | log.Ltime | log.Lmicroseconds)
	return &App{
		cfg:     cfg,
		appname: appName,
		version: version,
		DB:      nil,
	}
}

func (app *App) Config() *Config {
	return app.cfg
}

func (app *App) Auth() *Auth {
	return &app.Config().Auth
}

func (app *App) Address() string {
	return app.Config().Address()
}

func (app *App) StartServer() error {
	mux := app.Routes()

	address := app.Address()
	log.Printf("%s v%s HTTP Server listening on %s", app.appname, app.version, address)
	return http.ListenAndServe(address, mux)
}

func (app *App) CorsOrigin(r *http.Request) string {
	proto, dport := "http", 80
	if r.TLS != nil {
		proto, dport = "https", 443
	}
	remote := r.RemoteAddr
	if offset := strings.Index(remote, ":"); offset != -1 {
		host := remote[:offset]
		port, err := strconv.ParseInt(remote[offset+1:], 0, 32)
		// strip port if default
		if err == nil && int(port) == dport {
			remote = host
		}
	}
	return app.cfg.AllowOrigin(fmt.Sprintf("%s://%s", proto, remote))
}
