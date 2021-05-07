package binanceAPI

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	config "github.com/bot/config"
	db "github.com/bot/db"
	structures "github.com/bot/struct"
)

func GetPriceBySymbol(symb string) structures.PriceSymbol {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=" + symb + config.SecondSymbol)
	if err != nil {
		log.Println("err get Price By Symbol: ", err)
		return structures.PriceSymbol{}
	}

	defer resp.Body.Close()

	var dataSt structures.PriceSymbol

	err = json.NewDecoder(resp.Body).Decode(&dataSt)

	if err != nil {
		log.Println("err convert Price By Symbol: ", err)
		return structures.PriceSymbol{}
	}

	return dataSt
}

func GetAllOrders() {
	db := db.CreateConnection()
	data, e := db.Query(`SELECT * FROM "testtwo" WHERE status='NEW';`)
	if e != nil {
		fmt.Println("error from get all orders", e)
	}
	defer db.Close()
	var status structures.OrderData

	for data.Next() {
		if err := data.Scan(
			&status.Id,
			&status.Orderid,
			&status.Symbolbuy,
			&status.Symbolsell,
			&status.Buyprice,
			&status.Sellprice,
			&status.Datebuy,
			&status.Datesell,
			&status.Status,
			&status.Origquantity,
			&status.Executedquantity,
			&status.Message,
			&status.Profit); err != nil {
			fmt.Println("error from get status", err)
		}

		if status.Status == "NEW" {
			if time.Now().Unix()-int64(status.Datebuy) > 600 {
				price := GetPriceBySymbol(status.Symbolbuy)
				priceNow, err := strconv.ParseFloat(price.Price, 8)
				prevPrice, err := strconv.ParseFloat(status.Buyprice, 8)
				qunt, err := strconv.ParseFloat(status.Origquantity, 8)
				profit := (prevPrice - priceNow) * qunt
				sqlStatement := `
                    UPDATE "testtwo"
                    SET "sellprice" = $2, "datesell" = $3, "status" = $4, "profit" = $5
                    WHERE "id" = $1;`
				_, err = db.Exec(sqlStatement, status.Id, price.Price, time.Now().Unix(), "SELL", fmt.Sprintf("%f", profit))
				if err != nil {
					panic(err)
				}

				fmt.Println("Updated!")

			}
		}
	}
}
