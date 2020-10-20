package drive

import (
	"github.com/pkg/errors"
	"google.golang.org/api/drive/v3"
	"io"
	"os"
)

// Download downloads file from drive
func (s *Service) Download(f *drive.File) error {
	resp, err := s.drive.Files.Export(f.Id, f.MimeType).Fields().Download()
	if err != nil {
		return errors.Wrap(err, "download error")
	}
	defer func() { _ = resp.Body.Close() }()

	out, err := os.Create(f.Name)
	if err != nil {
		return errors.Wrap(err, "can't create file")
	}
	defer func() { _ = out.Close() }()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return errors.Wrap(err, "can't copy data")
	}
	return nil
}
