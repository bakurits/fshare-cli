package drive

import (
	"fmt"
	"github.com/bakurits/fileshare/pkg/auth"
	"github.com/bakurits/fileshare/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestService_Get(t *testing.T) {
	is := assert.New(t)
	client, err := auth.GetHTTPClient(testutils.RootDir() + "/credentials")
	if err != nil {
		log.Fatalf("Unable to retrieve http client: %v", err)
	}

	srv, err := NewService(client)
	is.NoError(err)
	f, err := srv.Get("შესარჩევი სრული")
	fmt.Println(f)
}
