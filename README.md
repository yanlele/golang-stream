# 总结

### handler设计
handler -> validation(1.request, 2.user)-> business logic -> response                           
1、 data model                   
2、 error handler

```go
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}
```
这种是go原生处理json 的一种方法， 标记