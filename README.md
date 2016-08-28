go-leancloud
============

a [leancloud](https://leancloud.cn/) [REST API](https://leancloud.cn/docs/rest_api.html) client library for go

install
------------
```go get -u github.com/zx9597446/go-leancloud```

quick usage
-----------
1. config client:
```go
var client = &leancloud.Client{}
cfg := leancloud.Config{}
cfg.AppId = ""
cfg.AppKey = ""
cfg.MasterKey = ""
cfg.UsingMaster = true
client.Cfg = cfg
```

2. Object CRUD
```go
className := "TestClass"
o := leancloud.NewObject()
o.Set("key", "value")
fetchOnSave := true
o.Create(client, className, fetchOnSave)

o.Set("newKey", "newValue")
o.Update(client, className)

o2 := leancloud.NewObject()
oid := o.ObjectId()
o2.Fetch(client, className, oid, "")

o2.Delete(client, className)
```

3. call leancloud function:
```go
param := leancloud.NewParams()
param["key"] = "value"
ret, err := leancloud.CallFunction(client, "functionName", param.Encode())
```

4. execute CQL:
```go
ret, err := leancloud.CQL(client, "select * from _User where username like 'abc%' limit 1")
```

5. call leancloud function(or object CRUD) with SessionToken:
```go
ret, err := leancloud.CallFunction(client.WithSessionToken(""), fn", "param")
```

6. dump http request:
```go
	client.BeforeRequest = func(r *http.Request) *http.Request {
		//data, _ := httputil.DumpRequestOut(r, true)
		//log.Println(string(data))
		return r
	}
```

API doc
------------
see [doc](http://godoc.org/github.com/zx9597446/go-leancloud)

examples
-----------
see [test](http://github.com/zx9597446/go-leancloud/blob/master/lean_test.go)
