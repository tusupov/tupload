package tupload

import (
	"mime/multipart"
	"net/http"
	"encoding/base64"
	"regexp"
	"bytes"
)


func ImageMultipart(header multipart.FileHeader, path string) (fileName string, err error) {

	isImage, ext := isImage(header.Header.Get("Content-Type"))
	if !isImage {
		return "", errorNoImage
	}

	return moveFile(header, path, ext)

}

func ImageUrl(url string, path string) (fileName string, err error) {

	var response *http.Response
	response, err = http.Get(url)
	if err != nil {
		return
	}

	isImage, ext := isImage(response.Header.Get("Content-Type"))
	if !isImage {
		return "", errorNoImage
	}

	return saveFile(response.Body, path, ext)

}

func ImageBase64(b64 string, path string) (fileName string, err error) {

	//create new regex checker
	var r *regexp.Regexp
	r, err = regexp.Compile("^data:(.+);base64,(.+)")
	if err != nil {
		return
	}

	list := r.FindStringSubmatch(b64)

	//is not valid
	if len(list) < 3 {
		return "", errorBase64
	}

	contentType := list[1]

	isImage, ext := isImage(contentType)
	if !isImage {
		return "", errorNoImage
	}

	b64 = list[2]

	var dec []byte
	dec, err = base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return
	}

	body := bytes.NewReader(dec)

	return saveFile(body, path, ext)

}

func isImage(contentType string) (bool bool, extension string) {
	return doMatch(contentType, imageType)
}
