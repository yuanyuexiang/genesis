package models

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const (
	getMenuURL               = "https://api.weixin.qq.com/cgi-bin/menu/get?access_token="
	deleteMenuURL            = "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token="
	createMenuURL            = "https://api.weixin.qq.com/cgi-bin/menu/create?access_token="
	addConditionalMenuURL    = "https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token="
	deleteConditionalMenuURL = "https://api.weixin.qq.com/cgi-bin/menu/delconditional?access_token="
)

// GetMenu insert a new Menu into database and returns
// last inserted Id on success.
func GetMenu() (data map[string]interface{}, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	strURL := getMenuURL + accessToken
	body, err := get(strURL)
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

// DeleteMenu deletes Menu by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMenu() (data map[string]interface{}, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	strURL := deleteMenuURL + accessToken
	body, err := get(strURL)
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

// CreateMenu deletes Menu by Id and returns error if
// the record to be deleted doesn't exist
func CreateMenu(requestData interface{}) (data map[string]interface{}, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	strURL := createMenuURL + accessToken
	postData, err := json.Marshal(requestData)
	if err != nil {
		return
	}
	body, err := post(strURL, postData)
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

// AddConditionalMenu deletes Menu by Id and returns error if
// the record to be deleted doesn't exist
func AddConditionalMenu(requestData interface{}) (data map[string]interface{}, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	strURL := addConditionalMenuURL + accessToken
	postData, err := json.Marshal(requestData)
	if err != nil {
		return
	}
	body, err := post(strURL, postData)
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

// DeleteConditionalMenu deletes Menu by Id and returns error if
// the record to be deleted doesn't exist
func DeleteConditionalMenu() (data map[string]interface{}, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	strURL := deleteConditionalMenuURL + accessToken
	body, err := get(strURL)
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func readFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("read file error\n")
		return "", err
	}

	defer file.Close()

	inputReader := bufio.NewReader(file)
	for {
		str, err := inputReader.ReadString('\n')
		if err == io.EOF {
			fmt.Printf("read eof")
			return "", err
		}
		if err != nil {
			return "", err
		}
		fmt.Printf("the input is:%s", str)
		return str, nil
	}
}
