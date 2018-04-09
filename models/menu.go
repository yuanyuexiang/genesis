package models

import (
	"bufio"
	"encoding/json"
	"errors"
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

// ResponseWechatMenu ResponseWechatMenu
type ResponseWechatMenu struct {
	Menu WechatMenu `json:"menu"`
}

// WechatMenu WechatMenu
type WechatMenu struct {
	Buttons []WechatButton `json:"button"`
}

// WechatButton WechatButton
type WechatButton struct {
	Type      string         `json:"type,omitempty"`
	Name      string         `json:"name,omitempty"`
	Key       string         `json:"key,omitempty"`
	URL       string         `json:"url,omitempty"`
	AppID     string         `json:"appid,omitempty"`
	PagePath  string         `json:"pagepath,omitempty"`
	SubButton []WechatButton `json:"sub_button,omitempty"`
}

// LocalMenu Menu
type LocalMenu struct {
	Menu Menu `json:"Menu"`
}

// Menu Menu
type Menu struct {
	Buttons []Button `json:"Buttons"`
}

// Button WechatButton
type Button struct {
	ID        int
	Type      string `json:"Type,omitempty"`
	Name      string `json:"Name,omitempty"`
	Key       string `json:"Key,omitempty"`
	URL       string `json:"URL,omitempty"`
	AppID     string `json:"AppID,omitempty"`
	PagePath  string `json:"PagePath,omitempty"`
	ParentID  int
	SubButton []Button `json:"SubButton,omitempty"`
}

// CreateMenu AddCreateMenu2Menu
func CreateMenu(localMenu *LocalMenu) (err error) {

	bytes, _ := json.Marshal(localMenu)
	fmt.Println("-----------------localMenu--", string(bytes))

	responseWechatMenu := ResponseWechatMenu{}
	wechatButtons := []WechatButton{}
	buttons := localMenu.Menu.Buttons
	for _, button := range buttons {
		if button.ParentID == -1 {
			wechatButton := WechatButton{
				Type:     button.Type,
				Name:     button.Name,
				Key:      button.Key,
				URL:      button.URL,
				AppID:    button.AppID,
				PagePath: button.PagePath}

			for _, subButton := range buttons {

				if subButton.ParentID == button.ID {
					subWechatButton := WechatButton{
						Type:     subButton.Type,
						Name:     subButton.Name,
						Key:      subButton.Key,
						URL:      subButton.URL,
						AppID:    subButton.AppID,
						PagePath: subButton.PagePath}
					wechatButton.SubButton = append(wechatButton.SubButton, subWechatButton)
				}
			}
			wechatButtons = append(wechatButtons, wechatButton)
		}
	}
	responseWechatMenu.Menu.Buttons = wechatButtons
	bytes, _ = json.Marshal(responseWechatMenu)
	fmt.Println("-----------------wechatButtons--", string(bytes))

	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	strURL := createMenuURL + accessToken
	postData, err := json.Marshal(responseWechatMenu.Menu)
	if err != nil {
		return
	}
	data := map[string]interface{}{}
	body, err := post(strURL, postData)

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	errcode := data["errcode"].(float64)
	errmsg := data["errmsg"].(string)
	if errcode != 0 {
		err = errors.New(errmsg)
	}
	fmt.Println("------------------", errcode, errmsg)
	return
}

// GetMenu insert a new Menu into database and returns
// last inserted Id on success.
func GetMenu() (localMenu *LocalMenu, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var responseWechatMenu *ResponseWechatMenu
	strURL := getMenuURL + accessToken
	body, err := get(strURL)
	err = json.Unmarshal(body, &responseWechatMenu)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	localMenu = &LocalMenu{}
	wechatButtons := responseWechatMenu.Menu.Buttons
	//buttons := localMenu.Menu.Buttons
	buttons := []Button{}
	index := 3
	for i, wechatButton := range wechatButtons {

		buttons = append(buttons, Button{
			ID:       i,
			ParentID: -1,
			Type:     wechatButton.Type,
			Name:     wechatButton.Name,
			Key:      wechatButton.Key,
			URL:      wechatButton.URL,
			AppID:    wechatButton.AppID,
			PagePath: wechatButton.PagePath})

		if wechatButton.SubButton != nil {
			for _, wechatButton := range wechatButton.SubButton {
				index++
				buttons = append(buttons, Button{
					ID:       index,
					ParentID: i,
					Type:     wechatButton.Type,
					Name:     wechatButton.Name,
					Key:      wechatButton.Key,
					URL:      wechatButton.URL,
					AppID:    wechatButton.AppID,
					PagePath: wechatButton.PagePath})
			}
		}
	}
	localMenu.Menu.Buttons = buttons
	bytes, _ := json.Marshal(buttons)
	fmt.Println("-----------------mediaInfo--", string(bytes))
	return localMenu, nil
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
func CreateMenu2(requestData interface{}) (data map[string]interface{}, err error) {
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
