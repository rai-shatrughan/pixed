package gateway

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var pMap = make(map[string]string)
var scheme = "http"

func init() {
	pMap["127.0.0.1:8001"] = "/api/agentmanagement/v3/"
	pMap["127.0.0.1:8002"] = "/api/iottimeseries/v3/"
	pMap["127.0.0.1:8003"] = "/api/assetmanagement/v3/"
	pMap["127.0.0.1:8004"] = "/api/mindconnect/v3/"
}

// NewProxy takes target host and creates a reverse proxy
func newProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	return httputil.NewSingleHostReverseProxy(url), nil
}

// ProxyRequestHandler handles the http request using proxy
func proxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	sema := make(chan struct{}, 8000)
	return func(w http.ResponseWriter, r *http.Request) {
		sema <- struct{}{}
		defer func() { <-sema }()
		// proxy.Transport = &handler.Transport{}
		proxy.ServeHTTP(w, r)
	}
}

func Start(port int) {
	for k, v := range pMap {
		// initialize a reverse proxy and pass the actual backend server url here
		proxy, err := newProxy(scheme + "://" + k)
		if err != nil {
			panic(err)
		}

		// handle all requests to your server using the proxy
		go http.HandleFunc(v, proxyRequestHandler(proxy))
	}
	log.Printf("Starting Gateway at : %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
