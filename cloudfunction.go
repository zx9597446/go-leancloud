package leancloud

func (cloud *Cloud) CloudFunction(fn string, jsonParam string) (*Result, error) {
	url := cloud.makeURLPrefix("functions", fn)
	return cloud.HttpPost(url, jsonParam)
}
