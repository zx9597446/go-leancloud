package leancloud

import (
	"encoding/json"
	"log"
)

type Result struct {
	StatusCode int
	Location   string
	Response   string
}

func (r *Result) Decode() (*Object, error) {
	o := NewObject()
	err := o.Decode(r.Response)
	return o, err
}

func (r *Result) CheckStatusCode() bool {
	return r.StatusCode >= 200 && r.StatusCode < 300
}

type Params map[string]interface{}

func NewParams() Params {
	return make(map[string]interface{}, 0)
}

func (p Params) Encode() string {
	data, err := json.Marshal(p)
	if err != nil {
		log.Panicln(err)
	}
	return string(data)
}
