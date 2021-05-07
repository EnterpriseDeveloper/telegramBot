package binanceAPI

import (
	//	"context"
	"fmt"
	"math/rand"
	"time"

	"log"
	"strconv"

	//	binance "github.com/adshao/go-binance/v2"
	config "github.com/bot/config"
	db "github.com/bot/db"
	structures "github.com/bot/struct"
)

func BinanceMakeOrder(orderStr structures.PriceSymbol, symb string, message string) {
	i, err := strconv.ParseFloat(orderStr.Price, 8)
	if err != nil {
		log.Println("convert price error: ", err)
		return
	}
	quantity := fmt.Sprintf("%.2f", float64(config.PriceLimit)/i)
	// client := binance.NewClient(config.BinanceKey, config.BinanceSecret)
	// order, error := client.NewCreateOrderService().Symbol(orderStr.Symbol).
	// 	Side(binance.SideTypeBuy).Type(binance.OrderTypeLimit).
	// 	TimeInForce(binance.TimeInForceTypeGTC).Quantity(quantity).
	// 	Price(orderStr.Price).Do(context.Background())
	// if error != nil {
	// 	fmt.Println("make order with :", orderStr.Symbol, error)
	// 	return
	// }
	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 10000000
	order := structures.Order{
		OrderID:          rand.Intn(max-min+1) + min,
		Price:            orderStr.Price,
		Symbol:           symb,
		OrigQuantity:     quantity,
		ExecutedQuantity: quantity,
		Status:           "NEW",
	}
	fmt.Println("Bought: ", order.Price, order.Symbol)
	db.SetOrderToDb(order, symb, message)
}
