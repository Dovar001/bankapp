package types



//Money  представляет собой денежную сумму в минимальных единицах(центы, копейки, и т.д)
type Money int64

type Currency string

//Коды валют
const(

	USD Currency = "USD"
	TJS Currency = "TJS"
	Rub Currency = "RUB"
)

//PAN представляет номер карты
type PAN string

//Card представляет  информацию о платёной карте 
type Card struct {
	ID int
	PAN PAN
	Balance Money
	Currency Currency
	Color string
	Name string
	Active bool
	Number string
}

type PaymentSource struct{
	Type string
	Number string
	Balance Money
}
