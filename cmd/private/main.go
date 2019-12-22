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
var filename string

func init() {
	flag.StringVar(&addr, "addr", ":80", "http service address")
	flag.StringVar(&filename, "import", "", "file to import messages")
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

	if filename != "" {
		if err = msgStorage.Import(filename); err != nil {
			log.Fatalf("messages import failed: %s", err)
		}
	}

	s := server.NewPrivateServer(msgStorage)
	s.Run(addr)
}