package main

import (
	"back-api/internal/authentication"
	"back-api/internal/message"
	"back-api/internal/password"
	"back-api/internal/server"
	"back-api/internal/storage"
	"back-api/internal/user"
	"flag"
	"log"
	"os"
)

const (
	EnvMongoUri    = "MONGODB_URI"
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

	authStorage, err := user.NewStorage(mongoClient, os.Getenv(EnvMongoDbName))

	if err != nil {
		log.Fatalln(err)
	}

	auth := authentication.NewAuth(authStorage, password.NewPasswordService())

	s := server.NewPrivateServer(msgStorage, auth)
	s.Run(addr)
}
