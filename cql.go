package leancloud

import "net/url"

func (cloud *Cloud) CQL(cql string) (*Result, error) {
	p := url.Values{}
	p.Add("cql", cql)
	u := cloud.makeURLPrefix("cloudQuery")
	return cloud.Get(u, p)
}
