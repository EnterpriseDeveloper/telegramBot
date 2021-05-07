package main

import (
	"log"
	"path/filepath"

	"github.com/zelenin/go-tdlib/client"

	"fmt"
	"strings"

	binanceAPI "github.com/bot/binanceAPI"
	config "github.com/bot/config"
	db "github.com/bot/db"
)

func main() {
	dataBase := db.CreateConnection()
	defer dataBase.Close()
	// client authorizer
	authorizer := client.ClientAuthorizer()
	go client.CliInteractor(authorizer)

	// or bot authorizer
	// botToken := "000000000:gsVCGG5YbikxYHC7bP5vRvmBqJ7Xz6vG6td"
	// authorizer := client.BotAuthorizer(botToken)

	const (
		apiId   = 4489698
		apiHash = "198dea4165bc5155e490bcbc267ec4df"
	)

	authorizer.TdlibParameters <- &client.TdlibParameters{
		UseTestDc:              false,
		DatabaseDirectory:      filepath.Join(".tdlib", "database"),
		FilesDirectory:         filepath.Join(".tdlib", "files"),
		UseFileDatabase:        true,
		UseChatInfoDatabase:    true,
		UseMessageDatabase:     true,
		UseSecretChats:         false,
		ApiId:                  apiId,
		ApiHash:                apiHash,
		SystemLanguageCode:     "en",
		DeviceModel:            "Server",
		SystemVersion:          "1.0.0",
		ApplicationVersion:     "1.0.0",
		EnableStorageOptimizer: true,
		IgnoreFileNames:        false,
	}

	logVerbosity := client.WithLogVerbosity(&client.SetLogVerbosityLevelRequest{
		NewVerbosityLevel: 0,
	})

	tdlibClient, err := client.NewClient(authorizer, logVerbosity)
	if err != nil {
		log.Fatalf("NewClient error: %s", err)
	}

	optionValue, err := tdlibClient.GetOption(&client.GetOptionRequest{
		Name: "version",
	})
	if err != nil {
		log.Fatalf("GetOption error: %s", err)
	}

	log.Printf("TDLib version: %s", optionValue.(*client.OptionValueString).Value)

	me, err := tdlibClient.GetMe()
	if err != nil {
		log.Fatalf("GetMe error: %s", err)
	}

	log.Printf("Me: %s %s [%s]", me.FirstName, me.LastName, me.Username)

	// Sub to updates
	listener := tdlibClient.GetListener()
	defer listener.Close()

	for update := range listener.Updates {
		// if update.GetClass() == client.ClassUpdate {
		// 	log.Printf("%#v", update)
		// }

		if update.GetClass() != client.ClassUpdate {
			continue
		}

		if update.GetType() != client.TypeUpdateNewMessage {
			continue
		}

		msg := update.(*client.UpdateNewMessage).Message

		log.Println(msg.Content.MessageContentType())
		if msg.Content.MessageContentType() == "messageText" {
			log.Println(msg.Content.(*client.MessageText).Text.Text)
			checkMessage(msg.Content.(*client.MessageText).Text.Text)
		} else if msg.Content.MessageContentType() == "messagePhoto" {
			log.Println(msg.Content.(*client.MessagePhoto).Caption.Text)
			checkMessage(msg.Content.(*client.MessagePhoto).Caption.Text)
		}
	}
}

func checkMessage(message string) {
	symb := binanceAPI.GetSymbols()
	for i := 0; i < len(symb); i++ {
		if strings.Index(message, symb[i]) != -1 {
			for z := 0; z < len(config.KeyWords); z++ {
				if strings.Index(strings.ToLower(message), strings.ToLower(config.KeyWords[z])) != -1 {
					price := binanceAPI.GetPriceBySymbol(symb[i])
					if price.Price != "" {
						binanceAPI.BinanceMakeOrder(price, symb[i], message)
					} else {
						fmt.Println("Symbol is incorect")
					}
					break
				}
			}
			break
		}
	}

	binanceAPI.GetAllOrders()
}
