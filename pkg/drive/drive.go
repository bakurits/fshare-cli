package drive

import (
	"io"
	"log"
	"net/http"

	"github.com/bakurits/fileshare/pkg/auth"

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
func NewService(client *http.Client) (*Service, error) {
	srv, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return &Service{}, errors.Wrap(err, "unable to retrieve Drive client")
	}
	return &Service{
		drive: srv,
	}, nil
}

// CreateDir creates a directory at google drive
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

// CreateFile creates a file
func (s *Service) CreateFile(name string, mimeType string, content io.Reader, parentID string) (*drive.File, error) {
	f := &drive.File{
		MimeType: mimeType,
		Name:     name,
		Parents:  []string{parentID},
	}
	file, err := s.drive.Files.Create(f).Media(content).Do()

	if err != nil {
		log.Println("Could not create file: " + err.Error())
		return nil, err
	}

	return file, nil
}

// Authorize get a service object from credentials json file
func Authorize(credentials string) (*Service, error) {
	client, err := auth.GetHTTPClient(credentials)

	if err != nil {
		return &Service{}, errors.Wrap(err, "unable to get client from drive")
	}

	service, err := NewService(client)

	if err != nil {
		return &Service{}, errors.Wrap(err, "unable to get service drive")
	}

	return service, nil
}
