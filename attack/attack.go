package attack

import (
	sructurs "awesomeProject1/struct"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func AttakForWeb(myurl string, nametycheck string) {
	exploit := ParsingExploit(nametycheck)
	resp, _ := http.Get(myurl + exploit) // переделать 2 строку на тестовый путь
	bytes, _ := ioutil.ReadAll(resp.Body)
	Body := string(bytes)
	fmt.Println("Код Ответа на Аттаку:")
	fmt.Println(Body)
}

func ParsingExploit(nametycheck string) string {

	content, err := ioutil.ReadFile("attack/file.json")
	if err != nil {
		log.Fatal(err)
	}
	user2 := sructurs.ExploitJson{}
	err = json.Unmarshal(content, &user2)
	if err != nil {
		log.Fatal(err)
	}

	exploit := "  "
	for i := 0; i < 3; i++ {
		if user2.MasTemplates[i].NameTeamplates == nametycheck {
			exploit = user2.MasTemplates[i].Exploits.Ex1
		}
	}
	return exploit
}
