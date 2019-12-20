package main

import (
	"back-api/internal/message"
	"back-api/internal/server"
	"back-api/internal/storage"
	"flag"
	"log"
	"os"
)

const (
	EnvMongoUri = "MONGODB_URI"
	EnvMongoDbName = "MONGODB_DBNAME"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", ":80", "http service address")
	flag.Parse()
}

func main() {
	mongoClient, err := storage.CreateMongoClient(os.Getenv(EnvMongoUri))

	if err != nil {
		log.Fatalln(err)
	}

	msgStorage, err := message.NewStorage(mongoClient, os.Getenv(EnvMongoDbName))

	if err != nil {
		log.Fatalln(err)
	}

	s := server.NewPrivateServer(msgStorage)
	s.Run(addr)
}