package main

import (
	"fmt"
	"github.com/anistark/gorouter"
	"net/http"
	"net/url"
	"os"
	"encoding/json"
)

type Payload struct {
	Stuff Data
}

type Data struct {
	Veggies Vegetables
}

type Vegetables map[string]int


func main() {
	r := gorouter.New(Root)
	//r.Use(fooMiddleware, barMiddleware, gorouter.Static()) // add global/router level middleware to run on every route.
	// Base Route
	r.Handle("GET", "/", Root)
	// User Routes
	r.Handle("GET", "/users", Users)
	//r.Handle("GET", "/users/:name", UserShow)
	//r.Handle("GET", "/users/:name/blog/new", UserBlogShow)
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
	//fmt.Fprint(w, "Users!\n")
	response, err := getJsonResponse();
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(response))
}

func UserShow(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprintf(w, "Hi %s", params["name"])
}

func UserBlogShow(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprintf(w, "This is %s Blog", params["name"])
}

func getJsonResponse()([]byte, error) {
	vegetables := make(map[string]int)
	vegetables["Carrats"] = 10
	vegetables["Beets"] = 0

	d := Data{vegetables}
	p := Payload{d}

	return json.MarshalIndent(p, "", "  ")
}
