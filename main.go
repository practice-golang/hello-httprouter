package main // import "hello-httprouter"

import (
	"fmt"
	"hello-httprouter/route"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	listen := "127.0.0.1:4416"

	router := httprouter.New()

	router.GET("/", route.Index)
	router.GET("/hello/:name", route.Hello)

	router.POST("/login", route.Login)
	router.POST("/user", route.User)

	handler := cors.Default().Handler(router)
	// c := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:4416"},
	// 	AllowedMethods:   []string{"GET"},
	// 	AllowedHeaders:   []string{"*"},
	// 	AllowCredentials: true,
	// 	Debug:            false,
	// })
	// handler := c.Handler(router)

	fmt.Printf("listening on %s\n", listen)
	log.Fatal(http.ListenAndServe(listen, handler))
}
