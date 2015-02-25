package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type authRewriteTransport struct {
	base                                               http.RoundTripper
	remote                                             *url.URL
	accessorUser, accessorPass, remoteUser, remotePass string
}

func (t *authRewriteTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	log.Println("Request from " + req.RemoteAddr + " to " + req.RequestURI)
	givenUser, givenPass, isGiven := req.BasicAuth()
	if isGiven && givenUser == t.accessorUser && givenPass == t.accessorPass {
		req.SetBasicAuth(t.remoteUser, t.remotePass)
	}
	req.Host = t.remote.Host
	resp, err = t.base.RoundTrip(req)
	return
}

func main() {
	port := os.Getenv("PORT")
	remoteUrl := os.Getenv("GITPROXY_REMOTE_URL")

	url, err := url.Parse(remoteUrl)
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.Transport = &authRewriteTransport{
		base:         http.DefaultTransport,
		remote:       url,
		accessorUser: os.Getenv("GITPROXY_ACCESSOR_USER"),
		accessorPass: os.Getenv("GITPROXY_ACCESSOR_PASS"),
		remoteUser:   os.Getenv("GITPROXY_REMOTE_USER"),
		remotePass:   os.Getenv("GITPROXY_REMOTE_PASS"),
	}

	log.Fatal(http.ListenAndServe(":"+port, proxy))
}
