//package main
//
//import (
//	"fmt"
//	"net/http"
//	"runtime"
//)
//
//func indexHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "hello world, I'm running on %s with an %s CPU ", runtime.GOOS, runtime.GOARCH)
//}
//
//func main() {
//	http.HandleFunc("/", indexHandler)
//	fmt.Println("Server started at http://localhost:8080")
//	http.ListenAndServe(":8080", nil)
//}

package main

import (
	"fmt"
	"github.com/anistark/gorouter"
	"net/http"
	"net/url"
	"os"
)

func main() {
	r := gorouter.New(Root)
	//r.Use(fooMiddleware, barMiddleware, gorouter.Static()) // add global/router level middleware to run on every route.
	r.Handle("GET", "/", Root)
	r.Handle("GET", "/users", Users)
	r.Handle("GET", "/users/:name", UserShow)
	r.Handle("GET", "/users/:name/blog/new", UserBlogShow)
	r.EnableLogging(os.Stdout) // Enable Router Logging to view on console.
	//http.ListenAndServe(":8080", r)
	r.Run(":8080") // Run Server
}

//func fooMiddleware(w http.ResponseWriter, r *http.Request, params url.Values) bool {
//	fmt.Println("Foo!")
//	return true
//}
//
//func barMiddleware(w http.ResponseWriter, r *http.Request, params url.Values) bool {
//	fmt.Println("Bar!")
//	return true
//}

func Root(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprint(w, "Root!\n")
}

func Users(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprint(w, "Users!\n")
}

func UserShow(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprintf(w, "Hi %s", params["name"])
}

func UserBlogShow(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprintf(w, "This is %s Blog", params["name"])
}
