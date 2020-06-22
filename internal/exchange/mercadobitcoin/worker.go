package mercadobitcoin

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ngoline/quantocustaobitcoin/internal/exchange"
	"github.com/ngoline/quantocustaobitcoin/internal/httpjson"

	"github.com/jinzhu/gorm"
)

// Worker represents this exchange
type Worker struct {
	ExchangeID int
	PairID     int
	Name       string
}

// NewExchange creates a default exchange
func NewExchange() Worker {
	obj := Worker{Name: "Mercado Bitcoin", ExchangeID: 1, PairID: 1}
	return obj
}

// GetName returns worker name
func (w Worker) GetName() string {
	return w.Name
}

// Init initializes the database
func (w Worker) Init() *gorm.DB {
	cs := fmt.Sprintf("%s:%s@tcp(%s)/qcobtc?charset=utf8&parseTime=True&loc=Local", os.Getenv("SQL_USER"), os.Getenv("SQL_PASSWORD"), os.Getenv("SQL_ADDRESS"))
	// fmt.Println(cs)
	db, err := gorm.Open("mysql", cs)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func getTrades(begin uint, end uint) []Trade {
	// Get trades from exchange
	url := fmt.Sprintf("https://www.mercadobitcoin.net/api/BTC/trades/%d/%d", begin, end)
	var trades []Trade
	err := httpjson.GetJSON(url, &trades)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	if len(trades) == 1000 {
		lastTrade := trades[len(trades)-1]
		newTrades := getTrades(lastTrade.Date, end)
		for t := range newTrades {
			if newTrades[t].ID > lastTrade.ID {
				newTrades = newTrades[t+1:]
				break
			}
		}

		trades = append(trades, newTrades...)
	}

	return trades
}

// SyncData Syncs Bitcoin Blocks
func (w Worker) SyncData(db *gorm.DB) {
	now := uint(time.Now().Unix())

	var lastTrade exchange.Trade
	db.Where("exchange_id = ? AND pair_id = ?", w.ExchangeID, w.PairID).Order("id desc").First(&lastTrade)
	begin := lastTrade.Date + 1
	end := begin + 30

	for begin < now {
		trades := getTrades(begin, end)

		if len(trades) > 0 {
			var lastTrade exchange.Trade
			for t := range trades {
				amount := uint64(trades[t].Amount * 1e8)
				price := trades[t].Price

				if lastTrade != (exchange.Trade{}) {
					if trades[t].Date == lastTrade.Date {
						var totalAmount = lastTrade.Amount + amount
						price = ((float32(lastTrade.Amount) * lastTrade.Price) + (float32(amount) * trades[t].Price)) / float32(totalAmount)
						amount = totalAmount
					} else {
						db.Create(&lastTrade)
						lastTrade = exchange.Trade{}
					}
				}

				lastTrade.Amount = amount
				lastTrade.ExchangeID = w.ExchangeID
				lastTrade.PairID = w.PairID
				lastTrade.Price = price
				lastTrade.Date = trades[t].Date
			}

			if lastTrade != (exchange.Trade{}) {
				db.Create(&lastTrade)
			}

			begin = trades[len(trades)-1].Date + 1
		} else {
			begin = end + 1
		}

		end = begin + 30
	}
}
