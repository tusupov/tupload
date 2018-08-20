package tupload

import "errors"

var (
	errorNoImage 		= errors.New("Not an image")
	errorBase64 		= errors.New("Not a base64 data")
	errorNotFoundExt	= errors.New("Could not determine type")
	errorSize			= errors.New("Incorrect image size")
)
