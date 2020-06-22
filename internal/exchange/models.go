package exchange

import (
	"time"
)

// Exchange represents a bitcoin exchange
type Exchange struct {
	ID   uint   `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Name string `gorm:"size:128"`
}

// Pair represents a currency pair
type Pair struct {
	ID    uint   `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Name  string `gorm:"size:8"`
	Base  string `gorm:"size:4"`
	Quote string `gorm:"size:4"`
}

// Trade represents an Exchange trade operation
type Trade struct {
	ID          uint      `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Price       float32   ``
	Amount      uint64    `json:"-"`
	AmountFloat float64   `gorm:"-" json:"amount"`
	Date        uint      ``
	Exchange    Exchange  `gorm:"FOREIGNKEY:id;association_foreignkey:exchange_id" json:"-"`
	ExchangeID  int       `json:"-"`
	Pair        Pair      `gorm:"FOREIGNKEY:id;association_foreignkey:pair_id" json:"-"`
	PairID      int       `json:"-"`
	BlockHeight int       `json:"-"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
