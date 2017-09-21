package models

import (
	"encoding/json"
	"fmt"
)

func init() {

}

const (
	materialPictureAddMaterial = "https://api.weixin.qq.com/cgi-bin/material/add_material?"
)

/*
{
 "title":VIDEO_TITLE,
 "introduction":INTRODUCTION
}
*/

type MaterialInfo struct {
	MediaID string `json:"media_id"`
	URL     string `json:"url"`
}

//https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=ACCESS_TOKEN&type=TYPE
// AddImageFileToWechat insert a new File into database and returns
// last inserted Id on success.
func AddImageFileToWechat(filePath string) (mediaInfo MaterialInfo, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
	}

	strURL := materialPictureAddMaterial + "access_token=" + accessToken + "&type=image"

	desc := map[string]string{"title": "title", "introduction": "introduction"}

	description, err := json.Marshal(desc)

	if err != nil {
		return
	}
	body, err := postFile(strURL, string(description), filePath)

	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &mediaInfo)
	return
}
