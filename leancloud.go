package leancloud

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const defaultSiteURL = "https://leancloud.cn"
const defaultVersion = "1.1"

type LeancloudConfig struct {
	AppId       string
	AppKey      string
	MasterKey   string
	UsingMaster bool
	Version     string
	SiteURL     string
}

var Config LeancloudConfig

type LeancloudResult struct {
	HttpStatusCode int
	RawData        string
	Err            error
}

func httpRequest(url, method, body string) (string, int, error) {
	//log.Println(url, method, body)
	r, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return "", 0, err
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Avoscloud-Application-Id", Config.AppId)
	//r.Header.Add("X-Avoscloud-Application-Key", Config.AppKey)
	r.Header.Set("X-AVOSCloud-Request-Sign", makeSign())
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//debug, _ := httputil.DumpRequest(r, true)
	//log.Println(string(debug))
	client := &http.Client{Transport: tr}
	res, err := client.Do(r)
	if err != nil {
		return "", 0, err
	}
	defer res.Body.Close()
	res_body, _ := ioutil.ReadAll(res.Body)
	return string(res_body), res.StatusCode, nil
}

func tryHttpRequest(method, body, urlParams string, urlParts ...string) *LeancloudResult {
	url := makeURL(urlParams, urlParts...)
	d, c, err := httpRequest(url, method, body)
	if err != nil {
		log.Println(err)
	}
	return &LeancloudResult{c, d, err}
}

func HttpGet(v url.Values, urlParts ...string) *LeancloudResult {
	return tryHttpRequest("GET", "", v.Encode(), urlParts...)
}

func HttpPut(body string, urlParts ...string) *LeancloudResult {
	return tryHttpRequest("PUT", body, "", urlParts...)
}

func HttpDelete(urlParts ...string) *LeancloudResult {
	return tryHttpRequest("DELETE", "", "", urlParts...)
}

func HttpPost(body string, urlParts ...string) *LeancloudResult {
	return tryHttpRequest("POST", body, "", urlParts...)
}

func makeURL(urlParams string, urlParts ...string) string {
	if Config.SiteURL == "" {
		Config.SiteURL = defaultSiteURL
	}
	if Config.Version == "" {
		Config.Version = defaultVersion
	}
	r := strings.Join(urlParts, "/")
	return fmt.Sprintf("%s/%s/%s?%s", Config.SiteURL, Config.Version, r, urlParams)
}

func makeSign() string {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	sign := ""
	if Config.UsingMaster {
		sign = fmt.Sprintf("%x", md5.Sum([]byte(timestamp+Config.MasterKey)))
		return fmt.Sprintf("%s,%s,%s", sign, timestamp, "master")
	} else {
		sign = fmt.Sprintf("%x", md5.Sum([]byte(timestamp+Config.AppKey)))
		return fmt.Sprintf("%s,%s", sign, timestamp)
	}
}

func makeJSON(v interface{}) string {
	d, _ := json.Marshal(v)
	return string(d)
}
