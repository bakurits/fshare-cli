package drive

import (
	"fmt"
	"log"
	"testing"

	"github.com/bakurits/fileshare/pkg/auth"
	"github.com/bakurits/fileshare/pkg/testutils"

	"github.com/stretchr/testify/assert"
)

func TestService_List(t *testing.T) {
	is := assert.New(t)
	client, err := auth.GetHTTPClient(testutils.RootDir() + "/credentials")
	if err != nil {
		log.Fatalf("Unable to retrieve http client: %v", err)
	}

	srv, err := NewService(client)
	is.NoError(err)
	files := srv
	fmt.Println(files)
}
