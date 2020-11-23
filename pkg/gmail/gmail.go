package gmail

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type Service struct {
	service *gmail.Service
}

// NewService returns new service instance
func NewService(client *http.Client) (*Service, error) {
	srv, err := gmail.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return &Service{}, errors.Wrap(err, "unable to retrieve Drive client")
	}
	return &Service{
		service: srv,
	}, nil
}

func (s *Service) SendMessage(userID string, message gmail.Message) error {
	_, err := s.service.Users.Messages.Send(userID, &message).Do()
	return err
}

func CreateMessage(from string, to string, subject string, content string) gmail.Message {
	var message gmail.Message

	messageBody := []byte("From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		content)

	message.Raw = base64.StdEncoding.EncodeToString(messageBody)

	return message
}

func ChunkSplit(body string, limit int, end string) string {
	var charSlice []rune

	// push characters to slice
	for _, char := range body {
		charSlice = append(charSlice, char)
	}

	var result string = ""

	for len(charSlice) >= 1 {
		// convert slice/array back to string
		// but insert end at specified limit

		result = result + string(charSlice[:limit]) + end

		// discard the elements that were copied over to result
		charSlice = charSlice[limit:]

		// change the limit
		// to cater for the last few words in
		//
		if len(charSlice) < limit {
			limit = len(charSlice)
		}
	}

	return result
}

func randStr(strSize int, randType string) string {
	var dictionary string

	if randType == "alphanum" {
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "alpha" {
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "number" {
		dictionary = "0123456789"
	}

	var bytes = make([]byte, strSize)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

func CreateMessageWithAttachment(from string, to string, subject string, content string, fileDir string, fileName string) gmail.Message {
	var message gmail.Message

	// read file for attachment purpose
	fileBytes, err := ioutil.ReadFile(fileDir + fileName)
	if err != nil {
		log.Fatalf("Unable to read file for attachment: %v", err)
	}

	fileMIMEType := http.DetectContentType(fileBytes)

	fileData := base64.StdEncoding.EncodeToString(fileBytes)

	boundary := randStr(32, "alphanum")

	messageBody := []byte("Content-Type: multipart/mixed; boundary=" + boundary + " \n" +
		"MIME-Version: 1.0\n" +
		"to: " + to + "\n" +
		"from: " + from + "\n" +
		"subject: " + subject + "\n\n" +

		"--" + boundary + "\n" +
		"Content-Type: text/plain; charset=" + string('"') + "UTF-8" + string('"') + "\n" +
		"MIME-Version: 1.0\n" +
		"Content-Transfer-Encoding: 7bit\n\n" +
		content + "\n\n" +
		"--" + boundary + "\n" +

		"Content-Type: " + fileMIMEType + "; name=" + string('"') + fileName + string('"') + " \n" +
		"MIME-Version: 1.0\n" +
		"Content-Transfer-Encoding: base64\n" +
		"Content-Disposition: attachment; filename=" + string('"') + fileName + string('"') + " \n\n" +
		ChunkSplit(fileData, 76, "\n") +
		"--" + boundary + "--")

	// see https://godoc.org/google.golang.org/api/gmail/v1#Message on .Raw
	// use URLEncoding here !! StdEncoding will be rejected by Google API

	message.Raw = base64.URLEncoding.EncodeToString(messageBody)

	return message
}
