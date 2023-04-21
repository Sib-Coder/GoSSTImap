package algoritm

import (
	sructurs "awesomeProject1/struct"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func CheckPusReq(patern string, myurl string, path string) bool {
	resp, _ := http.Get(myurl + path) // переделать 2 строку на тестовый путь
	bytes, _ := ioutil.ReadAll(resp.Body)
	Body := string(bytes)
	matched, _ := regexp.MatchString(patern, Body) // пустой «приемник» ошибки, ведь мы уверены, что пример отработает нормально
	//fmt.Println(patern)//отлаживал алгоритм
	return matched
}
func ParsingJson() sructurs.PathAndPatern {
	content, err := ioutil.ReadFile("algoritm/file.json")
	if err != nil {
		log.Fatal(err)
	}
	user2 := sructurs.MasPatern{}
	err = json.Unmarshal(content, &user2)
	if err != nil {
		log.Fatal(err)
	}
	datapatern := sructurs.PathAndPatern{}
	datapatern.Path1 = user2.MasPatternPath[0].PathInEX //это тестовые пути
	datapatern.Path2 = user2.MasPatternPath[1].PathInEX
	datapatern.Path3 = user2.MasPatternPath[2].PathInEX
	datapatern.Path4 = user2.MasPatternPath[3].PathInEX
	datapatern.Path5 = user2.MasPatternPath[4].PathInEX

	datapatern.Patern1 = user2.MasPatternPath[0].PaternInEX
	datapatern.Patern2 = user2.MasPatternPath[1].PaternInEX
	datapatern.Patern3 = user2.MasPatternPath[2].PaternInEX
	datapatern.Patern4 = user2.MasPatternPath[3].PaternInEX
	datapatern.Patern5 = user2.MasPatternPath[4].PaternInEX

	return datapatern

}

func Algoritm(url string) {

	datapath := ParsingJson()

	if CheckPusReq(datapath.Patern1, url, datapath.Path1) == true {
		if CheckPusReq(datapath.Patern2, url, datapath.Path2) == true {
			fmt.Println("Smarty")
		} else if CheckPusReq(datapath.Patern3, url, datapath.Path3) == true {
			fmt.Println("Mako")
		}
	} else if CheckPusReq(datapath.Patern4, url, datapath.Path4) == true {
		if CheckPusReq(datapath.Patern5, url, datapath.Path5) == true {
			fmt.Println("Jinja2 and Twig")
		} else {
			fmt.Println("Unknown")
		}
	} else {
		fmt.Println("Не знаю")
	}
}
