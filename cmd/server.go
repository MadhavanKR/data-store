package main

import (
	"github.com/milkmocha-datastore/pkg/httpServer"
	"github.com/milkmocha-datastore/pkg/redisAdapter"
)

func main() {
	redisAdapter.ConnectAndStore()
	server := httpServer.NewHttpServer()
	server.ListenAndServe()
}
