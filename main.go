package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kgaughan/mocha/internal/models"
	flag "github.com/spf13/pflag"
	"gorm.io/gorm"
)

var Version = "built incorrectly"

var (
	help        = flag.BoolP("help", "h", false, "show help")
	showVersion = flag.Bool("version", false, "show version")
	spider      = flag.Bool("spider", false, "run in spider mode")
	config      string
)

func init() {
	flag.StringVar(&config, "config", path.Join(getConfigRoot(), "mocha.toml"), "path to configuration file")

	flag.Usage = func() {
		name := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "%s - Generates an aggregated site from a set of feeds.\n\n", name)
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nVersion:\n  %v\n", Version)
	}
}

func getConfigRoot() string {
	if cfgRoot, ok := os.LookupEnv("XDG_CONFIG_HOME"); ok {
		return cfgRoot
	}
	return path.Join(os.Getenv("HOME"), ".config")
}

func main() {
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}
	if *showVersion {
		fmt.Println(Version)
		os.Exit(0)
	}

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal("gorm.Open:", err)
	}
	db.AutoMigrate(&models.Article{})

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World!")
	})
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
