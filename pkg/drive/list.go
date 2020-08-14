package drive

import (
	"google.golang.org/api/drive/v3"
)

func (s Service) List() []*drive.File {
	r, err := s.drive.Files.List().OrderBy("createdTime desc").PageSize(20).
		Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		return []*drive.File{}
	}
	return r.Files
}
