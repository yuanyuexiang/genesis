package models

type CDATA struct {
	Text string `xml:",cdata"`
}

type ErrorResponse struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}
