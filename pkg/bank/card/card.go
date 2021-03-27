package card
import "bank/pkg/bank/types"


func PaymentSource(cards []types.Card) []types.PaymentSource {

	var paymentcards []types.PaymentSource

for _, operation := range cards {
	if(operation.Balance>0 && operation.Active == true){
       
	paymentcards = append(paymentcards, types.PaymentSource{ Type: "card", Number: string(operation.PAN), Balance:operation.Balance}) 

	}
	
}
return paymentcards
}