package drive

import (
	"mime"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"google.golang.org/api/drive/v3"
)

func (s *Service) Upload(filePath string) (*drive.File, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("file does not exist " + filePath)
	}
	defer func() { _ = f.Close() }()
	_, fileName := filepath.Split(f.Name())

	tp := getFilePathMimeType(filePath)
	df := &drive.File{
		MimeType: tp,
		Name:     fileName,
		Parents:  []string{"root"},
	}
	file, err := s.drive.Files.Create(df).Media(f).Do()
	if err != nil {
		return nil, errors.Wrap(err, "could not create file")
	}
	return file, nil
}

func getFilePathMimeType(filePath string) string {
	ext := filepath.Ext(filePath)
	tp := mime.TypeByExtension(ext)
	return tp
}
