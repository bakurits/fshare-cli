package mailstore

import (
	"bufio"
	"fmt"
	"os"
)

// ReadMails reads a whole file into memory and returns a slice of its lines.
func ReadMails(path string) ([]string, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// WriteLines writes the lines to the given file.
func WriteMail(mail string, path string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	_, _ = fmt.Fprintln(w, mail)
	return w.Flush()
}
