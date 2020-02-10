package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args[1:]) != 1 {
		if len(os.Args[1:]) > 1 {
			fmt.Println("Mais parametros do que o necessario!")
		} else {
			fmt.Println("Falta o parametro de mensagem!")
		}
		fmt.Println("Uso: gobot mensagem")
		os.Exit(1)
	} else if len(os.Args[1]) == 0 {
		fmt.Println("Mensagem nula!")
	        fmt.Println("Uso: gobot mensagem")
		os.Exit(1)
        }	
	var mensagem = os.Args[1]
	var hora = time.Now()
	var host, herr = os.Hostname()
	if herr != nil {
		log.Panic(herr)
		os.Exit(1)
	}
	if len(os.Getenv("TELEGRAM_TOKEN")) == 0 {
		fmt.Println("Falta a variavel de ambiente TELEGRAM_TOKEN !")
		os.Exit(1)
	} else if len(os.Getenv("TELEGRAM_TARGET")) == 0 {
		fmt.Println("Falta a variavel de ambiente TELEGRAM_TARGET !")
		os.Exit(1)
	}
	var token = os.Getenv("TELEGRAM_TOKEN")
	var destino, terr = strconv.Atoi(os.Getenv("TELEGRAM_TARGET"))
	if terr != nil {
		log.Panic(terr)
	}
	bot, berr := tgbotapi.NewBotAPI(token)
	if berr != nil {
		log.Panic(berr)
		os.Exit(1)
	}
	if os.Getenv("TBOT_DEBUG") == "1" {
		bot.Debug = true
		log.Println("Usando a conta: ", bot.Self.UserName)
	}
	msg := tgbotapi.NewMessage(int64(destino), strings.Join([]string{mensagem, host, hora.Format("2006-01-02 15:04")}, " - "))
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
