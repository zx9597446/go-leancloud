package leancloud

func (cloud *Cloud) CloudFunction(fn string, jsonParam string) (*Result, error) {
	url := cloud.makeURLPrefix("functions", fn)
	return cloud.HttpPost(url, jsonParam)
}

func (cloud *Cloud) CloudFunctionResultAsObject(fn, jsonParam, className string) (*Object, error) {
	r, err := cloud.CloudFunction(fn, jsonParam)
	if err != nil {
		return nil, err
	}
	return r.Decode(className)
}
