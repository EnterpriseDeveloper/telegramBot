package db

import (
	"fmt"
	"time"

	//	binance "github.com/adshao/go-binance/v2"
	structures "github.com/bot/struct"
	_ "github.com/lib/pq"
)

//func SetOrderToDb(order *binance.CreateOrderResponse, symb string, message string) {
func SetOrderToDb(order structures.Order, symb string, message string) {

	db := CreateConnection()
	defer db.Close()

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS "testtwo" (
		id SERIAL primary key,
		orderid integer,
		symbolbuy text,
		symbolsell text, 
		buyprice text,
		sellprice text,
		datebuy integer,
		datesell integer,
		status text,
		origquantity text,
		executedquantity text,
		message text,
		profit text )`)

	if err != nil {
		fmt.Println("create table err:", err)
	}

	insertDynStmt := `INSERT INTO "testtwo" ("orderid", "symbolbuy", "symbolsell", "buyprice", "sellprice", "datebuy", "datesell", "status", "origquantity", "executedquantity", "message", "profit") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, e := db.Exec(
		insertDynStmt,
		order.OrderID,
		symb,
		"", // TODO symbol sell
		order.Price,
		0,
		time.Now().Unix(),
		0,
		order.Status,
		order.OrigQuantity,
		order.ExecutedQuantity,
		message,
		0)
	CheckError("set order", e)
	fmt.Println("Order: ", order.OrderID)
}

func getOrdersById() {
	db := CreateConnection()
	defer db.Close()

	// спочатку провіряемо ордера котрі не закрилися. Відключаемо їх якшто время пройшло Х
	// якшто ордер закрився то продаемо його обратно через время X або по профіту який ми заробили
	// провіряемо оредар який не продався і питаемося його закрити пізніше і дешевше через время Х
	// Запитати Майю время свічки яка летит вверх. Граемо на довгосрок чи короткосрок
	// поставити ліміти на покупку в usdt записати в базу
}
