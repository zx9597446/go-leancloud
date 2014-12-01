package leancloud

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

var cloud = &Cloud{}

func init() {
	log.SetFlags(log.Lshortfile)

	cfg := Config{}
	cfg.AppId = ""
	cfg.AppKey = ""
	cfg.MasterKey = ""
	cfg.UsingMaster = true
	cloud.Cfg = cfg
	cloud.BeforeRequest = func(r *http.Request) *http.Request {
		//data, _ := httputil.DumpRequestOut(r, true)
		//log.Println(string(data))
		return r
	}

	rand.Seed(time.Now().UnixNano())
}

func randString() string {
	return fmt.Sprintf("abc%d", rand.Int())
}

func TestObject(t *testing.T) {
	className := "NewClass"
	o1 := NewObject(className)
	o1.Set("key", "value")
	r1, err := o1.Create(cloud, true)
	if err != nil {
		t.Fatal(r1, err)
	}
	if o1.ObjectId() == "" {
		t.Fatal("null objectId")
	}
	o1.Set("updatekey", "updatevalue")
	r2, err := o1.Update(cloud)
	if err != nil {
		t.Fatal(r2, err)
	}
	o2 := NewObject(className)
	r3, err := o2.Fetch(cloud, o1.ObjectId(), "")
	if err != nil {
		t.Fatal(r3, err)
	}
	r4, err := o2.Delete(cloud)
	if err != nil {
		t.Fatal(r4, err)
	}
}

func TestDate(t *testing.T) {
	className := "Class2"
	o1 := NewObject(className)
	d := FormatDate(time.Now())
	o1.Set("key", d)
	r1, err := o1.Create(cloud, true)
	if err != nil {
		t.Fatal(r1, err)
	}
	r2, err := o1.Delete(cloud)
	if err != nil {
		t.Fatal(r2, err)
	}
}

func TestUser(t *testing.T) {
	u1 := NewUser()
	email := fmt.Sprintf("%s@email.com", randString())
	phone := fmt.Sprintf("1386818%0d%0d", rand.Intn(99), rand.Intn(99))
	username := randString()
	password := "password"
	r1, err := u1.Register(cloud, username, password, email, phone)
	if err != nil {
		t.Fatal(r1, err)
	}
	r2, err := u1.Login(cloud, username, password)
	if err != nil {
		t.Fatal(r2, err)
	}
}

func TestCQL(t *testing.T) {
	r, err := cloud.CQL("select * from _User where username like 'abc%'")
	if err != nil {
		t.Fatal(err, r)
	}
}

func TestCloudFunction(t *testing.T) {
	//r, err := cloud.CloudFunction("syncDate", "")
	//if err != nil {
	//t.Fatal(err, r)
	//}
}

func Test1(t *testing.T) {
	//t.Fatal(time.Now().UTC().Format(time.RFC3339))
}
