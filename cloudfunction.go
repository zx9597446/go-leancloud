package leancloud

func (cloud *Client) callfunc(fn, jsonParam string) (*result, error) {
	requestURL := cloud.makeURLPrefix("functions", fn)
	r, err := cloud.httpPost(requestURL, jsonParam)
	return r, err
}

func CallFunction(cloud *Client, fn, jsonParam string) (*Object, error) {
	r, err := cloud.callfunc(fn, jsonParam)
	if err != nil {
		return nil, err
	}
	return r.Decode()
}
