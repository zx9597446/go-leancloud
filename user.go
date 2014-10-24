package leancloud

import "net/url"

const userBaseURL = "users"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"mobilePhoneNumber"`
}

func UserRegister(username, password, email, phone string) *LeancloudResult {
	u := User{username, password, email, phone}
	return HttpPost("", makeJSON(u), userBaseURL)
}

func UserLogin(username, password string) *LeancloudResult {
	p := url.Values{}
	p.Add("username", username)
	p.Add("password", password)
	return HttpGet("", p, "login")
}

func GetUser(objectId string) *LeancloudResult {
	return HttpGet("", url.Values{}, userBaseURL, objectId)
}
