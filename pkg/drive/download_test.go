package drive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_Download(t *testing.T) {
	is := assert.New(t)
	client, err := getTestClient()
	is.NoError(err)

	srv, err := NewService(client.Client)
	is.NoError(err)

	f, err := srv.Get("The-Go-Programming-Language.pdf")
	is.NoError(err)

	err = srv.Download(f)
	is.NoError(err)
}
