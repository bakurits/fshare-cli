package drive

import (
	"testing"

	"github.com/bakurits/fileshare/pkg/auth"

	"github.com/stretchr/testify/assert"
)

func TestService_Download(t *testing.T) {
	is := assert.New(t)

	conf, err := getTestConfig()
	is.NoError(err)

	authConf := auth.GetConfig(conf)
	client, err := authConf.ClientFromTokenFile(conf.TokenPath)
	is.NoError(err)

	srv, err := NewService(client.Client)
	is.NoError(err)

	f, err := srv.Get("The-Go-Programming-Language.pdf")
	is.NoError(err)

	err = srv.Download(f)
	is.NoError(err)
}
