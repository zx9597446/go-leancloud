package leancloud

func (cloud *Cloud) CallFunction(fn, jsonParam string) (*Result, error) {
	requestURL := cloud.makeURLPrefix("functions", fn)
	r, err := cloud.HttpPost(requestURL, jsonParam)
	return r, err
}

func (cloud *Cloud) CallFunctionResultAsObject(fn, jsonParam string) (*Object, error) {
	r, err := cloud.CallFunction(fn, jsonParam)
	if err != nil {
		return nil, err
	}
	return r.Decode()
}
