go-leancloud
============

a [leancloud](https://leancloud.cn/) [REST API](https://leancloud.cn/docs/rest_api.html) client library for go

install
------------
	```go get -u github.com/zx9597446/go-leancloud```

quick usage
-----------
1. config cloud:
```go
cloud := &leancloud.Cloud{}
cfg := leancloud.Config{}
cfg.AppId = ""
cfg.AppKey = ""
cfg.MasterKey = ""
cfg.UsingMaster = true
cloud.Cfg = cfg
```

2. Object CRUD
```go
className := "TestClass"
o := leancloud.NewObject()
o.Set("key", "value")
fetchOnSave := true
o.Create(cloud, className, fetchOnSave)

o.Set("newKey", "newValue")
o.Update(cloud, className)

o2 := leancloud.NewObject()
oid := o.ObjectId()
o2.Fetch(cloud, className, oid, "")

o2.Delete(cloud, className)
```

3. call leancloud function:
```go
param := leancloud.NewParams()
param["key"] = "value"
ret, err := cloud.CallFunction("functionName", param.Encode())
```

4. execute CQL:
```go
ret, err := cloud.CQL("select * from _User where username like 'abc%' limit 1")
```

5. call leancloud function(or object CRUD) with SessionToken:
```go
ret, err := cloud.WithSessionToken(token).CallFunction("fn", "param")
```

6. dump http request:
```go
	cloud.BeforeRequest = func(r *http.Request) *http.Request {
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
