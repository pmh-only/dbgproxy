package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	client := &fasthttp.HostClient{
		Addr: TARGET_ADDRESS,
	}

	server := &fasthttp.Server{
		Handler: createProxyHandler(client),
	}

	err := server.ListenAndServe(LISTEN_ADDRESS)
	if err != nil {
		log.Fatalln(err)
	}
}
