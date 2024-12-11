package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)


func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
   name := params.ByName("name")
   w.Write([]byte(name))
}

func main() {
    router := httprouter.New()

    router.GET("/:name", IndexHandler)

    listen, err := net.Listen("tcp", ":1234")   

     if err!= nil {
        log.Fatal(err)
    }
    
    server := &http.Server{
        Handler: router,
        WriteTimeout: 15 * time.Second,
        ReadTimeout: 15 * time.Second, 
    }
    
    log.Fatal(server.Serve(listen))
}