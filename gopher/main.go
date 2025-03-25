package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Gopher struct {
	Name string `json:"gopher"`
}

func GopherFunc() (gopher *Gopher) {
	gopher = &Gopher{Name: "gopher"}
	return gopher
}

func main() {
	const gopherName = "GOPHER"
	gogopher := GopherFunc()
	gogopher.Name = gopherName
	fmt.Println(*gogopher)

	// JSON形式で表示する場合
	jsonData, err := json.Marshal(gogopher)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}

