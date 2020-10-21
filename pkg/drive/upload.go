package drive

import (
	"mime"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"google.golang.org/api/drive/v3"
)

func (s *Service) Upload(filepath string) (*drive.File, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, errors.New("file does not exist " + filepath)
	}
	defer func() { _ = f.Close() }()

	tp := getFilePathMimeType(filepath)
	df := &drive.File{
		MimeType: tp,
		Name:     f.Name(),
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
