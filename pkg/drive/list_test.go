package drive

import (
	"fmt"
	"testing"

	"github.com/bakurits/fileshare/pkg/auth"
	"github.com/stretchr/testify/assert"
)

func TestService_List(t *testing.T) {
	is := assert.New(t)
	conf, err := getTestConfig()
	is.NoError(err)

	authConf := auth.GetConfig(conf)
	client, err := authConf.ClientFromTokenFile(conf.TokenPath)
	is.NoError(err)

	srv, err := NewService(client.Client)
	is.NoError(err)
	files := srv
	fmt.Println(files)
}
