package leancloud

import "net/url"

const classBaseURL = "classes"

func (cloud *Client) makeClassURL(parts ...string) string {
	return cloud.makeURLPrefix(classBaseURL, parts...)
}

func (cloud *Client) createObject(className, jsonData string) (*result, error) {
	return cloud.httpPost(cloud.makeClassURL(className), jsonData)
}

func (cloud *Client) getObject(className, objectId, include string) (*result, error) {
	p := url.Values{}
	p.Add("include", include)
	url := cloud.makeClassURL(className, objectId)
	return cloud.httpGet(url, p)
}

func (cloud *Client) getObjectDirectly(location string) (*result, error) {
	return cloud.httpGet(location, nil)
}

func (cloud *Client) updateObject(className, objectId, jsonData string) (*result, error) {
	url := cloud.makeClassURL(className, objectId)
	return cloud.httpPut(url, jsonData)
}

func (cloud *Client) deleteObject(className, objectId string) (*result, error) {
	url := cloud.makeClassURL(className, objectId)
	return cloud.httpDelete(url)
}

func (cloud *Client) queryObject(className, whereJson, limit, skip, order, keys string) (*result, error) {
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
	return cloud.httpGet(url, p)
}
