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
	StatusCode int
	JSON       string
}

type leancloudParams struct {
	method       string
	url          string
	body         string
	sessionToken string
}

func httpRequest(param leancloudParams) (*LeancloudResult, error) {
	r, err := http.NewRequest(param.method, param.url, strings.NewReader(param.body))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Avoscloud-Application-Id", Config.AppId)
	//r.Header.Add("X-Avoscloud-Application-Key", Config.AppKey)
	r.Header.Set("X-AVOSCloud-Request-Sign", makeSign())
	if param.sessionToken != "" {
		r.Header.Set("X-AVOSCloud-Session-Token", param.sessionToken)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	sbody, _ := ioutil.ReadAll(res.Body)
	ret := &LeancloudResult{res.StatusCode, string(sbody)}
	return ret, nil
}

func makeURL(queryParams string, urlParts ...string) string {
	if Config.SiteURL == "" {
		Config.SiteURL = defaultSiteURL
	}
	if Config.Version == "" {
		Config.Version = defaultVersion
	}
	r := strings.Join(urlParts, "/")
	return fmt.Sprintf("%s/%s/%s?%s", Config.SiteURL, Config.Version, r, queryParams)
}

func mustRequest(param leancloudParams, query url.Values, urlParts ...string) *LeancloudResult {
	param.url = makeURL(query.Encode(), urlParts...)
	r, err := httpRequest(param)
	if err != nil {
		log.Fatalln(err)
	}
	return r
}

func HttpGet(sessionToken string, query url.Values, urlParts ...string) *LeancloudResult {
	return mustRequest(leancloudParams{"GET", "", "", sessionToken}, query, urlParts...)
}

func HttpPut(sessionToken, body string, urlParts ...string) *LeancloudResult {
	return mustRequest(leancloudParams{"PUT", "", body, sessionToken}, url.Values{}, urlParts...)
}

func HttpDelete(sessionToken string, urlParts ...string) *LeancloudResult {
	return mustRequest(leancloudParams{"DELETE", "", "", sessionToken}, url.Values{}, urlParts...)
}

func HttpPost(sessionToken, body string, urlParts ...string) *LeancloudResult {
	return mustRequest(leancloudParams{"POST", "", body, sessionToken}, url.Values{}, urlParts...)
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
