package drive

import (
	"github.com/bakurits/fileshare/pkg/auth"
	"github.com/bakurits/fileshare/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Download(t *testing.T) {
	is := assert.New(t)
	client, err := auth.GetHTTPClient(testutils.RootDir() + "/credentials")
	is.NoError(err)

	srv, err := NewService(client)
	is.NoError(err)

	f, err := srv.Get("შესარჩევი სრული")
	is.NoError(err)

	err = srv.Download(f)
	is.NoError(err)
}
