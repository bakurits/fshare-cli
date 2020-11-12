package main

import (
	"github.com/bakurits/fileshare/pkg/auth"
	"github.com/bakurits/fileshare/pkg/cfg"
	"github.com/bakurits/fileshare/pkg/webapp/db"
	"github.com/bakurits/fileshare/pkg/webapp/server"
	"log"
	"net/http"
	"os"
)

func main() {
	conf, err := cfg.GetConfig(os.Getenv("config"))
	if err != nil {
		log.Fatal(err)
	}
	repository, err := db.NewRepository(conf.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	s := &server.Server{
		AuthConfig:    auth.GetConfig(conf),
		Repository:    repository,
		StaticFileDir: conf.StaticFileDir,
	}
	s.Init()

	err = http.ListenAndServe(conf.Port, s)
	log.Fatal(err)
}
