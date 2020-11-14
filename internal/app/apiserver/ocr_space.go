package apiserver

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

// Parse data from URL
func (c *Config) ParseFromURL(fileURL string) (OCRText, error){
	var results OCRText
	resp, err := http.PostForm(c.Url, url.Values{
		"url": 							{fileURL},
		"language": 					{c.Language},
		"apikey": 						{c.ApiKey},
		"isOverlayRequired":			{"true"},
		"isSearchablePdfHideTextLayer": {"true"},
		"scale":						{"true"},
	})
	if err != nil{
		return results, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return results, err
	}

	err = json.Unmarshal(body, &results)
	if err != nil{
		return results,err
	}

	return results, err
}

// Parse to OCR From Base64
func (c *Config) ParseFromBase64(baseString string, results *OCRText)  error{

	resp, err := http.PostForm(c.Url, url.Values{
		"base64Image":					{baseString},
		"language":						{c.Language},
		"apikey":						{c.ApiKey},
		"isOverlayRequired":			{"true"},
		"isSearchablePdfHideTextLayer": {"true"},
		"scale":						{"true"},
	})
	if err != nil{
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return err
	}

	err = json.Unmarshal(body,&results)
	if err != nil{
		return err
	}

	return nil
}

// Parse Local Files
func (c *Config)ParseFromPost(r *http.Request,results *OCRText) error{
	var err error
	params := map[string]string{
		"language":					c.Language,
		"apikey":					c.ApiKey,
		"isOverlayRequired":		"true",
		"scale":					"true",
	}

	buf := &bytes.Buffer{}

	writer := multipart.NewWriter(buf)


	for key, val := range params{
		_ = writer.WriteField(key,val)
	}

	if _,err = io.Copy(buf, r.Body); err != nil{
		log.Println("Error io Copying our buffer", err)
	}

	err = writer.Close()
	if err != nil{
		log.Println("Writer error", err)
		return err
	}

	req, err := http.NewRequest(http.MethodPost,c.Url,buf)
	if err != nil{
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil{
		log.Fatal(err)
	} else{
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(response.Body)
		if err != nil{
			return err
		}
		response.Body.Close()
		err = json.Unmarshal(body.Bytes(),&results)
		if err != nil{
			return err
		}
	}
	return nil
}

func (c Config) ParseFromLocal(localPath string) (OCRText, error) {
	var results OCRText
	params := map[string]string{
		"language":                     c.Language,
		"apikey":                       c.ApiKey,
		"isOverlayRequired":            "true",
		"scale": 						"true",
	}

	file, err := os.Open(localPath)
	if err != nil {
		return results, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base(localPath))
	if err != nil {
		return results, err
	}
	_, err = io.Copy(part, file)
	if err != nil{
		return results, err
	}
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return results, err
	}

	req, err := http.NewRequest(http.MethodPost, c.Url, body)
	if err != nil{
		return results,err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		response.Body.Close()
		err = json.Unmarshal(body.Bytes(), &results)
		if err != nil {
			return results, err
		}
	}

	return results, nil
}

func (ocr *OCRText) JustText() string{
	text := ""
	if ocr.IsErroredOnProcessing{
		for _, page := range ocr.ErrorMessage{
			text += page
		}
	} else {
		for _, page := range ocr.ParsedResults{
			text += page.ParsedText
		}
	}
	return text
}

