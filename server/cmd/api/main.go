package main

import (
	"flag"
	"fmt"
	"github.com/chigopher/pathlib"
	"github.com/deeprave/go-auth/api"
	"github.com/deeprave/go-auth/repository/db"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	version    = "1.0"
	configName = "config.yml"
)

func setOptions(cfg *api.Config) (bool, bool, string) {
	if err := godotenv.Load(); err != nil {
		log.Print("WARNING: failed to load .env file")
	}

	// dns (db connection string) may contain secrets so is not appropriate for a config file
	// instead this is persisted via the environment or .env which is never committed to SCM
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://localhost:5432/auth"
	}

	flag.IntVar(&cfg.Port, "port", 4000, "listen port")
	flag.StringVar(&cfg.Host, "host", "localhost", "listen host")
	flag.StringVar(&cfg.Env, "env", api.Development, "application environment")
	flag.StringVar(&dsn, "dsn", dsn, "database connection string")
	flag.StringVar(&cfg.Auth.Secret, "jwt-secret", cfg.Auth.Secret, "jwt secret (default overridden by config)")
	flag.StringVar(&cfg.Auth.Issuer, "jwt-iss", cfg.Auth.Issuer, "jwt issuer")
	flag.StringVar(&cfg.Auth.Audience, "jwt-aud", cfg.Auth.Audience, "jwt audience")
	flag.StringVar(&cfg.Auth.Cookie.Domain, "cookie", cfg.Auth.Cookie.Domain, "cookie domain")

	showVersion := flag.Bool("version", false, "program version")
	writeConfig := flag.Bool("write-config", false, "write current config to file")

	flag.Parse()

	return *showVersion, *writeConfig, dsn
}

func main() {
	cfg, err := api.NewConfigFromFile(configName, version)
	if err != nil {
		log.Fatalln(err)
	}
	showVersion, writeConfig, dsn := setOptions(cfg)

	if showVersion {
		exe := pathlib.NewPath(os.Args[0])
		fmt.Printf("%s v%s\n", exe.Name(), version)
		os.Exit(0)
	}

	if writeConfig {
		if outfile, err := cfg.Write(configName); err != nil {
			fmt.Printf("Error: %s - %v", configName, err)
			os.Exit(1)
		} else {
			fmt.Printf("Sucessfully saved config to %s", outfile)
			os.Exit(0)
		}
	}

	app := api.NewApp(cfg, "main ", "server", version)

	app.DB, err = db.NewPG(dsn)
	if err != nil {
		log.Fatal("failed to open db:", err)
	}
	defer app.DB.Close()
	log.Print(app.StartServer())
}
