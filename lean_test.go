package leancloud

import "net/http"
import "testing"

func init() {
	Config.AppId = "appid"
	Config.AppKey = "appkey"
	Config.MasterKey = "masterkey"
	Config.UsingMaster = false
}

func checkResult(r *LeancloudResult, t *testing.T, expectCode int) {
	if r.Err != nil || r.HttpStatusCode != expectCode {
		t.Fatalf("checkResult failed: %v, %d, %s", r.Err, r.HttpStatusCode, r.RawData)
	}
	t.Log(r.RawData)
}

func TestNewObject(t *testing.T) {
	type testStruct struct {
		Name string
		Key  string
	}
	jsn := makeJSON(testStruct{"name", "key"})
	r := NewObject("testClass", jsn)
	checkResult(r, t, http.StatusCreated)
}

func TestGetObject(t *testing.T) {
	id := "544755c8e4b0327b4b90d3d7"
	r := GetObject("testClass", id, "game")
	checkResult(r, t, http.StatusOK)
}
