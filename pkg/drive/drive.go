package drive

import (
	"github.com/bakurits/fileshare/pkg/auth"
	"io"
	"log"
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

// Service google drive API wrapper
type Service struct {
	Drive *drive.Service
}

// NewService returns new service instance
func NewService(client *http.Client) (Service, error) {
	srv, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return Service{}, errors.Wrap(err, "unable to retrieve Drive client")
	}
	return Service{
		Drive: srv,
	}, nil
}

// create a folder in drive
func CreateDir(service *drive.Service, name string, parentId string) (*drive.File, error) {
	d := &drive.File{
		Name:     name,
		MimeType: "application/vnd.google-apps.folder",
		Parents:  []string{parentId},
	}

	file, err := service.Files.Create(d).Do()

	if err != nil {
		log.Println("Could not create dir: " + err.Error())
		return nil, err
	}

	return file, nil
}

// create a file
func CreateFile(service *drive.Service, name string, mimeType string, content io.Reader, parentId string) (*drive.File, error) {
	f := &drive.File{
		MimeType: mimeType,
		Name:     name,
		Parents:  []string{parentId},
	}
	file, err := service.Files.Create(f).Media(content).Do()

	if err != nil {
		log.Println("Could not create file: " + err.Error())
		return nil, err
	}

	return file, nil
}

// get a service object from credentials json file
func Authorize(credentials string) (Service, error) {
	client, err := auth.GetHTTPClient(credentials)

	if err != nil {
		return Service{}, errors.Wrap(err, "unable to get client from drive")
	}

	service, err := NewService(client)

	if err != nil {
		return Service{}, errors.Wrap(err, "unable to get service drive")
	}

	return service, nil
}
