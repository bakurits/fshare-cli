package main

import (
	"github.com/bakurits/fileshare/pkg/auth"
	"github.com/bakurits/fileshare/pkg/cfg"
	"github.com/bakurits/fileshare/pkg/webapp/db"
	"github.com/bakurits/fileshare/pkg/webapp/server"
	"log"
	"net/http"
)

type Config struct {
	StaticFileDir  string
	CredentialsDir string

	ConnectionString string
	DBDialect        string

	Server string
	Port   string

	GoogleCredentialsPath string
	GoogleCredentials     cfg.GoogleCredentials
}

func main() {
	var conf Config
	err := cfg.GetConfig(&conf)
	if err != nil {
		log.Fatal(err)
	}

	web, err := cfg.LoadGoogleCredentials(conf.GoogleCredentialsPath, cfg.CredentialTypeWeb)
	if err != nil {
		log.Fatal(err)
	}
	conf.GoogleCredentials = web

	repository, err := db.NewRepository(conf.DBDialect, conf.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	s := &server.Server{
		AuthConfig:    auth.GetConfig(conf.GoogleCredentials.ClientID, conf.GoogleCredentials.ClientSecret, conf.Server+conf.Port+"/auth"),
		Repository:    repository,
		StaticFileDir: conf.StaticFileDir,
	}
	s.Init()

	err = http.ListenAndServe(conf.Port, s)
	log.Fatal(err)
}
