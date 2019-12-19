package main

import (
	"back-api/internal/server"
	"flag"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", ":80", "http service address")
	flag.Parse()
}

func main() {
	s := server.NewPrivateServer()
	s.Run(addr)
}