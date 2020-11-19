package drive

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_List(t *testing.T) {
	is := assert.New(t)

	client, err := getTestClient()
	is.NoError(err)

	srv, err := NewService(client.Client)
	is.NoError(err)
	files := srv

	fileList, err := files.List(100)
	if err != nil {
		return
	}

	for _, file := range fileList {
		fmt.Println(file.Name)
	}
}
