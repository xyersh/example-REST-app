package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/xyersh/examle-REST-app/internal/user"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("name")
	w.Write([]byte(fmt.Sprintf("Hello, %s", name)))
}

func main() {
	log.Println("Create router")
	router := httprouter.New()

	log.Println("Register handler")
	handler := user.NewHandler()
	handler.Register(router)

	Start(router)
}

func Start(router *httprouter.Router) {
	log.Println("Start application")

	listener, err := net.Listen("tcp", "127.0.0.1:8899")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatalln(server.Serve(listener))

}
