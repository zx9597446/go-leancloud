package leancloud

import (
	"errors"
	"net/url"
)

const userBaseURL = "users"
const userClass = "_User"

type User struct {
	Object
}

func NewUser() *User {
	o := NewObject()
	return &User{*o}
}

func (u *User) Register(cloud *Client, username, password, email, phone string) (*result, error) {
	u.Set("username", username)
	u.Set("password", password)
	u.Set("email", email)
	u.Set("mobilePhoneNumber", phone)
	url := cloud.makeURLPrefix(userBaseURL)
	return cloud.httpPost(url, u.Encode())
}

func (u *User) Login(cloud *Client, username, password string) (*result, error) {
	p := url.Values{}
	p.Add("username", username)
	p.Add("password", password)
	uri := cloud.makeURLPrefix("login")
	r, err := cloud.httpGet(uri, p)
	if err != nil {
		return r, err
	}
	o, err := r.Decode()
	if err == nil {
		u.Data = o.Data
	}
	return r, err
}

func (u *User) fetchFrom(cloud *Client, objectId string) (*result, error) {
	url := cloud.makeURLPrefix(userBaseURL, objectId)
	r, err := cloud.httpGet(url, nil)
	if err != nil {
		return r, err
	}
	u1, err := r.Decode()
	if err == nil {
		u.Data = u1.Data
	}
	return r, err
}

func FetchUser(cloud *Client, userId string) (*User, error) {
	u := NewUser()
	_, err := u.fetchFrom(cloud, userId)
	return u, err
}

func GetUserSessionToken(cloud *Client, userId string) (string, error) {
	r, err := CQLf(cloud, "select * from _User where objectId = '%s'", userId)
	if err != nil {
		return "", err
	}
	users, err := r.GetResults()
	if err != nil {
		return "", err
	}
	if len(users) == 0 {
		return "", errors.New("empty results")
	}
	token, ok := users[0].Get("sessionToken").(string)
	if ok {
		return token, nil
	} else {
		return "", errors.New("convert to string failed")
	}
}
