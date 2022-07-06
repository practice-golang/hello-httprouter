package main // import "hello-httprouter"

import (
	"fmt"
	"hello-httprouter/route"
	"log"
	"net/http"

	// "github.com/julienschmidt/httprouter"
	"hello-httprouter/router"

	"github.com/rs/cors"
)

func main() {
	listen := "127.0.0.1:4416"

	r := router.New()

	r.GET("/", route.Index)
	r.GET("/hello/:name", route.Hello)

	handler := cors.Default().Handler(r)

	// router := httprouter.New()

	// router.GET("/", route.Index)
	// router.GET("/hello/:name", route.Hello)

	// router.POST("/login", route.Login)
	// router.POST("/user", route.User)

	// handler := cors.Default().Handler(router)
	// // c := cors.New(cors.Options{
	// // 	AllowedOrigins:   []string{"http://localhost:4416"},
	// // 	AllowedMethods:   []string{"GET"},
	// // 	AllowedHeaders:   []string{"*"},
	// // 	AllowCredentials: true,
	// // 	Debug:            false,
	// // })
	// // handler := c.Handler(router)

	fmt.Printf("listening on %s\n", listen)
	log.Fatal(http.ListenAndServe(listen, handler))
}
