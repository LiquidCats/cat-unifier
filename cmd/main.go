package main

import (
	"cat-unifier/internal/kernel/common/library/decorators"
	"cat-unifier/internal/kernel/common/library/factories"
	"cat-unifier/internal/kernel/common/library/repositiries"
	"cat-unifier/internal/kernel/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load("./configs/.env")
	if nil != err {
		log.Fatal(err)
	}

	configRepo := repositiries.NewConfigRepository(make(map[string]interface{}))

	readerFactory := factories.NewReaderFactory(configRepo)
	reader, err := readerFactory.Make()
	if nil != err {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	httpDecorator := decorators.NewHttpDecorator(reader)
	httpRoutes := http.NewHttpRoutes(r, configRepo, httpDecorator)

	server := http.NewServer(configRepo)
	server.Serve(httpRoutes)
}
