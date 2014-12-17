package leancloud

import (
	"fmt"
	"net/url"
)

func (cloud *Client) cql(q string) (*result, error) {
	p := url.Values{}
	p.Add("cql", q)
	u := cloud.makeURLPrefix("cloudQuery")
	return cloud.httpGet(u, p)
}

func CQL(cloud *Client, q string) (*Object, error) {
	r, err := cloud.cql(q)
	if err != nil {
		return nil, err
	}
	return r.Decode()
}

func CQLf(cloud *Client, format string, a ...interface{}) (*Object, error) {
	q := fmt.Sprintf(format, a...)
	return CQL(cloud, q)
}
