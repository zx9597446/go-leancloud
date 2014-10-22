package leancloud

import "net/url"

const classBaseURL = "classes"

func NewObject(className string, jsn string) *LeancloudResult {
	return HttpPost(jsn, classBaseURL, className)
}

func GetObject(className, objectId, include string) *LeancloudResult {
	p := url.Values{}
	p.Add("include", include)
	return HttpGet(p, classBaseURL, className, objectId)
}

func UpdateObject(className, objectId string, jsn string) *LeancloudResult {
	return HttpPut(jsn, classBaseURL, className, objectId)
}

func DeleteObject(className, objectId string) *LeancloudResult {
	return HttpDelete(classBaseURL, className, objectId)
}

func QueryObject(className, where, limit, skip, order string) *LeancloudResult {
	p := url.Values{}
	p.Add("where", where)
	p.Add("limit", limit)
	p.Add("skip", skip)
	p.Add("order", order)
	return HttpGet(p, classBaseURL, className)
}
