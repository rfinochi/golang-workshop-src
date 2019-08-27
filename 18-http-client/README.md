# HTTP Client

The HTTP Client implementation from the [http](https://golang.org/pkg/net/http/) package allows users to emulate actions performed by a web browser.

For this tutorial, we shall make use of the awesome https://httpbin.org HTTP tester (source code [here](https://github.com/postmanlabs/httpbin)). This application can also be used locally by running the [kennethreitz/httpbin](https://hub.docker.com/r/kennethreitz/httpbin/) Docker image.

## Making a GET request

```go
resp, _ := http.Get("https://httpbin.org/get")
defer resp.Body.Close()

data, _ := ioutil.ReadAll(resp.Body)
fmt.Println(string(data))
```

## Making a POST request

```go
payload := "Hello world!"
resp, _ := http.Post("https://httpbin.org/post", "text/plain", strings.NewReader(payload))
defer resp.Body.Close()

data, _ := ioutil.ReadAll(resp.Body)
fmt.Println(string(data))
```

Note that if we don't get an error from `http.Get` or `http.Post`, we have to remember to close the response Body when we don't need it any more, since the library delegates this task to us. The `defer` statement helps us in this sense by closing the Body just before the enclosing function returns.

## Extracting information from a JSON response

Given this JSON:

```json
{
  "headers": {
    "Accept-Encoding": "gzip",
    "Host": "httpbin.org",
    "User-Agent": "Go-http-client/1.1"
  }
}
```

We can extract the `User-Agent` field like so:

```go
type response struct {
	Headers struct {
		UserAgent string `json:"User-Agent"`
	} `json:"headers"`
}

var resp response
_ := json.Unmarshal([]byte(jsonData), &resp)

fmt.Println(resp.Headers.UserAgent)
```

## HTTP request timeouts

When the HTTP request takes too long, we might want to have it terminate automatically after a fixed amount of time.

We can't use the convenience function `http.Get(url)` in this case.

Create a new request object with method GET:

```go
req, _ := http.NewRequest(http.MethodGet, url, nil)
```

Create a context with timeout:

```go
ctx, cancel := context.WithTimeout(context.Background(), timeout)
defer cancel()
```

Execute request by passing the request with the timeout context to the default HTTP client:

```go
resp, _ := http.DefaultClient.Do(req.WithContext(ctx))
defer resp.Body.Close()
```

## What are contexts?

The [context](https://golang.org/pkg/context/) package allows us to create a context object which can be used to pass some extra information to various functions.

`context.Background()` creates a new empty context, which does nothing.

Contexts can be chained together:

```go
ctx := context.Background()
ctxT1, cancel := context.WithTimeout(ctx, 3*time.Second)
ctxT2, cancelNew := context.WithTimeout(ctxT1, 2*time.Second)
```