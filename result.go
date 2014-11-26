package leancloud

type Result struct {
	StatusCode int
	Location   string
	Response   string
}

func (r *Result) Decode(className string) (*Object, error) {
	o := NewObject(className)
	err := o.Decode(r.Response)
	return o, err
}

func (r *Result) CheckStatusCode() bool {
	return r.StatusCode >= 200 && r.StatusCode < 300
}
