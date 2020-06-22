package mercadobitcoin

// Trade represents a MercadoBitcoin trade object
type Trade struct {
	ID     int     `json:"tid"`
	Price  float32 ``
	Amount float64 `json:"amount"`
	Date   uint    ``
}
