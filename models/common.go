package models

//CDATA CDATA
type CDATA struct {
	Text string `xml:",cdata"`
}

//ErrorResponse ErrorResponse
type ErrorResponse struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}
