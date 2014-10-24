package leancloud

import "net/url"

const classBaseURL = "classes"

func NewObject(className, jsonData, sessionToken string) *LeancloudResult {
	return HttpPost(sessionToken, jsonData, classBaseURL, className)
}

func GetObject(className, objectId, include, sessionToken string) *LeancloudResult {
	p := url.Values{}
	p.Add("include", include)
	return HttpGet(sessionToken, p, classBaseURL, className, objectId)
}

func UpdateObject(className, objectId, jsonData, sessionToken string) *LeancloudResult {
	return HttpPut(sessionToken, jsonData, classBaseURL, className, objectId)
}

func DeleteObject(className, objectId, sessionToken string) *LeancloudResult {
	return HttpDelete(sessionToken, classBaseURL, className, objectId)
}

func QueryObject(className, where, limit, skip, order, sessionToken string) *LeancloudResult {
	p := url.Values{}
	p.Add("where", where)
	p.Add("limit", limit)
	p.Add("skip", skip)
	p.Add("order", order)
	return HttpGet(sessionToken, p, classBaseURL, className)
}
