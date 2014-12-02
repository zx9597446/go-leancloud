package leancloud

func (cloud *Cloud) CloudFunction(fn, jsonParam string) (*Result, error) {
	requestURL := cloud.makeURLPrefix("functions", fn)
	//param := url.Values{}
	//if uid != "" {
	//param.Add("uid", uid)
	//}
	//if jsonParam != "" {
	//param.Add("params", jsonParam)
	//}
	//old := cloud.HeaderContentType
	//cloud.HeaderContentType = "application/x-www-form-urlencoded"
	r, err := cloud.HttpPost(requestURL, jsonParam)
	//cloud.HeaderContentType = old
	return r, err
}

func (cloud *Cloud) CloudFunctionResultAsObject(fn, jsonParam string) (*Object, error) {
	r, err := cloud.CloudFunction(fn, jsonParam)
	if err != nil {
		return nil, err
	}
	return r.Decode()
}
