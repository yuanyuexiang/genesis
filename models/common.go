package models

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

// ReturnData ReturnData
type ReturnData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// GetReturnData GetReturnData
func GetReturnData(code int, message string, data interface{}) (rd *ReturnData) {
	rd = &ReturnData{code, message, data}
	return
}

//CDATA CDATA
type CDATA struct {
	Text string `xml:",cdata"`
}

//ErrorResponse ErrorResponse
type ErrorResponse struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}

func get(url string) (data []byte, err error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	if response.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
		err = errors.New("wechat server error")
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	bodystr := string(body)
	fmt.Println(bodystr)
	data = body
	return
}

func post(url string, postData []byte) (data []byte, err error) {
	request, err := http.NewRequest("POST", url, strings.NewReader(string(postData)))
	if err != nil {
		return
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	if response.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
		err = errors.New("wechat server error")
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	bodystr := string(body)
	fmt.Println(bodystr)
	data = body
	return
}

func postFile(url, description, filePath string) (data []byte, err error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()
	fw, err := w.CreateFormFile("media", filePath)
	if err != nil {
		return
	}
	if _, err = io.Copy(fw, f); err != nil {
		return
	}
	if description != "" {
		w.WriteField("description", description)
	}
	w.Close()
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	bodystr := string(body)
	fmt.Println(bodystr)
	data = body
	return
}
