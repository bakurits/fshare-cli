package main

import (
	"github.com/bakurits/fileshare/pkg/auth"
	"github.com/bakurits/fileshare/pkg/cfg"
	"github.com/bakurits/fileshare/pkg/webapp/db"
	"github.com/bakurits/fileshare/pkg/webapp/server"
	"os"

	"log"
	"net/http"
)

func main() {

	var conf auth.WebConfig
	err := cfg.GetConfig(os.Getenv("config.json"), &conf)
	if err != nil {
		log.Fatal(err)
	}

	web, err := auth.LoadGoogleCredentials(conf.GoogleCredentialsPath)
	if err != nil {
		log.Fatal(err)
	}
	conf.GoogleCredentials = web

	repository, err := db.NewRepository(conf.DBDialect, conf.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	s := &server.Server{
		AuthConfig:    auth.GetWebConfig(conf),
		Repository:    repository,
		StaticFileDir: conf.StaticFileDir,
	}
	s.Init()

	err = http.ListenAndServe(conf.Port, s)
	log.Fatal(err)
}
