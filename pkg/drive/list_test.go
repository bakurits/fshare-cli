package drive

import (
	"fileshare/pkg/auth"
	"fileshare/pkg/testutils"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestService_List(t *testing.T) {
	is := assert.New(t)
	client, err := auth.GetHTTPClient(testutils.RootDir() + "/credentials")
	if err != nil {
		log.Fatalf("Unable to retrieve http client: %v", err)
	}

	srv, err := NewService(client)
	is.NoError(err)
	files := srv.List()
	fmt.Println(files)
}
