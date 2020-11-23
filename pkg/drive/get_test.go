package drive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_Get(t *testing.T) {
	is := assert.New(t)

	client, err := getTestClient()
	is.NoError(err)

	srv, err := NewService(client.Client)
	is.NoError(err)

	f, err := srv.Get("შესარჩევი სრული")
	is.NoError(err)

	is.NotEmpty(f.Name)
}
