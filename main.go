package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: time.Second * 2,
	}

	url := "https://jsonplaceholder.typicode.com/posts"
	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Status code error: %d %s", res.StatusCode, res.Status))
	}

	fmt.Println(res.Header.Get("Content-Type"))

	type Post struct {
		userId int `json:"userId"`
		id int `json:"id"`
		title string `json:"title"`
		body string `json:"body"`
	}

	var posts []Post
	resBody, _ := ioutil.ReadAll(res.Body)

	fmt.Printf("%T\n", resBody) // []uint8
	fmt.Println(string(resBody))

	err = json.Unmarshal(resBody, &posts)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T\n", posts) // []main.Post
	fmt.Printf("%+v\n", posts)
}
