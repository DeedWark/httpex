# httpex
[IN PROGRESS]

My HTTP Go package to make http requests easy to implement

## Headers
```go
headers := http.Header{
    "Content-Type": {"application/json"},
}
```

## Easy Requests
```go
g, err := httpex.Get(url string, headers http.Header)
h, err := httpex.Head(url string, headers http.Header)
p, err := httpex.Post(url string, body string, headers http.Header)
d, err := httpex.Delete(url string, body string, headers http.Header)

get, err := httpex.Get("https://github.com", headers)
if err != nil {
    log.Fatal(err)
}
```

## Custom Requests
```go
r, err := httpex.Request(method string, url string, data string, headers http.Header)

req, err := httpex.Request("GET", "https://github.com", "", headers)
if err != nil {
    log.Fatal(err)
}
// req = type of *http.Response
```

## Response
```go
r, err := httpex.Response(request *http.response)

res, err := httpex.Response(req)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(res))
// res = type of []byte
```

## Serve
```go
func serve() {
    opt := httpex.ServeOpt{
        Uri:       "/",
        Address:   ":8080",
        WebPage:   myFunc,
        LogMethod: os.Stderr,
        LogFormat: "CombinedLoggingHandler",
    }
    opt.Serve()
}
```

