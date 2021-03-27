package card

import
(
	"fmt"
	"bank/pkg/bank/types"
)

func ExamplePaymentSource() {

	cards := []types.Card{

		{
			Balance:10_000_00,
			Active:true,
			PAN:"5058 xxxx xxxx 6600",
		},
		{
			Balance:10_000_00,
			Active:true,
			PAN:"5058 xxxx xxxx 4443",
		},
		
		{
			Balance:-10_000_00,
			Active:true,
			PAN:"5058 xxxx xxxx 0700",
		},
		
	}

	 paymentcard := PaymentSource(cards)
	 fmt.Println(paymentcard)

	//Output:[{card 5058 xxxx xxxx 6600 1000000} {card 5058 xxxx xxxx 4443 1000000}]
	
}