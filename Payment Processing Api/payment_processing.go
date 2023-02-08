/*
Notice: Here I used the stripe-go library to process payments through the Stripe API.
To use other payment APIs, you'll need to consult the documentation for the specific payment gateway 
you're using to learn how to process payments through their API.
*/

package payment_processing

import (
	"encoding/json"
	"log"
	"net/http"
  "fmt"
  
  stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

// Payment represents a payment request
type Payment struct {
	Amount float64 `json:"amount"`
  Currency string `json:"currency"`
	Method string  `json:"method"`
  Description string `json:"description"`
}

func MakePayment() {
	http.HandleFunc("/pay", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var payment Payment
		if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Process the payment
		if err := processPayment(payment); err != nil {
			http.Error(w, "Failed to process payment", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

// 	log.Fatal(http.ListenAndServe(":8080", nil))
}


// processing payments using the Stripe payment gateway

func processPayment(payment Payment) error {
	stripe.Key = "YOUR_STRIPE_SECRET_KEY"

	params := &stripe.ChargeParams{
		Amount:   stripe.Int64(int64(payment.Amount * 100)), // convert amount to cents
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Desc:     stripe.String(string(payment.Description),
	}
                            
  if payment.Method == "credit_card" {
		params.SetSource("tok_visa") // replace with a real credit card token
	} else if payment.Method == "paypal" {
		params.SetSource("src_paypal") // replace with a real PayPal source
	} else {
		return fmt.Errorf("unsupported payment method: %s", payment.Method)
	}
                            
  _, err := charge.New(params)
                            
	return nil
}
