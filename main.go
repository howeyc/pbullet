// pbullet project main.go
package main

import (
	"fmt"
	"net/http"
	"net/url"
)

var pushUrl string

type PushDev struct {
	Id string
}

func SetAPIKey(apiKey string) {
	pUrl := url.URL{}
	pUrl.Scheme = "https"
	pUrl.User = url.UserPassword(apiKey, "")
	pUrl.Host = "www.pushbullet.com"
	pUrl.Path = "/api/pushes"
	pushUrl = pUrl.String()
}

func (pd *PushDev) PushNote(title, body string) (resp *http.Response, err error) {
	pushVals := url.Values{}
	pushVals.Set("device_id", pd.Id)
	pushVals.Set("type", "note")
	pushVals.Set("title", title)
	pushVals.Set("body", body)

	return http.PostForm(pushUrl, pushVals)
}

func main() {
	SetAPIKey("e4ac3e11929d522888c58ed67268b643")

	pushDev := &PushDev{"37413"}

	resp, err := pushDev.PushNote("testy sub", "long body orar\n \" ya")
	fmt.Println(resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Hello World!")
}
