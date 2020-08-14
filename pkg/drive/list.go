package drive

import (
	"google.golang.org/api/drive/v3"
)

// List returns list of files contained in drive
func (s Service) List() []*drive.File {
	r, err := s.drive.Files.List().OrderBy("createdTime desc").PageSize(20).
		Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		return []*drive.File{}
	}
	return r.Files
}
