// pbullet project main.go
package pbullet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

var pushUrl string
var getUrl string

type DeviceInfo struct {
	Manufacturer   string `json:"manufacturer"`
	Model          string `json:"model"`
	AndroidVersion string `json:"android_version"`
	SDKVersion     string `json:"sdk_version"`
	AppVersion     string `json:"app_version"`
	Nickname       string `json:"nickname"`
}

type Device struct {
	Id      int        `json:"id"`
	DevInfo DeviceInfo `json:"extras"`
	Owner   string     `json:"owner_name"`
}

type DeviceList struct {
	Devices       []Device `json:"devices"`
	SharedDevices []Device `json:"shared_devices"`
}

func SetAPIKey(apiKey string) {
	pUrl := url.URL{}
	pUrl.Scheme = "https"
	pUrl.User = url.UserPassword(apiKey, "")
	pUrl.Host = "www.pushbullet.com"
	pUrl.Path = "/api/pushes"
	pushUrl = pUrl.String()

	gUrl := url.URL{}
	gUrl.Scheme = "https"
	gUrl.User = url.UserPassword(apiKey, "")
	gUrl.Host = "www.pushbullet.com"
	gUrl.Path = "/api/devices"
	getUrl = gUrl.String()
}

func GetDevices() (DeviceList, error) {
	var devList DeviceList
	resp, err := http.Get(getUrl)
	if err != nil {
		return devList, err
	}
	fmt.Println(resp)
	respBytes, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(respBytes, &devList)
	return devList, err
}

func (pd *Device) PushNote(title, body string) (resp *http.Response, err error) {
	pushVals := url.Values{}
	pushVals.Set("device_id", strconv.Itoa(pd.Id))
	pushVals.Set("type", "note")
	pushVals.Set("title", title)
	pushVals.Set("body", body)

	return http.PostForm(pushUrl, pushVals)
}

func (pd *Device) PushAddress(name, address string) (resp *http.Response, err error) {
	pushVals := url.Values{}
	pushVals.Set("device_id", strconv.Itoa(pd.Id))
	pushVals.Set("type", "note")
	pushVals.Set("name", name)
	pushVals.Set("address", address)

	return http.PostForm(pushUrl, pushVals)
}

func (pd *Device) PushLink(title, urlAddress string) (resp *http.Response, err error) {
	pushVals := url.Values{}
	pushVals.Set("device_id", strconv.Itoa(pd.Id))
	pushVals.Set("type", "note")
	pushVals.Set("title", title)
	pushVals.Set("url", urlAddress)

	return http.PostForm(pushUrl, pushVals)
}
