package gotwitter

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type (
	UploadMedia struct {
		FileName string
		Body     *bytes.Buffer
		Writer   *multipart.Writer
	}

	UploadMediaResponse struct {
	}
)

func (u *UploadMedia) Upload(fileName string) error {

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	u.Writer = multipart.NewWriter(u.Body)
	part, err := u.Writer.CreateFormFile("media", fileName)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}
	u.Writer.Close()

	return nil
}

func (u *UploadMedia) EventMethod() string {
	return http.MethodPost
}

func (u *UploadMedia) EventName() string {
	return "upload_media"
}
func (u *UploadMedia) EventEndpoint() string {
	return EndpointUploadMedia
}
func (u *UploadMedia) EventHeader(oauthHeader string) map[string]string {

	headerMap := make(map[string]string)
	headerMap["Authorization"] = oauthHeader
	headerMap["Content-Type"] = u.Writer.FormDataContentType()

	return headerMap
}
