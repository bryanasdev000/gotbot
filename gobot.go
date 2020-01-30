package main

import (
    "os"
    "fmt"
    "strconv"
    "log"	
    "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// TODO
	// host + msg + metadata
	//if len(os.Args[1:]) != 1 {		
	//	fmt.Println("./gobot 'TOKEN' 'DESTINO' 'MENSAGEM'")
	//	os.Exit(1)
	//}	
	//fmt.Println(os.Args)
	fmt.Println("Gobot")
	//var token = os.Args[1]
	//var mensagem = os.Args[3]
	//var destino,terr = strconv.Atoi(os.Args[2])
	//if terr != nil {
	//	log.Panic(terr)
	//}
	//var destino = os.Getenv("TELEGRAM_TARGET")
	var token = os.Getenv("TELEGRAM_TOKEN")
	var mensagem = os.Args[1]
	var destino,terr = strconv.Atoi(os.Getenv("TELEGRAM_TARGET"))
	if terr != nil {
		log.Panic(terr)
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
	//bot.Debug = true
	msg := tgbotapi.NewMessage(int64(destino), mensagem)
	if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
