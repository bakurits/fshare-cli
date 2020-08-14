package drive

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"net/http"
)

type Service struct {
	drive *drive.Service
}

func NewService(client *http.Client) (Service, error) {
	srv, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return Service{}, errors.Wrap(err, "unable to retrieve Drive client")
	}
	return Service{
		drive: srv,
	}, nil
}
