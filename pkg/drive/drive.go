package drive

import (
	"log"
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

// Service google drive API wrapper.
type Service struct {
	drive *drive.Service
}

// NewService returns new service instance.
func NewService(client *http.Client) (*Service, error) {
	srv, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return &Service{}, errors.Wrap(err, "unable to retrieve Drive client")
	}
	return &Service{
		drive: srv,
	}, nil
}

// CreateDir creates a directory at google drive.
func (s *Service) CreateDir(name string, parentID string) (*drive.File, error) {
	d := &drive.File{
		Name:     name,
		MimeType: "application/vnd.google-apps.folder",
		Parents:  []string{parentID},
	}

	file, err := s.drive.Files.Create(d).Do()

	if err != nil {
		log.Println("Could not create dir: " + err.Error())
		return nil, err
	}

	return file, nil
}
