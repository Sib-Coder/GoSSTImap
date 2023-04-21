package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Paterns struct {
	NumberInAggoritm string
	PathInEX         string
	PaternInEX       string
}
type AutoGenerated struct {
	MasPatternPath []struct {
		NumberInAggoritm string `json:"NumberInAggoritm"`
		PathInEX         string `json:"PathInEX"`
		PaternInEX       string `json:"PaternInEX"`
	} `json:"MasPatternPath"`
}

func main() {

	content, err := ioutil.ReadFile("file.json")
	if err != nil {
		log.Fatal(err)
	}
	user2 := AutoGenerated{}
	err = json.Unmarshal(content, &user2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user2.MasPatternPath[0].PathInEX)

}