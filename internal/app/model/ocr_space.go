package model

import (
	"bytes"
	"encoding/json"
	"github.com/JetBrainer/BackOCRService/internal/app/apiserver"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type Config struct {
	ApiKey   string `toml:"apikey"`
	Language string `toml:"language"`
	Url		 string `toml:"url"`
	DBUrl	 string `toml:"database_url"`
	HttpPort string `toml:"port"`
}

func (c Config) ParseForm(fileURL string) (apiserver.OCRText, error){
	var results apiserver.OCRText
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

func (c Config) ParseFromBase64(baseString string) (apiserver.OCRText, error){
	var results apiserver.OCRText
	resp, err := http.PostForm(c.Url, url.Values{
		"base64Image":					{baseString},
		"language":						{c.Language},
		"apikey":						{c.ApiKey},
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

	err = json.Unmarshal(body,&results)
	if err != nil{
		return results,err
	}

	return results,nil
}

func (c Config)ParseFromLocal (localPath string) (apiserver.OCRText,error){
	var results apiserver.OCRText
	params := map[string]string{
		"language":					c.Language,
		"apikey":					c.ApiKey,
		"isOverlayRequired":		"true",
		"isSearchablePdfTextLayer":	"true",
		"scale":					"true",
	}
	file, err := os.Open(localPath)
	if err != nil{
		return results, err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file",filepath.Base(localPath))
	if err != nil{
		return results,err
	}

	_, err = io.Copy(part,file)
	if err != nil{
		return results,err
	}

	for key, val := range params{
		_ = writer.WriteField(key,val)
	}
	err = writer.Close()
	if err != nil{
		return results,err
	}

	req, err := http.NewRequest(http.MethodPost,c.Url,body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil{
		log.Fatal(err)
	} else{
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(response.Body)
		if err != nil{
			return results,err
		}
		response.Body.Close()
		err = json.Unmarshal(body.Bytes(),&results)
		if err != nil{
			return results,err
		}
	}
	return results, nil
}

