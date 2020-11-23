package drive

import (
	"fmt"

	"github.com/adam-lavrik/go-imath/ix"
	"google.golang.org/api/drive/v3"
)

// List returns list of files contained in drive.
func (s *Service) List(limit int) ([]*drive.File, error) {
	var fileList []*drive.File
	nextPageToken := ""
	for {
		q := s.drive.Files.List().OrderBy("modifiedTime desc")
		if nextPageToken != "" {
			q = q.PageToken(nextPageToken)
		}
		r, err := q.Do()
		if err != nil {
			fmt.Printf("An error occurred: %v\n", err)
			return fileList, err
		}
		fileList = append(fileList, r.Files...)
		nextPageToken = r.NextPageToken
		if nextPageToken == "" || len(fileList) >= limit {
			break
		}
	}
	return fileList[:ix.Min(limit, len(fileList))], nil
}
