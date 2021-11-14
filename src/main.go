package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello Usagisan/")
	//resp, err := http.Get("http://localhost:3500/v1.0/invoke/demo2/method/hello")
	//req, _ := http.NewRequest("GET", "http://demo2:8082/demo", nil)
	req, _ := http.NewRequest("POST", "http://localhost:3500/v1.0/publish/pubsub/demo", nil)
	//req.Header.Set("HTTP2_HEADER", "test")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(writer, resp.StatusCode)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(writer, string(body))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
