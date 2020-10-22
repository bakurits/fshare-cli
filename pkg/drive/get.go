package drive

import (
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/api/drive/v3"
)

// Get finds file with given name on drive
// if there are multiple files with given name method returns latest modified
func (s *Service) Get(name string) (*drive.File, error) {
	r, err := s.drive.Files.List().Q(fmt.Sprintf("name = '%s'", name)).PageSize(1).
		Fields("nextPageToken, files(id, name, mimeType)").Do()
	if err != nil {
		return nil, errors.Wrap(err, "error while filtering")
	}
	if len(r.Files) == 0 {
		return nil, errors.New("file not found")
	}
	return r.Files[0], nil
}
