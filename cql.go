package leancloud

import "net/url"

func (cloud *Cloud) CQL(cql string) (*Result, error) {
	p := url.Values{}
	p.Add("cql", cql)
	u := cloud.makeURLPrefix("cloudQuery")
	return cloud.HttpGet(u, p)
}

func (cloud *Cloud) CQLResultAsObject(cql string) (*Object, error) {
	r, err := cloud.CQL(cql)
	if err != nil {
		return nil, err
	}
	return r.Decode()
}
