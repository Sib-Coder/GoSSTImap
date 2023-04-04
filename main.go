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
	//patern := "777777"
	//chek1 := CheckPusReq(patern, myurl)
	//fmt.Println(chek1)
	Algoritm(myurl)

}

func CheckPusReq(patern string, myurl string, path string) bool {
	resp, _ := http.Get(myurl + path) // переделать 2 строку на тестовый путь
	bytes, _ := ioutil.ReadAll(resp.Body)
	Body := string(bytes)
	matched, _ := regexp.MatchString(patern, Body) // пустой «приемник» ошибки, ведь мы уверены, что пример отработает нормально
	fmt.Println(patern)
	return matched
}
func Algoritm(url string) {
	path1 := "${7*7}" //это не патерны - это тестовые пути
	path2 := "a{*comment*}b"
	path3 := "${\"z\".join(\"ab\")}" //пофиксить строку
	path4 := "{{7*7}}"
	path5 := "{{7*'7'}}"

	/// заменить верхнее на патерны
	patern1 := "${7*7}" //это не патерны - это тестовые пути
	patern2 := "a{*comment*}b"
	patern3 := "${\"z\".join(\"ab\")}" //пофиксить строку
	patern4 := "49"
	patern5 := "7777777"

	if CheckPusReq(patern1, url, path1) == true {
		if CheckPusReq(patern2, url, path2) == true {
			fmt.Println("Smarty")
		} else if CheckPusReq(patern3, url, path3) == true {
			fmt.Println("Mako")
		}
	} else if CheckPusReq(patern4, url, path4) == true {
		if CheckPusReq(patern5, url, path5) == true {
			fmt.Println("Jinja2 and Twig")
		} else {
			fmt.Println("Unknown")
		}
	} else {
		fmt.Println("Не знаю")
	}
}
