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
o.Save(client, className, fetchOnSave)

o.Set("newKey", "newValue")
o.Update(client, className)
o.Delete(client, className)

obj, err := leancloud.FetchObject(client, className, objectId, include)

leancloud.DeleteObject(client, className, objectId)

```

3. call leancloud function:
```go
param := leancloud.NewParams()
param["key"] = "value"
ret, err := leancloud.CallFunction(client, "functionName", param.Encode())
```

4. execute CQL:
```go
//ret is a map["result"]={obj1, obj2...}
ret, err := leancloud.CQLf(client, "select * from _User where username like '%s' limit 1", "abc")

//to retrive all objects in ret, using this:
objects, err := ret.GetResults()

//to retrive one single object in ret, using this:
object, err := ret.GetResultByIdx(idx)

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
