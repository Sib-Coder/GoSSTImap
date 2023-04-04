package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	//surl := flag.String("url", "http://127.0.0.1:5000/", "Give URL Address")
	fmt.Println("Start programm......")
	myurl := "http://localhost:9980/"
	//data := RequestToUrl(myurl)
	//fmt.Println(data)
	patern := "777777"
	chek1 := CheckPusReq(patern, myurl)
	fmt.Println(chek1)

}

// func RequestToUrl(myurl string) string { //можно объединить 2 эти функции в 1 и сделать единую функцию с входными параметрами url и патерн
//
//		resp, _ := http.Get(myurl + "{{7*'7'}}")
//		bytes, _ := ioutil.ReadAll(resp.Body)
//		Body := string(bytes)
//		return Body
//	}
func CheckPusReq(patern string, myurl string) bool {
	resp, _ := http.Get(myurl + "{{7*'7'}}")
	bytes, _ := ioutil.ReadAll(resp.Body)
	Body := string(bytes)
	matched, _ := regexp.MatchString(patern, Body) // пустой «приемник» ошибки, ведь мы уверены, что пример отработает нормально
	return matched
}
