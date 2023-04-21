package main

import (
	"awesomeProject1/algoritm"
	"awesomeProject1/attack"
	"flag"
	"fmt"
)

func main() {
	//myurl := "http://localhost:9980/"
	myurl := flag.String("url", "0", "Give URL Address")
	Start()
	flag.Parse()
	url := *myurl
	if url != "0" {
		name := algoritm.Algoritm(url)
		if name == "  " {
			fmt.Println("Неизвестный шаблонизатор")
		} else {
			attack.AttakForWeb(url, name)
		}
	} else {
		fmt.Println("\nEror\nПерезапустите программу с флагом --url  ")
	}

}
func Start() {
	fmt.Println("Start programm......")
	fmt.Println("┏━━━┓━━━━┏━━━┓┏━━━┓┏━━━━┓┏━━┓━━━━━━━━━━━━━\n┃┏━┓┃━━━━┃┏━┓┃┃┏━┓┃┃┏┓┏┓┃┗┫┣┛━━━━━━━━━━━━━\n┃┃━┗┛┏━━┓┃┗━━┓┃┗━━┓┗┛┃┃┗┛━┃┃━┏┓┏┓┏━━┓━┏━━┓\n┃┃┏━┓┃┏┓┃┗━━┓┃┗━━┓┃━━┃┃━━━┃┃━┃┗┛┃┗━┓┃━┃┏┓┃\n┃┗┻━┃┃┗┛┃┃┗━┛┃┃┗━┛┃━┏┛┗┓━┏┫┣┓┃┃┃┃┃┗┛┗┓┃┗┛┃\n┗━━━┛┗━━┛┗━━━┛┗━━━┛━┗━━┛━┗━━┛┗┻┻┛┗━━━┛┃┏━┛\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃┃━━\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┗┛━━")
	fmt.Println("                         %              \n                       %%%%%            \n           %%%%       #%%%%%            \n           %%%%%%     %%%%%%       %    \n           #%%%%%%    ,%%%%      #%%%   \n    %%%      %%%%                %%%%%  \n   #%%%%%                       %%%%%%  \n   %%%%%%%%       %%%%%%%%#      #%%%   \n    %%%%%%%     %%%%%%%%%%%%%%%         \n      %%#     #%%%%%%%%%%%%%%%%%%       \n             %%%%%%%%%%%%%%%%%%%%       \n             %%%%%%%%%%%%%%%%%          \n              #%%%%%%%#                 ")
}
