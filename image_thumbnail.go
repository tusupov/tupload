package tupload

import (
	"os"
	"github.com/nfnt/resize"
	"net/http"
	"image"
	"path/filepath"
	"image/jpeg"
	"image/png"
	"image/gif"
	"strings"
)

const (

	// Nearest-neighbor interpolation
	NearestNeighbor resize.InterpolationFunction = iota

	// Bilinear interpolation
	Bilinear

	// Bicubic interpolation (with cubic hermite spline)
	Bicubic

	// Mitchell-Netravali interpolation
	MitchellNetravali

	// Lanczos interpolation (a=2)
	Lanczos2

	// Lanczos interpolation (a=3)
	Lanczos3

)

func ImageThumbnail(maxWidth, maxHeight uint, imgFilePath, path string, interp resize.InterpolationFunction) (string, error) {

	// check file
	file, errFile := os.Open(imgFilePath)
	if errFile != nil {
		return "", errFile
	}
	defer file.Close()


	// check file to image
	buffer := make([]byte, 512)
	_, errRead := file.Read(buffer)
	if errRead != nil {
		return "", errRead
	}
	file.Seek(0, 0)

	contentType := http.DetectContentType(buffer)

	isImage, _ := isImage(contentType)
	if !isImage {
		return "", errorNoImage
	}


	// decode image
	img, _, errDecode := image.Decode(file)
	if errDecode != nil {
		return "", errDecode
	}


	// resize image
	newImg := resize.Thumbnail(maxWidth, maxHeight, img, interp)


	fileName := filepath.Base(imgFilePath)


	// create folders
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", err
	}


	// create new file to copy
	f, err := os.OpenFile(path + "/" + fileName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return "", err
	}
	defer f.Close()


	// copy resized image to new file
	switch strings.TrimPrefix(filepath.Ext(fileName), ".") {

	case "jpg":
		err := jpeg.Encode(f, newImg, nil)
		if err != nil {
			return "", err
		}

	case "png":
		err := png.Encode(f, newImg)
		if err != nil {
			return "", err
		}

	case "gif":
		err := gif.Encode(f, newImg, nil)
		if err != nil {
			return "", err
		}

	default:
		os.Remove(f.Name())
		return "", errorNotFoundExt

	}

	return fileName, nil

}
