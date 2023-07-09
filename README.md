# GoSSTImap
## Утилита для сканирования веб сайта на наличие SSTI. 
Данная утилита была написана в рамках курсовой работы за 3 курс и представляет из себя реализацию обхода алгоритма представленного ниже.
![Аглоритм]()

## Начало работы 
1) Если у вас нет в наличии тестового сервиса можно взять разработанный мою [тестовый сервис](https://github.com/Sib-Coder/TestSSTIWebServer).
```bash
git clone https://github.com/Sib-Coder/TestSSTIWebServer

docker-compose up

```
2) Склонировать репозиторий
```bash
git clone https://github.com/Sib-Coder/GoSSTImap
```
3) Проинициализированить проект
```golang
go mod init GoSSTImap
```
4) Запустить утилиту передав Url через флаг --url
```golang
go run main.go --url <Url>
```

## Пример работы
![пример работы]()
 

