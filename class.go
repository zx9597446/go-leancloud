package leancloud

import "net/url"

const classBaseURL = "classes"

func (cloud *Cloud) makeClassURL(parts ...string) string {
	return cloud.makeURLPrefix(classBaseURL, parts...)
}

func (cloud *Cloud) CreateObject(className, jsonData string) (*Result, error) {
	return cloud.HttpPost(cloud.makeClassURL(className), jsonData)
}

func (cloud *Cloud) GetObject(className, objectId, include string) (*Result, error) {
	p := url.Values{}
	p.Add("include", include)
	url := cloud.makeClassURL(className, objectId)
	return cloud.HttpGet(url, p)
}

func (cloud *Cloud) GetObjectDirectly(location string) (*Result, error) {
	return cloud.HttpGet(location, nil)
}

func (cloud *Cloud) UpdateObject(className, objectId, jsonData string) (*Result, error) {
	url := cloud.makeClassURL(className, objectId)
	return cloud.HttpPut(url, jsonData)
}

func (cloud *Cloud) DeleteObject(className, objectId string) (*Result, error) {
	url := cloud.makeClassURL(className, objectId)
	return cloud.HttpDelete(url)
}

func (cloud *Cloud) QueryObject(className, whereJson, limit, skip, order, keys string) (*Result, error) {
	p := url.Values{}
	if whereJson != "" {
		p.Add("where", whereJson)
	}
	if limit != "" {
		p.Add("limit", limit)
	}
	if skip != "" {
		p.Add("skip", skip)
	}
	if order != "" {
		p.Add("order", order)
	}
	if keys != "" {
		p.Add("keys", keys)
	}
	url := cloud.makeClassURL(className)
	return cloud.HttpGet(url, p)
}
