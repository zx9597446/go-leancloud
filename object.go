package leancloud

import (
	"encoding/json"
	"errors"
	"log"
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
	return json.Unmarshal([]byte(data), &o.Data)
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
	if !fetchOnSave {
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
	delete(o.Data, "createdAt")
	delete(o.Data, "updatedAt")
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

func (o *Object) CreatedAt() string {
	return o.Get("createdAt").(string)
}

func (o *Object) UpdatedAt() string {
	return o.Get("updatedAt").(string)
}

func (o *Object) ObjectId() string {
	return o.Get("objectId").(string)
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
