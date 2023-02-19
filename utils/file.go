package utils

import (
	"errors"
	"mime/multipart"
)

func CheckContentType(file *multipart.FileHeader, contentTypes ...string) error {
	if len(contentTypes) > 0 {
		for _, value := range contentTypes {
			contentType := file.Header.Get("Content-Type")
			if contentType == "image/"+value {
				return nil
			}
		}

		return errors.New("not allowed file type")
	} else {
		return errors.New("not found content type to be checking")
	}
}
