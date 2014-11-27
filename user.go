package leancloud

import "net/url"

const userBaseURL = "users"
const userClass = "_User"

type User struct {
	Object
}

func NewUser() *User {
	o := NewObject(userClass)
	return &User{*o}
}

func (u *User) Register(cloud *Cloud, username, password, email, phone string) (*Result, error) {
	u.Set("username", username)
	u.Set("password", password)
	u.Set("email", email)
	u.Set("mobilePhoneNumber", phone)
	url := cloud.makeURLPrefix(userBaseURL)
	return cloud.HttpPost(url, u.Encode())
}

func (u *User) Login(cloud *Cloud, username, password string) (*Result, error) {
	p := url.Values{}
	p.Add("username", username)
	p.Add("password", password)
	uri := cloud.makeURLPrefix("login")
	r, err := cloud.HttpGet(uri, p)
	if err != nil {
		return r, err
	}
	o, err := r.Decode(u.ClassName)
	if err == nil {
		u.Data = o.Data
	}
	return r, err
}

func (u *User) Get(cloud *Cloud, objectId string) (*Result, error) {
	url := cloud.makeURLPrefix(userBaseURL, objectId)
	r, err := cloud.HttpGet(url, nil)
	if err != nil {
		return r, err
	}
	u1, err := r.Decode(userClass)
	if err == nil {
		u.Data = u1.Data
	}
	return r, err
}
