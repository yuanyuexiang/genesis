package models

import (
	"crypto/sha1"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	_ "reflect"
	"sort"
	"strings"
)

type Message struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   int64
	MsgType      CDATA
	Content      CDATA
}

type CDATA struct {
	Text string `xml:",cdata"`
}

func init() {
}

func PushMessage(m *Message) (id int64, err error) {

	//m.Content = "123"
	//m.FromUserName = "123453453"
	m.MsgType.Text = "12123122242342342"
	b, _ := xml.MarshalIndent(m, "", "    ")
	fmt.Println(string(b))
	return
}

//
func CheckMessageInterface(signature, timestamp, nonce, echostr string) (err error) {

	if signature != makeSignature(timestamp, nonce) {
		err = errors.New("CHECK FAILED")
	}
	return err
}

//生成签名
func makeSignature(timestamp, nonce string) string {

	token := beego.AppConfig.String("token")
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}
