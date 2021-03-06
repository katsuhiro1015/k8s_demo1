package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"k8s_demo1/output"
	"log"
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

func list_handler(writer http.ResponseWriter, request *http.Request) {
	c, err := output.NewClient("http://localhost:3500/v1.0/invoke/demo2/method/")
	if err != nil {
		panic(err)
	}

	params := output.FindPetsParams{Tags: &[]string{"yosistamp"}}
	// http.Response として返却
	res, err := c.FindPets(context.Background(), &params)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(writer, string(b))

}

func add_handler(writer http.ResponseWriter, request *http.Request) {
	c, err := output.NewClient("http://localhost:3500/v1.0/invoke/demo2/method/")
	if err != nil {
		panic(err)
	}

	tag := "yosistamp"
	params := output.AddPetJSONRequestBody{
		Name: "usagisan",
		Tag:  &tag,
	}
	// http.Response として返却
	res, err := c.AddPet(context.Background(), params)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(writer, string(b))

}
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/oapi/add", add_handler)
	http.HandleFunc("/oapi", list_handler)
	http.ListenAndServe(":8080", nil)
}
