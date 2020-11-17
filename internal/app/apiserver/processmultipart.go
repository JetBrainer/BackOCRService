package apiserver

import (
	"fmt"
	"io"
	"strings"
)


func prepareUploadBody(fileReader io.Reader, lang, apikey string) (string,io.Reader){
	boundary := "MultiPartBoundary"

	fieldFormat := "--%s\r\nContent-Disposition: form-data; name=\"%s\"\r\n\r\n%s\r\n"
	tokenPart := fmt.Sprintf(fieldFormat, boundary, "token", apikey)

	fieldFormat2 := "--%s\r\nContent-Disposition: form-data; name=\"%s\"\r\n\r\n%s\r\n"
	langPart := fmt.Sprintf(fieldFormat2, boundary, "language", lang)

	fieldFormat3 := "--%s\r\nContent-Disposition: form-data; name=\"%s\"\r\n\r\n%s\r\n"
	overlay := fmt.Sprintf(fieldFormat3, boundary, "isOverlayRequired", "true")

	fieldFormat4 := "--%s\r\nContent-Disposition: form-data; name=\"%s\"\r\n\r\n%s\r\n"
	scale := fmt.Sprintf(fieldFormat4, boundary, "scale", "true")

	//fileName := "build.jpeg"
	//fileHeader := "Content-type: application/octet-stream"
	//fileFormat := "--%s\r\nContent-Disposition: form-data; name=\"file\"; filename=\"%s\"\r\n%s\r\n\r\n"
	//filePart := fmt.Sprintf(fileFormat, boundary, fileName, fileHeader)

	bodyTop := fmt.Sprintf("%s%s%s%s", langPart, tokenPart,overlay,scale)

	body := io.MultiReader(strings.NewReader(bodyTop), fileReader)

	contentType := fmt.Sprintf("multipart/form-data; boundary=%s", boundary)

	return contentType, body
}