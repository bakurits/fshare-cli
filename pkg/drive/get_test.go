package drive

import (
	"fmt"
	"github.com/bakurits/fileshare/pkg/auth"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Get(t *testing.T) {
	is := assert.New(t)
	conf, err := getTestConfig()
	is.NoError(err)

	authConf := auth.GetConfig(conf)
	client, err := authConf.ClientFromTokenFile(conf.TokenPath)
	is.NoError(err)

	srv, err := NewService(client.Client)
	is.NoError(err)

	f, err := srv.Get("შესარჩევი სრული")
	fmt.Println(f)
}
