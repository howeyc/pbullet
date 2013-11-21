// Go library for the Push Bullet REST API
// More info: https://www.pushbullet.com/api
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

// Extra device info returned by GetDevices API call.
type DeviceInfo struct {
	Manufacturer   string `json:"manufacturer"`
	Model          string `json:"model"`
	AndroidVersion string `json:"android_version"`
	SDKVersion     string `json:"sdk_version"`
	AppVersion     string `json:"app_version"`
	Nickname       string `json:"nickname"`
}

// Device is the structure needed to push Notes/Addresses/Links/etc.
// Only need to populate Id field. Other fields are informational only.
type Device struct {
	Id      int        `json:"id"`
	DevInfo DeviceInfo `json:"extras"`
	Owner   string     `json:"owner_name"`
}

// GetDevices returns two lists, owned devices and devices that are shared.
type DeviceList struct {
	Devices       []Device `json:"devices"`
	SharedDevices []Device `json:"shared_devices"`
}

// Set the API key used for all Get and Push API calls.
func SetAPIKey(apiKey string) {
	pUrl := url.URL{}
	pUrl.Scheme = "https"
	pUrl.User = url.UserPassword(apiKey, "")
	pUrl.Host = "api.pushbullet.com"
	pUrl.Path = "/api/pushes"
	pushUrl = pUrl.String()

	gUrl := url.URL{}
	gUrl.Scheme = "https"
	gUrl.User = url.UserPassword(apiKey, "")
	gUrl.Host = "api.pushbullet.com"
	gUrl.Path = "/api/devices"
	getUrl = gUrl.String()
}

// Get devices configured on PushBullet
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

// Push a note to a device.
func (pd *Device) PushNote(title, body string) (resp *http.Response, err error) {
	pushVals := url.Values{}
	pushVals.Set("device_id", strconv.Itoa(pd.Id))
	pushVals.Set("type", "note")
	pushVals.Set("title", title)
	pushVals.Set("body", body)

	return http.PostForm(pushUrl, pushVals)
}

// Push an address to a device.
func (pd *Device) PushAddress(name, address string) (resp *http.Response, err error) {
	pushVals := url.Values{}
	pushVals.Set("device_id", strconv.Itoa(pd.Id))
	pushVals.Set("type", "note")
	pushVals.Set("name", name)
	pushVals.Set("address", address)

	return http.PostForm(pushUrl, pushVals)
}

// Push a link to a device.
func (pd *Device) PushLink(title, urlAddress string) (resp *http.Response, err error) {
	pushVals := url.Values{}
	pushVals.Set("device_id", strconv.Itoa(pd.Id))
	pushVals.Set("type", "note")
	pushVals.Set("title", title)
	pushVals.Set("url", urlAddress)

	return http.PostForm(pushUrl, pushVals)
}
