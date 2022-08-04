package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// remoteURL, err := url.Parse("http://192.168.31.50:8001")
	localURL, err := url.Parse("http://192.168.31.50:8001")
	if err != nil {
		log.Fatal(err)
	}
	// proxy := httputil.NewSingleHostReverseProxy(remoteURL)
	proxy := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Println("HOST", req.Host)
		fmt.Println("BODY", req.Body)
		fmt.Println("HEADER", req.Header)
		fmt.Println("PROTO", req.Proto)
		fmt.Println("URI", req.RequestURI)
		fmt.Println("Addr", req.RemoteAddr)
		fmt.Println("scheme", req.URL.Scheme)

		//replaced a working local site running on my local server

		req.Host = "Caddy"
		req.RemoteAddr = "192.168.31.50:8001"
		req.RequestURI = ""
		req.URL = localURL

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			rw.WriteHeader((http.StatusInternalServerError))
			log.Fatalln("ERROR UNABLE TO CONNECT:", err)
			return
		}

		rw.WriteHeader(resp.StatusCode)
		io.Copy(rw, resp.Body)

	})

	http.ListenAndServe(":80", proxy)
}
