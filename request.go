package leancloud

import (
	"crypto/md5"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const defaultSiteURL = "https://leancloud.cn/1.1"

type Config struct {
	AppId       string
	AppKey      string
	MasterKey   string
	UsingMaster bool
	SiteURL     string
}

type Client struct {
	Cfg              Config
	HeaderProduction string
	BeforeRequest    func(*http.Request) *http.Request
	sessionToken     string
}

func (c *Client) WithSessionToken(token string) *Client {
	ac := c.clone()
	ac.sessionToken = token
	return ac
}

func (cloud *Client) clone() *Client {
	return &Client{cloud.Cfg, cloud.HeaderProduction, cloud.BeforeRequest, cloud.sessionToken}
}

func (cloud *Client) makeSign() string {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	sign := ""
	if cloud.Cfg.UsingMaster {
		sign = fmt.Sprintf("%x", md5.Sum([]byte(timestamp+cloud.Cfg.MasterKey)))
		return fmt.Sprintf("%s,%s,%s", sign, timestamp, "master")
	} else {
		sign = fmt.Sprintf("%x", md5.Sum([]byte(timestamp+cloud.Cfg.AppKey)))
		return fmt.Sprintf("%s,%s", sign, timestamp)
	}
}

func (cloud *Client) makeURL(parts ...string) (url string) {
	var path string
	if len(parts) == 0 {
		log.Panicln("can not make url", parts)
	} else if len(parts) == 1 {
		path = parts[0]
	} else {
		path = strings.Join(parts, "/")
	}
	if cloud.Cfg.SiteURL == "" {
		url = fmt.Sprintf("%s/%s", defaultSiteURL, path)
	} else {
		url = fmt.Sprintf("%s/%s", cloud.Cfg.SiteURL, path)
	}
	return
}

func (cloud *Client) makeURLPrefix(prefix string, parts ...string) string {
	tmp := make([]string, 0)
	tmp = append(tmp, prefix)
	tmp = append(tmp, parts...)
	return cloud.makeURL(tmp...)
}

func (cloud *Client) httpRequest(url, method, body string) (*result, error) {
	r, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	r.Header.Set("X-Avoscloud-Application-Id", cloud.Cfg.AppId)
	r.Header.Set("X-AVOSCloud-Request-Sign", cloud.makeSign())
	if cloud.sessionToken != "" {
		r.Header.Set("X-AVOSCloud-Session-Token", cloud.sessionToken)
	}
	r.Header.Set("Content-Type", "application/json")
	if cloud.HeaderProduction != "" {
		r.Header.Set("X-AVOSCloud-Application-Production", cloud.HeaderProduction)
	}
	if cloud.BeforeRequest != nil {
		r = cloud.BeforeRequest(r)
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
	var location string
	if u, err := res.Location(); err == nil {
		location = u.String()
	}
	ret := &result{res.StatusCode, location, string(sbody)}
	if !ret.CheckStatusCode() {
		return ret, errors.New(ret.Response)
	}
	return ret, nil
}

func (cloud *Client) httpGet(url string, param url.Values) (*result, error) {
	withQuery := fmt.Sprintf("%s?%s", url, param.Encode())
	return cloud.httpRequest(withQuery, "GET", "")
}

func (cloud *Client) httpPut(url, body string) (*result, error) {
	return cloud.httpRequest(url, "PUT", body)
}

func (cloud *Client) httpDelete(url string) (*result, error) {
	return cloud.httpRequest(url, "DELETE", "")
}

func (cloud *Client) httpPost(url, body string) (*result, error) {
	return cloud.httpRequest(url, "POST", body)
}
