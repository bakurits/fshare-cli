package drive

import (
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

// Service google drive API wrapper
type Service struct {
	drive *drive.Service
}

// NewService returns new service instance
func NewService(client *http.Client) (Service, error) {
	srv, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return Service{}, errors.Wrap(err, "unable to retrieve Drive client")
	}
	return Service{
		drive: srv,
	}, nil
}
