# grpc-websocket-proxy

[![GoDoc](https://godoc.org/github.com/tmc/grpc-websocket-proxy/wsproxy?status.svg)](http://godoc.org/github.com/tmc/grpc-websocket-proxy/wsproxy)

Wrap your grpc-gateway mux with this helper to expose streaming endpoints over websockets.

Usage:
```diff
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := echoserver.RegisterEchoServiceHandlerFromEndpoint(ctx, mux, *grpcAddr, opts); err != nil {
		return err
	}
-	http.ListenAndServe(*httpAddr, mux)
+	http.ListenAndServe(*httpAddr, wsproxy.WebsocketProxy(mux))
```


# Package wsproxy

`import "github.com/tmc/grpc-websocket-proxy/wsproxy"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package wsproxy implements a websocket proxy for grpc-gateway backed services




## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func WebsocketProxy(h http.Handler) http.HandlerFunc](#WebsocketProxy)


#### <a name="pkg-files">Package files</a>
[doc.go](/src/github.com/tmc/grpc-websocket-proxy/wsproxy/doc.go) [websocket_proxy.go](/src/github.com/tmc/grpc-websocket-proxy/wsproxy/websocket_proxy.go) 



## <a name="pkg-variables">Variables</a>
``` go
var (
    MethodOverrideParam = "method"
    TokenCookieName     = "token"
)
```


## <a name="WebsocketProxy">func</a> [WebsocketProxy](/src/target/websocket_proxy.go?s=758:810#L21)
``` go
func WebsocketProxy(h http.Handler) http.HandlerFunc
```
WebsocketProxy attempts to expose the underlying handler as a bidi websocket stream with newline-delimited
JSON as the content encoding.

The HTTP Authorization header is either populated from the Sec-Websocket-Protocol field or by a cookie.
The cookie name is specified by the TokenCookieName value.

example:


	Sec-Websocket-Protocol: Bearer, foobar

is converted to:


	Authorization: Bearer foobar

Method can be overwritten with the MethodOverrideParam get parameter in the requested URL


