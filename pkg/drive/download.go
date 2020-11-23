package drive

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/pkg/errors"
	"google.golang.org/api/drive/v3"
)

// Download downloads file from drive.
func (s *Service) Download(f *drive.File) error {
	resp, err := s.drive.Files.Get(f.Id).Download()
	if err != nil {
		return errors.Wrap(err, "download error")
	}
	defer func() { _ = resp.Body.Close() }()

	if err = downloadTmp(resp, f.Name); err != nil {
		return err
	}

	if err = os.Rename(f.Name+".tmp", f.Name); err != nil {
		return errors.Wrap(err, "can't store file")
	}
	fmt.Println("")
	return nil
}

// WriteCounter counts the number of bytes written to it. It implements to the io.Writer interface
// and we can pass this into io.TeeReader() which will report progress on each write cycle.
type writeCounter struct {
	Total uint64
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.printProgress()
	return n, nil
}

func (wc writeCounter) printProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

func downloadTmp(resp *http.Response, fileName string) error {
	out, err := os.Create(fileName + ".tmp")
	if err != nil {
		return errors.Wrap(err, "can't create file")
	}
	defer func() { _ = out.Close() }()

	counter := &writeCounter{}

	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return errors.Wrap(err, "can't copy data")
	}
	return nil
}
