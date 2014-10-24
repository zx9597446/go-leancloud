package leancloud

import (
	"net/http"
)
import "testing"

func init() {
	Config.AppId = ""
	Config.AppKey = ""
	Config.MasterKey = ""
	Config.UsingMaster = false
	Config.SiteURL = "https://avoscloud.us"
}

func checkResult(r *LeancloudResult, t *testing.T, expectCode int) {
	if r.StatusCode != expectCode {
		t.Fatalf("checkResult failed: %d, %s", r.StatusCode, r.JSON)
	}
	t.Log(r.JSON)
}

func TestNewObject(t *testing.T) {
	type testStruct struct {
		Name string
		Key  string
	}
	jsn := makeJSON(testStruct{"name", "key"})
	r := NewObject("testClass", jsn, "")
	checkResult(r, t, http.StatusCreated)
}

func TestGetObject(t *testing.T) {
	id := "544755c8e4b0327b4b90d3d7"
	r := GetObject("testClass", id, "game", "")
	checkResult(r, t, http.StatusOK)
}

func TestRegister(t *testing.T) {
	//r := UserRegister("zhang", "sdee", "ade@123.com", "13348598764")
	//checkResult(r, t, http.StatusCreated)
}

func TestLogin(t *testing.T) {
	r := UserLogin("zhang", "sdee")
	checkResult(r, t, http.StatusOK)
}

func TestGetUser(t *testing.T) {
	r := GetUser("5448ba00e4b0882ddff5dbf2")
	checkResult(r, t, http.StatusOK)
}
