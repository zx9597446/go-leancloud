package leancloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
)

type Object struct {
	Data      map[string]interface{}
	ClassName string
}

var ErrNoClassOrNoObjectId = errors.New("no className or no objectId")

func NewObject(className string) *Object {
	o := &Object{}
	o.Data = make(map[string]interface{})
	o.ClassName = className
	return o
}

func (o *Object) Decode(data string) error {
	err := json.Unmarshal([]byte(data), &o.Data)
	if err == nil {
		o.ConvertToDate("createdAt")
		o.ConvertToDate("updatedAt")
	}
	return err
}

func (o *Object) ConvertToDate(key string) bool {
	v, ok := o.Data[key]
	if !ok {
		return false
	}
	o.Data[key] = NewDate(v.(string))
	return true
}

func (o *Object) Encode() string {
	data, err := json.Marshal(o.Data)
	if err != nil {
		log.Panicln(err)
	}
	return string(data)
}

func (o *Object) Get(key string) interface{} {
	return o.Data[key]
}

func (o *Object) Set(key string, value interface{}) {
	o.Data[key] = value
}

func (o *Object) Create(cloud *Cloud, fetchOnSave bool) (*Result, error) {
	r, err := cloud.CreateObject(o.ClassName, o.Encode())
	if !fetchOnSave || err != nil {
		return r, err
	}
	r, err = cloud.GetObjectDirectly(r.Location)
	if err != nil {
		return r, err
	}
	o1, err := r.Decode(o.ClassName)
	if err == nil {
		o.Data = o1.Data
	}
	return r, err
}

func (o *Object) Update(cloud *Cloud) (*Result, error) {
	if o.ObjectId() == "" || o.ClassName == "" {
		return nil, ErrNoClassOrNoObjectId
	}
	return cloud.UpdateObject(o.ClassName, o.ObjectId(), o.Encode())
}

func (o *Object) Delete(cloud *Cloud) (*Result, error) {
	if o.ObjectId() == "" || o.ClassName == "" {
		return nil, ErrNoClassOrNoObjectId
	}
	return cloud.DeleteObject(o.ClassName, o.ObjectId())
}

func (o *Object) Fetch(cloud *Cloud, objectId, include string) (*Result, error) {
	if o.ClassName == "" || objectId == "" {
		return nil, ErrNoClassOrNoObjectId
	}
	r, err := cloud.GetObject(o.ClassName, objectId, include)
	if err != nil {
		return r, err
	}
	o1, err := r.Decode(o.ClassName)
	if err == nil {
		o.Data = o1.Data
	}
	return r, err
}

func (o *Object) ObjectId() string {
	return o.Get("objectId").(string)
}

func (o *Object) AsPointer() Pointer {
	return NewPointer(o.ClassName, o.ObjectId())
}

func (o *Object) createdAt() Date {
	return o.Get("createdAt").(Date)
}

func (o *Object) updatedAt() Date {
	return o.Get("updatedAt").(Date)
}

type Date struct {
	Type string `json:"__type"`
	ISO  string `json:"iso"`
}

func NewDate(date string) Date {
	return Date{"Date", date}
}

//iso 格式是以 ISO 8601 标准和毫秒级精度储存:YYYY-MM-DDTHH:MM:SS.MMMMZ
func FormatDate(t time.Time) Date {
	t1 := t.UTC()
	s := fmt.Sprintf("%4d-%2d-%2dT%2d:%2d:%2d.%4dZ",
		t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute(), t1.Second(), t1.Nanosecond()%1e6/1e3)
	return Date{"Date", s}
}

type Pointer struct {
	ClassName string `json:"className"`
	Type      string `json:"__type"`
	ObjectId  string `json:"objectId"`
}

func NewPointer(class, oid string) Pointer {
	return Pointer{class, "Pointer", oid}
}

func NewUserPointer(oid string) Pointer {
	return NewPointer("_User", oid)
}
