package leancloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
)

type Object struct {
	Data map[string]interface{}
}

var ErrNoObjectIdOrClassName = errors.New("no objectId or no class name")

func NewObject() *Object {
	o := &Object{}
	o.Data = make(map[string]interface{}, 0)
	return o
}

func FetchObject(cloud *Client, className, objectId, include string) (*Object, error) {
	if className == "" || objectId == "" {
		return nil, ErrNoObjectIdOrClassName
	}
	r, err := cloud.getObject(className, objectId, include)
	if err != nil {
		return nil, err
	}
	return r.Decode()
}

func DeleteObject(cloud *Client, className, objectId string) error {
	if objectId == "" || className == "" {
		return ErrNoObjectIdOrClassName
	}
	_, err := cloud.deleteObject(className, objectId)
	return err
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

func (o *Object) Save(cloud *Client, className string, fetchOnSave bool) error {
	r, err := cloud.createObject(className, o.Encode())
	if !fetchOnSave || err != nil {
		return err
	}
	r, err = cloud.getObjectDirectly(r.Location)
	if err != nil {
		return err
	}
	o1, err := r.Decode()
	if err == nil {
		o.Data = o1.Data
	}
	return err
}

func (o *Object) Update(cloud *Client, className string) error {
	if o.ObjectId() == "" || className == "" {
		return ErrNoObjectIdOrClassName
	}
	_, err := cloud.updateObject(className, o.ObjectId(), o.Encode())
	return err
}

func (o *Object) Delete(cloud *Client, className string) error {
	if o.ObjectId() == "" || className == "" {
		return ErrNoObjectIdOrClassName
	}
	_, err := cloud.deleteObject(className, o.ObjectId())
	return err
}

func (o *Object) ObjectId() string {
	return o.Get("objectId").(string)
}

func (o *Object) AsPointer(className string) Pointer {
	return NewPointer(className, o.ObjectId())
}

func (o *Object) GetResults() ([]*Object, error) {
	results, ok := o.Data["results"]
	if !ok {
		return nil, errors.New("no such key `results`")
	}
	interfaces, ok := results.([]interface{})
	if !ok {
		return nil, errors.New("convert to []interface{} failed")
	}
	objects := make([]*Object, 0)
	for _, v := range interfaces {
		o := NewObject()
		data, ok := v.(map[string]interface{})
		if !ok {
			return nil, errors.New("convert to map[string]interface{} failed")
		}
		o.Data = data
		objects = append(objects, o)
	}
	return objects, nil
}

func (o *Object) CreatedAt() Date {
	return o.Get("createdAt").(Date)
}

func (o *Object) UpdatedAt() Date {
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
