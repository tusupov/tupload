package handle

import (
	"net/http"
	"encoding/json"
	"github.com/tusupov/tupload"
)

func UploadImage(w http.ResponseWriter, r *http.Request, rootPath string) {

	imageWidth        := uint(100)
	imageHeight       := uint(100)
	imageFolder       := "/images/"
	imageResizeFolder := "/images/resize/"

	if r.Method == "POST" {

		result := NewResult()

		//max upload file 100 MB
		r.ParseMultipartForm(100<<20) //100 mb

		if r.MultipartForm != nil {

			//from file
			for _, file := range r.MultipartForm.File["image"] {

				if fileName, err := tupload.ImageMultipart(*file, rootPath + imageFolder); err != nil {
					result.AddError(file.Filename, err)
				} else {
					result.Add(fileName, "")
				}

			}

			//from url
			for _, url := range r.MultipartForm.Value["image"] {

				if fileName, err := tupload.ImageUrl(url, rootPath + imageFolder); err != nil {
					result.AddError(fileName, err)
				} else {
					result.Add(fileName, "")
				}
			}

		} else {

			//from json
			var bodyText map[string]interface{}
			json.NewDecoder(r.Body).Decode(&bodyText)

			if imageMap, prs := bodyText["image"]; prs {

				if image, ok := imageMap.(string); ok {
					if fileName, err := tupload.ImageBase64(image, rootPath + imageFolder); err != nil {
						result.AddError("", err)
					} else {
						result.Add(fileName, "")
					}
				}

				if images, ok := imageMap.([]string); ok {
					for _, image := range images {
						if fileName, err := tupload.ImageBase64(image, rootPath + imageFolder); err != nil {
							result.AddError("", err)
						} else {
							result.Add(fileName, "")
						}
					}
				}

			}

		}

		//create thumbnail for all success image
		for k, v := range result.Results {

			if len(v.Err) == 0 && len(v.Image) > 0 {

				thumbFileName, errThumb := tupload.ImageThumbnail(
					imageWidth,
					imageHeight,
					rootPath + imageFolder + v.Image,
					rootPath + imageResizeFolder,
					tupload.NearestNeighbor,
				)

				result.Results[k].Image = imageFolder + v.Image
				if errThumb == nil {
					result.Results[k].Thumb = imageResizeFolder + thumbFileName
				}

			}

		}

		json.NewEncoder(w).Encode(result)

	}

}
